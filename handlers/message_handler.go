package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/chatcollab/chatcollab/services"
)

// MessageHandler handles HTTP requests for messages
type MessageHandler struct {
	service *services.MessageService
}

// NewMessageHandler creates a new MessageHandler
func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		service: services.NewMessageService(),
	}
}

// Create creates a new message
func (h *MessageHandler) Create(c *gin.Context) {
	var input struct {
		Content   string `json:"content" binding:"required"`
		AgentID   string `json:"agentId" binding:"required"`
		SessionID string `json:"sessionId" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	message, err := h.service.CreateMessage(input.Content, input.AgentID, input.SessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, message)
}

// Get retrieves a message by ID
func (h *MessageHandler) Get(c *gin.Context) {
	id := c.Param("id")
	
	message, err := h.service.GetMessage(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}
	
	c.JSON(http.StatusOK, message)
}

// Update updates a message's content
func (h *MessageHandler) Update(c *gin.Context) {
	id := c.Param("id")
	
	var input struct {
		Content string `json:"content" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := h.service.UpdateMessage(id, input.Content)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}
	
	c.Status(http.StatusNoContent)
}

// Delete deletes a message
func (h *MessageHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	
	err := h.service.DeleteMessage(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.Status(http.StatusNoContent)
}

// GetSessionMessages retrieves all messages for a session
func (h *MessageHandler) GetSessionMessages(c *gin.Context) {
	sessionID := c.Param("id")
	
	messages, err := h.service.GetSessionMessages(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, messages)
}

// GetAgentMessages retrieves all messages for an agent
func (h *MessageHandler) GetAgentMessages(c *gin.Context) {
	agentID := c.Param("id")
	
	messages, err := h.service.GetAgentMessages(agentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, messages)
}

// GetNewMessages retrieves all messages after a specific time
func (h *MessageHandler) GetNewMessages(c *gin.Context) {
	sessionID := c.Param("id")
	
	var input struct {
		After time.Time `json:"after" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	messages, err := h.service.GetNewMessages(sessionID, input.After)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, messages)
}

// RegisterRoutes registers routes for the message handler
func (h *MessageHandler) RegisterRoutes(router *gin.Engine) {
	messages := router.Group("/api/messages")
	{
		messages.POST("", h.Create)
		messages.GET("/:id", h.Get)
		messages.PUT("/:id", h.Update)
		messages.DELETE("/:id", h.Delete)
	}
	
	router.GET("/api/sessions/:id/messages", h.GetSessionMessages)
	router.POST("/api/sessions/:id/messages/new", h.GetNewMessages)
	router.GET("/api/agents/:id/messages", h.GetAgentMessages)
}