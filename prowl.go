package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type ProwClient struct {
	ProwlApiKey string
	ProwlApiUrl string
}

type ProwlAddQuery struct {
	Application string `json:"application"`
	Event       string `json:"event"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	URL         string `json:"url"`
}

func (p *ProwClient) sendAlert(payload *AlertPayload) error {
	return p.add("oncall", payload.Labels["alertname"], payload.Annotations.Description, 0, "")
}

func (p *ProwClient) add(application string, event string, description string, priority int, URL string) error {
	url := fmt.Sprintf("%s/%s", p.ProwlApiUrl, "add")
	request, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	query := request.URL.Query()
	query.Add("apikey", p.ProwlApiKey)
	query.Add("application", application)
	query.Add("event", event)
	query.Add("description", description)
	query.Add("priority", strconv.Itoa(priority))
	query.Add("url", URL)
	request.URL.RawQuery = query.Encode()

	fmt.Printf("Sending notification for %s (%s)\n", event, description)
	client := http.Client{}
	done, err := client.Do(request)
	if err != nil {
		return err
	}

	if done.StatusCode != http.StatusOK {
		return fmt.Errorf("error posting request to %s: %v", p.ProwlApiUrl, request.Body)
	}

	return nil
}

func NewProwClient(prowlApiKey string) *ProwClient {
	return &ProwClient{prowlApiKey, "https://api.prowlapp.com/publicapi"}
}
