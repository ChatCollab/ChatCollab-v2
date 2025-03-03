package models

import (
	"github.com/google/uuid"
)

// Agent represents a chat agent/participant
type Agent struct {
	ID           string `json:"id"`
	IsOnline     bool   `json:"isOnline"`
	Name         string `json:"name"`
	Role         string `json:"role"`
	Prompt       string `json:"prompt"`
	Model        string `json:"model"`
	ReasoningLog string `json:"reasoningLog"`
	SessionID    string `json:"sessionId"`
}

// NewAgent creates a new Agent with a generated UUID
func NewAgent(name, role, prompt, model, sessionID string) *Agent {
	return &Agent{
		ID:           uuid.New().String(),
		IsOnline:     true,
		Name:         name,
		Role:         role,
		Prompt:       prompt,
		Model:        model,
		ReasoningLog: "",
		SessionID:    sessionID,
	}
}

// SetOnline updates the agent's online status
func (a *Agent) SetOnline(isOnline bool) {
	a.IsOnline = isOnline
}

// AppendReasoningLog adds to the agent's reasoning log
func (a *Agent) AppendReasoningLog(log string) {
	if a.ReasoningLog == "" {
		a.ReasoningLog = log
	} else {
		a.ReasoningLog += "\n" + log
	}
}