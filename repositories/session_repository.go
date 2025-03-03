package repositories

import (
	"time"

	"github.com/chatcollab/chatcollab/db"
	"github.com/chatcollab/chatcollab/models"
)

// SessionRepository handles database operations for sessions
type SessionRepository struct{}

// Create inserts a new session into the database
func (r *SessionRepository) Create(session *models.Session) error {
	_, err := db.DB.Exec(
		"INSERT INTO sessions (id, last_heartbeat) VALUES (?, ?)",
		session.ID, session.LastHeartbeat,
	)
	return err
}

// GetByID retrieves a session by its ID
func (r *SessionRepository) GetByID(id string) (*models.Session, error) {
	var session models.Session
	err := db.DB.QueryRow(
		"SELECT id, last_heartbeat FROM sessions WHERE id = ?",
		id,
	).Scan(&session.ID, &session.LastHeartbeat)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// Update updates an existing session
func (r *SessionRepository) Update(session *models.Session) error {
	_, err := db.DB.Exec(
		"UPDATE sessions SET last_heartbeat = ? WHERE id = ?",
		session.LastHeartbeat, session.ID,
	)
	return err
}

// Delete removes a session from the database
func (r *SessionRepository) Delete(id string) error {
	_, err := db.DB.Exec("DELETE FROM sessions WHERE id = ?", id)
	return err
}

// ListAll retrieves all sessions
func (r *SessionRepository) ListAll() ([]*models.Session, error) {
	rows, err := db.DB.Query("SELECT id, last_heartbeat FROM sessions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []*models.Session
	for rows.Next() {
		var session models.Session
		if err := rows.Scan(&session.ID, &session.LastHeartbeat); err != nil {
			return nil, err
		}
		sessions = append(sessions, &session)
	}

	return sessions, nil
}

// GetActiveSessions retrieves all active sessions based on the heartbeat timeout
func (r *SessionRepository) GetActiveSessions(timeout time.Duration) ([]*models.Session, error) {
	cutoffTime := time.Now().Add(-timeout)
	rows, err := db.DB.Query(
		"SELECT id, last_heartbeat FROM sessions WHERE last_heartbeat > ?",
		cutoffTime,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []*models.Session
	for rows.Next() {
		var session models.Session
		if err := rows.Scan(&session.ID, &session.LastHeartbeat); err != nil {
			return nil, err
		}
		sessions = append(sessions, &session)
	}

	return sessions, nil
}