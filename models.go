package main

// This file contains the data structures used to decode the JSON payload from Oncall

type AlertPayload struct {
	EndsAt       string            `json:"endsAt"`
	StartsAt     string            `json:"startsAt"`
	Labels       map[string]string `json:"labels"`
	Status       string            `json:"status"`
	GroupLabels  map[string]string `json:"groupLabels"`
	CommonLabels map[string]string `json:"commonLabels"`
	Annotations  struct {
		Description string `json:"description"`
	} `json:"annotations"`
	Alerts []struct {
		Labels       map[string]string `json:"labels"`
		Status       string            `json:"status"`
		Annotations  map[string]string `json:"annotations"`
		StartsAt     string            `json:"startsAt"`
		EndsAt       string            `json:"endsAt"`
		GeneratorURL string            `json:"generatorURL"`
	} `json:"alerts"`
}

type Event struct {
	Type string `json:"type"`
	Time string `json:"time"`
}

type AlertGroup struct {
	ID    string `json:"id"`
	State string `json:"state"`
}

type oncall_webhook struct {
	Event        Event             `json:"event"`
	User         map[string]string `json:"user"`
	AlertGroup   AlertGroup        `json:"alert_group"`
	AlertPayload AlertPayload      `json:"alert_payload"`
}
