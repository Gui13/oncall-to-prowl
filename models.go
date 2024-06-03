package main

type AlertPayload struct {
	EndsAt      string            `json:"endsAt"`
	StartsAt    string            `json:"startsAt"`
	Labels      map[string]string `json:"labels"`
	Status      string            `json:"status"`
	Annotations struct {
		Description string `json:"description"`
	} `json:"annotations"`
}

type Event struct {
	Type string `json:"type"`
	Time string `json:"time"`
}

type AlertGroup struct {
	Labels map[string]string `json:"labels"`
	ID     string            `json:"id"`
	State  string            `json:"state"`
}

type oncall_webhook struct {
	Event        Event             `json:"event"`
	User         map[string]string `json:"user"`
	AlertGroup   AlertGroup        `json:"alert_group"`
	AlertPayload AlertPayload      `json:"alert_payload"`
}
