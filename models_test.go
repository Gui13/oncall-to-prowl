package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestOncallWebhook(t *testing.T) {
	var wh oncall_webhook
	event, err := os.Open("event.json")
	if err != nil {
		t.Fatal(err)
	}
	dec := json.NewDecoder(event)
	err = dec.Decode(&wh)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOncallWebhookWithoutLabels(t *testing.T) {
	var wh oncall_webhook
	event, err := os.Open("event_nolabels.json")
	if err != nil {
		t.Fatal(err)
	}
	dec := json.NewDecoder(event)
	err = dec.Decode(&wh)
	if err != nil {
		t.Fatal(err)
	}

	if len(wh.AlertPayload.Alerts) != 2 {
		t.Fatalf("Expected 2 alerts, got %d", len(wh.AlertPayload.Alerts))
	}
}
