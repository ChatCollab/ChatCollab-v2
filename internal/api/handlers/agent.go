package handlers

import (
    "net/http"

    "chatcollab/internal/services/agent"
    "chatcollab/internal/services/openrouter"
    "github.com/gin-gonic/gin"
)

type AgentHandler struct {
    service *agent.Service
    client  *openrouter.Client
}

func NewAgentHandler(service *agent.Service, client *openrouter.Client) *AgentHandler {
    return &AgentHandler{
        service: service,
        client:  client,
    }
}

func (h *AgentHandler) Create(c *gin.Context) {
    sessionID := c.Param("session_id")
    var req struct {
        Name  string `json:"name" binding:"required"`
        Model string `json:"model" binding:"required"`
    }

    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    agent, err := h.service.Create(sessionID, req.Name, req.Model)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, agent)
}

func (h *AgentHandler) List(c *gin.Context) {
    sessionID := c.Param("session_id")
    agents, err := h.service.ListBySession(sessionID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, agents)
}

func (h *AgentHandler) Delete(c *gin.Context) {
    sessionID := c.Param("session_id")
    id := c.Param("id")
    err := h.service.Delete(sessionID, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Agent deleted successfully"})
}

type ChatRequest struct {
    Prompt string `json:"prompt" binding:"required"`
}

func (h *AgentHandler) Chat(c *gin.Context) {
    agentID := c.Param("id")
    sessionID := c.Param("session_id")

    var req ChatRequest
    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
	model, err := h.service.GetModelByID(sessionID, agentID)

	if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
	}

    response, err := h.client.Chat(model, req.Prompt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, response)
}