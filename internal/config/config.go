package config

import "os"

type Config struct {
    DatabaseURL     string
    OpenRouterKey   string
    Port            string
}

func Load() (*Config, error) {
    return &Config{
        DatabaseURL:     getEnvOrDefault("DATABASE_URL", "sqlite3://agent_manager.db"),
        OpenRouterKey:   os.Getenv("OPENROUTER_API_KEY"),
        Port:            getEnvOrDefault("PORT", "8080"),
    }, nil
}

func getEnvOrDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}