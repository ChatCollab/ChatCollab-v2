package db

import "database/sql"

func Migrate(db *sql.DB) error {
    queries := []string{
        `CREATE TABLE IF NOT EXISTS sessions (
            id TEXT PRIMARY KEY,
            name TEXT NOT NULL,
            created_at DATETIME NOT NULL,
            last_active DATETIME NOT NULL
        )`,
        `CREATE TABLE IF NOT EXISTS agents (
            id TEXT PRIMARY KEY,
            session_id TEXT NOT NULL,
            name TEXT NOT NULL,
            model TEXT NOT NULL,
            status TEXT NOT NULL,
            created_at DATETIME NOT NULL,
            last_active DATETIME NOT NULL,
            FOREIGN KEY (session_id) REFERENCES sessions(id) ON DELETE CASCADE
        )`,
    }

    for _, query := range queries {
        if _, err := db.Exec(query); err != nil {
            return err
        }
    }
    
    return nil
}