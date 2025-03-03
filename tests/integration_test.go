package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/chatcollab/chatcollab/db"
	"github.com/chatcollab/chatcollab/handlers"
)

func setupTestRouter() *gin.Engine {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	
	// Initialize router
	router := gin.Default()
	
	// Register API routes
	sessionHandler := handlers.NewSessionHandler()
	sessionHandler.RegisterRoutes(router)
	
	agentHandler := handlers.NewAgentHandler()
	agentHandler.RegisterRoutes(router)
	
	messageHandler := handlers.NewMessageHandler()
	messageHandler.RegisterRoutes(router)
	
	return router
}

func TestIntegrationFlow(t *testing.T) {
	// Create a temporary test database
	testDBPath := "./integration_test.db"
	defer os.Remove(testDBPath)
	
	// Initialize database
	err := db.Initialize(testDBPath)
	assert.NoError(t, err)
	defer db.Close()
	
	// Setup router
	router := setupTestRouter()
	
	// Step 1: Create a new session
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/sessions", nil)
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusCreated, w.Code)
	
	var session map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &session)
	assert.NoError(t, err)
	
	sessionID := session["id"].(string)
	assert.NotEmpty(t, sessionID)
	
	// Step 2: Create a new agent
	agentData := map[string]string{
		"name":      "Test Agent",
		"role":      "assistant",
		"prompt":    "You are a helpful assistant",
		"model":     "gpt-4",
		"sessionId": sessionID,
	}
	
	agentJSON, _ := json.Marshal(agentData)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/agents", bytes.NewBuffer(agentJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusCreated, w.Code)
	
	var agent map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &agent)
	assert.NoError(t, err)
	
	agentID := agent["id"].(string)
	assert.NotEmpty(t, agentID)
	
	// Step 3: Send a message
	messageData := map[string]string{
		"content":   "Hello, this is a test message",
		"agentId":   agentID,
		"sessionId": sessionID,
	}
	
	messageJSON, _ := json.Marshal(messageData)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/api/messages", bytes.NewBuffer(messageJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusCreated, w.Code)
	
	var message map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &message)
	assert.NoError(t, err)
	
	messageID := message["id"].(string)
	assert.NotEmpty(t, messageID)
	
	// Step 4: Get session messages
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/sessions/"+sessionID+"/messages", nil)
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var messages []interface{}
	err = json.Unmarshal(w.Body.Bytes(), &messages)
	assert.NoError(t, err)
	assert.Len(t, messages, 1)
	
	// Step 5: Get session agents
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/api/sessions/"+sessionID+"/agents", nil)
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusOK, w.Code)
	
	var agents []interface{}
	err = json.Unmarshal(w.Body.Bytes(), &agents)
	assert.NoError(t, err)
	assert.Len(t, agents, 1)
	
	// Step 6: Update agent status
	statusData := map[string]bool{
		"isOnline": false,
	}
	
	statusJSON, _ := json.Marshal(statusData)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("PUT", "/api/agents/"+agentID+"/online", bytes.NewBuffer(statusJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusNoContent, w.Code)
	
	// Step 7: Delete session (cleanup)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("DELETE", "/api/sessions/"+sessionID, nil)
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusNoContent, w.Code)
}