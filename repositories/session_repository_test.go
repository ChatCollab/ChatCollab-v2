package repositories

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/chatcollab/chatcollab/db"
	"github.com/chatcollab/chatcollab/models"
)

func setupTestDB(t *testing.T) func() {
	// Create a temporary test database
	testDBPath := "./repo_test.db"
	
	// Clean up from previous tests
	_ = os.Remove(testDBPath)
	
	// Initialize the database
	err := db.Initialize(testDBPath)
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}
	
	// Return cleanup function
	return func() {
		db.Close()
		os.Remove(testDBPath)
	}
}

func TestSessionRepository(t *testing.T) {
	cleanup := setupTestDB(t)
	defer cleanup()
	
	repo := SessionRepository{}
	
	// Test Create
	session := models.NewSession()
	err := repo.Create(session)
	assert.NoError(t, err)
	
	// Test GetByID
	retrieved, err := repo.GetByID(session.ID)
	assert.NoError(t, err)
	assert.Equal(t, session.ID, retrieved.ID)
	assert.WithinDuration(t, session.LastHeartbeat, retrieved.LastHeartbeat, time.Second)
	
	// Test Update
	updatedTime := time.Now().Add(1 * time.Hour)
	session.LastHeartbeat = updatedTime
	err = repo.Update(session)
	assert.NoError(t, err)
	
	// Verify update
	retrieved, err = repo.GetByID(session.ID)
	assert.NoError(t, err)
	assert.WithinDuration(t, updatedTime, retrieved.LastHeartbeat, time.Second)
	
	// Test ListAll
	sessions, err := repo.ListAll()
	assert.NoError(t, err)
	assert.Len(t, sessions, 1)
	assert.Equal(t, session.ID, sessions[0].ID)
	
	// Test GetActiveSessions
	activeSessions, err := repo.GetActiveSessions(2 * time.Hour)
	assert.NoError(t, err)
	assert.Len(t, activeSessions, 1)
	
	// Add an inactive session
	oldSession := models.NewSession()
	oldSession.LastHeartbeat = time.Now().Add(-3 * time.Hour)
	err = repo.Create(oldSession)
	assert.NoError(t, err)
	
	// Verify we still get only one active session
	activeSessions, err = repo.GetActiveSessions(2 * time.Hour)
	assert.NoError(t, err)
	assert.Len(t, activeSessions, 1)
	
	// Test Delete
	err = repo.Delete(session.ID)
	assert.NoError(t, err)
	
	// Verify deletion
	_, err = repo.GetByID(session.ID)
	assert.Error(t, err) // Should error because session was deleted
}