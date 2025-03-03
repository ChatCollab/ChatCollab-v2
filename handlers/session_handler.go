package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/chatcollab/chatcollab/services"
)

// SessionHandler handles HTTP requests for sessions
type SessionHandler struct {
	service *services.SessionService
}

// NewSessionHandler creates a new SessionHandler
func NewSessionHandler() *SessionHandler {
	return &SessionHandler{
		service: services.NewSessionService(),
	}
}

// Create creates a new session
func (h *SessionHandler) Create(c *gin.Context) {
	session, err := h.service.CreateSession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, session)
}

// Get retrieves a session by ID
func (h *SessionHandler) Get(c *gin.Context) {
	id := c.Param("id")
	
	session, err := h.service.GetSession(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}
	
	c.JSON(http.StatusOK, session)
}

// UpdateHeartbeat updates a session's heartbeat
func (h *SessionHandler) UpdateHeartbeat(c *gin.Context) {
	id := c.Param("id")
	
	err := h.service.UpdateHeartbeat(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}
	
	c.Status(http.StatusNoContent)
}

// Delete deletes a session
func (h *SessionHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	
	err := h.service.DeleteSession(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.Status(http.StatusNoContent)
}

// List lists all sessions
func (h *SessionHandler) List(c *gin.Context) {
	sessions, err := h.service.ListSessions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, sessions)
}

// ListActive lists all active sessions
func (h *SessionHandler) ListActive(c *gin.Context) {
	// Default timeout of 5 minutes
	timeout := 5 * time.Minute
	
	sessions, err := h.service.ListActiveSessions(timeout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, sessions)
}

// RegisterRoutes registers routes for the session handler
func (h *SessionHandler) RegisterRoutes(router *gin.Engine) {
	sessions := router.Group("/api/sessions")
	{
		sessions.POST("", h.Create)
		sessions.GET("", h.List)
		sessions.GET("/active", h.ListActive)
		sessions.GET("/:id", h.Get)
		sessions.PUT("/:id/heartbeat", h.UpdateHeartbeat)
		sessions.DELETE("/:id", h.Delete)
	}
}