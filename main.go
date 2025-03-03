package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/chatcollab/chatcollab/db"
	"github.com/chatcollab/chatcollab/handlers"
)

func main() {
	// Initialize database
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "chatcollab.db"
	}
	
	if err := db.Initialize(dbPath); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()
	
	// Initialize Gin router
	router := gin.Default()
	
	// Serve static files
	router.Static("/static", "./static")
	router.Static("/css", "./templates/css")
	
	// Load HTML templates
	router.LoadHTMLGlob("templates/*.html")
	
	// Home page
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "ChatCollab",
		})
	})
	
	// Register API routes
	sessionHandler := handlers.NewSessionHandler()
	sessionHandler.RegisterRoutes(router)
	
	agentHandler := handlers.NewAgentHandler()
	agentHandler.RegisterRoutes(router)
	
	messageHandler := handlers.NewMessageHandler()
	messageHandler.RegisterRoutes(router)
	
	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}