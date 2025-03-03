package models

import (
	"time"

	"github.com/google/uuid"
)

// Session represents a chat session
type Session struct {
	ID            string    `json:"id"`
	LastHeartbeat time.Time `json:"lastHeartbeat"`
}

// NewSession creates a new Session with a generated UUID
func NewSession() *Session {
	return &Session{
		ID:            uuid.New().String(),
		LastHeartbeat: time.Now(),
	}
}

// UpdateHeartbeat updates the session's last heartbeat time
func (s *Session) UpdateHeartbeat() {
	s.LastHeartbeat = time.Now()
}

// IsActive checks if the session is active based on a timeout duration
func (s *Session) IsActive(timeout time.Duration) bool {
	return time.Since(s.LastHeartbeat) < timeout
}