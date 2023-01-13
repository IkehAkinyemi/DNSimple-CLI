package model

import (
	"time"
)

// ZoneRecord defines a zone record type.
type ZoneRecord struct {
	ID           int       `json:"id"`
	ZoneID       string    `json:"zone_id"`
	Name         string    `json:"name"`
	Content      string    `json:"content"`
	TTL          int       `json:"ttl"`
	Type         string    `json:"type"`
	Regions      []string  `json:"regions"`
	SystemRecord bool      `json:"system_record"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ZoneRecords defines the expected response type.
type ZoneRecords struct {
	Data       []ZoneRecord `json:"data"`
	Pagination struct {
		CurrentPage  int `json:"current_page"`
		PerPage      int `json:"per_page"`
		TotalEntries int `json:"total_entries"`
		TotalPages   int `json:"total_page"`
	} `json:"pagination"`
}

// Envelope encapsulates data.
type Envelope map[string]any

// Config defines the requirement for CLI configuration.
type Config struct {
	AccessToken string `yaml:"token"`
	Env         string `yaml:"env"`
	ZoneName    string `yaml:"zone"`
	AccountID   int    `yaml:"account"`
}
