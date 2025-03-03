package models

import (
	"time"

	"github.com/google/uuid"
)

// Message represents a chat message
type Message struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
	AgentID   string    `json:"agentId"`
	SessionID string    `json:"sessionId"`
}

// NewMessage creates a new Message with a generated UUID
func NewMessage(content, agentID, sessionID string) *Message {
	return &Message{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		Content:   content,
		AgentID:   agentID,
		SessionID: sessionID,
	}
}