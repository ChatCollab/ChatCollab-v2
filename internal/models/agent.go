package models

import "time"

type Agent struct {
    ID          string    `json:"id"`
    SessionID   string    `json:"session_id"`
    Name        string    `json:"name"`
    Model       string    `json:"model"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
    LastActive  time.Time `json:"last_active"`
}