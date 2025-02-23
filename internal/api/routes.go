package api

import (
    "database/sql"

    "chatcollab/internal/api/handlers"
    "chatcollab/internal/config"
    "chatcollab/internal/services/agent"
    "chatcollab/internal/services/openrouter"
    "chatcollab/internal/services/session"
    "github.com/gin-gonic/gin"
)

type Server struct {
    router *gin.Engine
    cfg    *config.Config
    db     *sql.DB
}

func NewServer(cfg *config.Config, db *sql.DB) *Server {
    server := &Server{
        router: gin.Default(),
        cfg:    cfg,
        db:     db,
    }

    server.setupRoutes()
    return server
}

func (s *Server) setupRoutes() {
    sessionService := session.NewService(s.db)
    agentService := agent.NewService(s.db)
    openrouterClient := openrouter.NewClient(s.cfg.OpenRouterKey)

    sessionHandler := handlers.NewSessionHandler(sessionService)
    agentHandler := handlers.NewAgentHandler(agentService, openrouterClient)

    // Session routes
    s.router.POST("/sessions", sessionHandler.Create)
    s.router.GET("/sessions/:session_id", sessionHandler.Get)

    // Agent routes
    s.router.POST("/sessions/:session_id/agents", agentHandler.Create)
    s.router.GET("/sessions/:session_id/agents", agentHandler.List)
    s.router.DELETE("/sessions/:session_id/agents/:id", agentHandler.Delete)
    s.router.POST("/sessions/:session_id/agents/:id/chat", agentHandler.Chat)
}

func (s *Server) Run() {
    s.router.Run(":" + s.cfg.Port)
}