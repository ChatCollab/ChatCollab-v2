package models

import "time"

type Session struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    CreatedAt   time.Time `json:"created_at"`
    LastActive  time.Time `json:"last_active"`
}