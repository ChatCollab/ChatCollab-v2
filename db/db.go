package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DB is the database connection
var DB *sql.DB

// Initialize sets up the database connection and creates tables if they don't exist
func Initialize(dbPath string) error {
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	err = createTables()
	if err != nil {
		return err
	}

	log.Println("Database initialized successfully")
	return nil
}

// createTables creates the necessary tables if they don't exist
func createTables() error {
	// Create Session table
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS sessions (
		id TEXT PRIMARY KEY,
		last_heartbeat DATETIME NOT NULL
	)`)
	if err != nil {
		return err
	}

	// Create Agent table
	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS agents (
		id TEXT PRIMARY KEY,
		is_online BOOLEAN NOT NULL,
		name TEXT NOT NULL,
		role TEXT NOT NULL,
		prompt TEXT NOT NULL,
		model TEXT NOT NULL,
		reasoning_log TEXT,
		session_id TEXT,
		FOREIGN KEY (session_id) REFERENCES sessions(id)
	)`)
	if err != nil {
		return err
	}

	// Create Message table
	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS messages (
		id TEXT PRIMARY KEY,
		created_at DATETIME NOT NULL,
		content TEXT NOT NULL,
		agent_id TEXT NOT NULL,
		session_id TEXT NOT NULL,
		FOREIGN KEY (agent_id) REFERENCES agents(id),
		FOREIGN KEY (session_id) REFERENCES sessions(id)
	)`)
	if err != nil {
		return err
	}

	return nil
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}