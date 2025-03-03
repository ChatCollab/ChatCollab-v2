package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMessage(t *testing.T) {
	content := "Hello, this is a test message"
	agentID := "agent123"
	sessionID := "session123"
	
	message := NewMessage(content, agentID, sessionID)
	
	assert.NotEmpty(t, message.ID, "Message ID should not be empty")
	assert.Equal(t, content, message.Content, "Message content should match input")
	assert.Equal(t, agentID, message.AgentID, "Message agentID should match input")
	assert.Equal(t, sessionID, message.SessionID, "Message sessionID should match input")
	assert.WithinDuration(t, time.Now(), message.CreatedAt, 2*time.Second, "CreatedAt should be close to current time")
}