package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAgent(t *testing.T) {
	name := "Test Agent"
	role := "assistant"
	prompt := "You are a helpful assistant"
	model := "gpt-4"
	sessionID := "session123"
	
	agent := NewAgent(name, role, prompt, model, sessionID)
	
	assert.NotEmpty(t, agent.ID, "Agent ID should not be empty")
	assert.Equal(t, name, agent.Name, "Agent name should match input")
	assert.Equal(t, role, agent.Role, "Agent role should match input")
	assert.Equal(t, prompt, agent.Prompt, "Agent prompt should match input")
	assert.Equal(t, model, agent.Model, "Agent model should match input")
	assert.Equal(t, sessionID, agent.SessionID, "Agent sessionID should match input")
	assert.True(t, agent.IsOnline, "New agent should be online by default")
	assert.Empty(t, agent.ReasoningLog, "New agent should have empty reasoning log")
}

func TestSetOnline(t *testing.T) {
	agent := NewAgent("Test", "assistant", "prompt", "model", "session123")
	
	// By default, agent is online
	assert.True(t, agent.IsOnline, "Agent should be online by default")
	
	// Set to offline
	agent.SetOnline(false)
	assert.False(t, agent.IsOnline, "Agent should be offline after SetOnline(false)")
	
	// Set back to online
	agent.SetOnline(true)
	assert.True(t, agent.IsOnline, "Agent should be online after SetOnline(true)")
}

func TestAppendReasoningLog(t *testing.T) {
	agent := NewAgent("Test", "assistant", "prompt", "model", "session123")
	
	// Initial log should be empty
	assert.Empty(t, agent.ReasoningLog, "New agent should have empty reasoning log")
	
	// Append first log entry
	firstLog := "First reasoning entry"
	agent.AppendReasoningLog(firstLog)
	assert.Equal(t, firstLog, agent.ReasoningLog, "Reasoning log should contain first entry")
	
	// Append second log entry
	secondLog := "Second reasoning entry"
	agent.AppendReasoningLog(secondLog)
	assert.Equal(t, firstLog+"\n"+secondLog, agent.ReasoningLog, "Reasoning log should contain both entries")
}