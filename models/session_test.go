package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewSession(t *testing.T) {
	session := NewSession()
	
	assert.NotEmpty(t, session.ID, "Session ID should not be empty")
	assert.WithinDuration(t, time.Now(), session.LastHeartbeat, 2*time.Second, "LastHeartbeat should be close to current time")
}

func TestUpdateHeartbeat(t *testing.T) {
	session := NewSession()
	
	// Make the test sleep a little to ensure we get a different timestamp
	time.Sleep(100 * time.Millisecond)
	
	oldHeartbeat := session.LastHeartbeat
	session.UpdateHeartbeat()
	
	assert.True(t, session.LastHeartbeat.After(oldHeartbeat), "LastHeartbeat should be updated to a later time")
}

func TestIsActive(t *testing.T) {
	session := NewSession()
	
	// Check if the session is active with a 5-minute timeout
	assert.True(t, session.IsActive(5*time.Minute), "New session should be active")
	
	// Manually set the heartbeat to be older
	session.LastHeartbeat = time.Now().Add(-10 * time.Minute)
	
	// Check if the session is now inactive
	assert.False(t, session.IsActive(5*time.Minute), "Session with old heartbeat should be inactive")
}