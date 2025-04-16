package models

type PollRequest struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
}

type Poll struct {
	ID       string         `json:"id"`
	Question string         `json:"question"`
	Options  map[string]int `json:"options,omitempty"`
}
