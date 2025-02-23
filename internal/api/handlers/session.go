package handlers

import (
    "net/http"

    "chatcollab/internal/services/session"
    "github.com/gin-gonic/gin"
)

type SessionHandler struct {
    service *session.Service
}

func NewSessionHandler(service *session.Service) *SessionHandler {
    return &SessionHandler{service: service}
}

func (h *SessionHandler) Create(c *gin.Context) {
    var req struct {
        Name string `json:"name" binding:"required"`
    }

    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    session, err := h.service.Create(req.Name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, session)
}

func (h *SessionHandler) Get(c *gin.Context) {
    id := c.Param("session_id")
    session, err := h.service.Get(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
        return
    }
    c.JSON(http.StatusOK, session)
}