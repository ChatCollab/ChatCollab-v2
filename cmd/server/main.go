package main

import (
    "log"

    "chatcollab/internal/api"
    "chatcollab/internal/config"
    "chatcollab/internal/db"
)

func main() {
    cfg, err := config.Load()
    if err != nil {
        log.Fatal(err)
    }

    database, err := db.Connect(cfg.DatabaseURL)
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    if err := db.Migrate(database); err != nil {
        log.Fatal(err)
    }

    server := api.NewServer(cfg, database)
    server.Run()
}