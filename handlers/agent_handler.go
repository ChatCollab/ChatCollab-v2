package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/chatcollab/chatcollab/services"
)

// AgentHandler handles HTTP requests for agents
type AgentHandler struct {
	service *services.AgentService
}

// NewAgentHandler creates a new AgentHandler
func NewAgentHandler() *AgentHandler {
	return &AgentHandler{
		service: services.NewAgentService(),
	}
}

// Create creates a new agent
func (h *AgentHandler) Create(c *gin.Context) {
	var input struct {
		Name      string `json:"name" binding:"required"`
		Role      string `json:"role" binding:"required"`
		Prompt    string `json:"prompt" binding:"required"`
		Model     string `json:"model" binding:"required"`
		SessionID string `json:"sessionId" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	agent, err := h.service.CreateAgent(input.Name, input.Role, input.Prompt, input.Model, input.SessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, agent)
}

// Get retrieves an agent by ID
func (h *AgentHandler) Get(c *gin.Context) {
	id := c.Param("id")
	
	agent, err := h.service.GetAgent(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agent not found"})
		return
	}
	
	c.JSON(http.StatusOK, agent)
}

// Update updates an agent
func (h *AgentHandler) Update(c *gin.Context) {
	id := c.Param("id")
	
	agent, err := h.service.GetAgent(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agent not found"})
		return
	}
	
	var input struct {
		IsOnline     *bool   `json:"isOnline"`
		Name         *string `json:"name"`
		Role         *string `json:"role"`
		Prompt       *string `json:"prompt"`
		Model        *string `json:"model"`
		ReasoningLog *string `json:"reasoningLog"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if input.IsOnline != nil {
		agent.IsOnline = *input.IsOnline
	}
	if input.Name != nil {
		agent.Name = *input.Name
	}
	if input.Role != nil {
		agent.Role = *input.Role
	}
	if input.Prompt != nil {
		agent.Prompt = *input.Prompt
	}
	if input.Model != nil {
		agent.Model = *input.Model
	}
	if input.ReasoningLog != nil {
		agent.ReasoningLog = *input.ReasoningLog
	}
	
	if err := h.service.UpdateAgent(agent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, agent)
}

// Delete deletes an agent
func (h *AgentHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	
	err := h.service.DeleteAgent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.Status(http.StatusNoContent)
}

// List lists all agents
func (h *AgentHandler) List(c *gin.Context) {
	agents, err := h.service.ListAgents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, agents)
}

// ListSessionAgents lists all agents for a session
func (h *AgentHandler) ListSessionAgents(c *gin.Context) {
	sessionID := c.Param("id")
	
	agents, err := h.service.ListSessionAgents(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, agents)
}

// UpdateOnlineStatus updates an agent's online status
func (h *AgentHandler) UpdateOnlineStatus(c *gin.Context) {
	id := c.Param("id")
	
	var input struct {
		IsOnline *bool `json:"isOnline" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if input.IsOnline == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "isOnline field is required"})
		return
	}
	
	err := h.service.SetAgentOnlineStatus(id, *input.IsOnline)
	
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agent not found"})
		return
	}
	
	c.Status(http.StatusNoContent)
}

// AppendReasoningLog appends to an agent's reasoning log
func (h *AgentHandler) AppendReasoningLog(c *gin.Context) {
	id := c.Param("id")
	
	var input struct {
		Log string `json:"log" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := h.service.AppendAgentReasoningLog(id, input.Log)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agent not found"})
		return
	}
	
	c.Status(http.StatusNoContent)
}

// RegisterRoutes registers routes for the agent handler
func (h *AgentHandler) RegisterRoutes(router *gin.Engine) {
	agents := router.Group("/api/agents")
	{
		agents.POST("", h.Create)
		agents.GET("", h.List)
		agents.GET("/:id", h.Get)
		agents.PUT("/:id", h.Update)
		agents.DELETE("/:id", h.Delete)
		agents.PUT("/:id/online", h.UpdateOnlineStatus)
		agents.POST("/:id/reasoning", h.AppendReasoningLog)
	}
	
	router.GET("/api/sessions/:id/agents", h.ListSessionAgents)
}