package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var ProwlApiKey string
var VerboseMode bool

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func readyHandler(writer http.ResponseWriter, request *http.Request) {
	if ProwlApiKey != "" {
		writer.WriteHeader(200)
	} else {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte("Prowl API key not set"))
	}
}

func eventHandler(writer http.ResponseWriter, request *http.Request) {
	// retrieve body as bytes for easier dump if error
	defer request.Body.Close()
	bodyBytes, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Printf("Error receiving request body: %s\n", err)
		writer.WriteHeader(500)
		return
	}

	oncall := oncall_webhook{}
	err = json.Unmarshal(bodyBytes, &oncall)
	if err != nil {
		fmt.Printf("Error decoding oncall payload: %s\n", err)
		if VerboseMode {
			fmt.Printf("payload is: %v", string(bodyBytes))
		}
		writer.WriteHeader(400)
		return
	}

	if VerboseMode {
		fmt.Printf("Oncall event:\n%+v\n", oncall)
	}

	prowlClient := NewProwClient(ProwlApiKey)

	err = prowlClient.sendAlert(&oncall.AlertPayload)
	if err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte("Prowl call failed: " + err.Error()))
	} else {
		_, _ = writer.Write([]byte("Handled the event!\n"))
	}

}

func main() {
	var serverPort string
	var serverAddr string

	flag.StringVar(&ProwlApiKey, "k", os.Getenv("PROWL_API_KEY"), "The prowl API key value, you can pass it as PROWL_API_KEY in the environment")
	flag.StringVar(&serverAddr, "u", os.Getenv("SERVER_ADDRESS"), "The server address on which to listen, default is 0.0.0.0")
	flag.StringVar(&serverPort, "p", os.Getenv("SERVER_PORT"), "Server port to listen on, default is 8080")
	flag.BoolVar(&VerboseMode, "v", false, "Be verbose")
	flag.Parse()

	if ProwlApiKey == "" {
		fmt.Println("Prowl API key not set, please set it using PROWL_API_KEY or launching with -k xxx")
		os.Exit(1)
	}

	if VerboseMode {
		fmt.Println("Verbose mode enabled")
		fmt.Printf("Prowl Api Key: %s\n", ProwlApiKey)
	}

	if serverAddr == "" {
		serverAddr = "0.0.0.0"
	}

	if serverPort == "" {
		serverPort = "8080"
	}
	listenAddr := fmt.Sprintf("%s:%s", serverAddr, serverPort)
	fmt.Printf("Listening on http://%s\n", listenAddr)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/ready", readyHandler)
	mux.Handle("POST /event", httpLogger(eventHandler))

	err := http.ListenAndServe(listenAddr, mux)
	if err != nil {
		panic(err)
	}

}

func httpLogger(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	handler := http.HandlerFunc(next)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		handler.ServeHTTP(w, r)
		finish := time.Since(now)
		fmt.Printf("%s: [%s] %s %s in %s\n", time.Now().Format("2006-01-02T15:04:05.000Z07:00"), r.RequestURI, r.RemoteAddr, r.Method, finish)
	})
}
