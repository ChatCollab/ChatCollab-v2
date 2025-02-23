package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func Connect(dbURL string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "agent_manager.db")
    if err != nil {
        return nil, err
    }
    
    if err := db.Ping(); err != nil {
        return nil, err
    }
    
    return db, nil
}