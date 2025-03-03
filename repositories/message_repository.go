package repositories

import (
	"time"

	"github.com/chatcollab/chatcollab/db"
	"github.com/chatcollab/chatcollab/models"
)

// MessageRepository handles database operations for messages
type MessageRepository struct{}

// Create inserts a new message into the database
func (r *MessageRepository) Create(message *models.Message) error {
	_, err := db.DB.Exec(
		"INSERT INTO messages (id, created_at, content, agent_id, session_id) VALUES (?, ?, ?, ?, ?)",
		message.ID, message.CreatedAt, message.Content, message.AgentID, message.SessionID,
	)
	return err
}

// GetByID retrieves a message by its ID
func (r *MessageRepository) GetByID(id string) (*models.Message, error) {
	var message models.Message
	err := db.DB.QueryRow(
		"SELECT id, created_at, content, agent_id, session_id FROM messages WHERE id = ?",
		id,
	).Scan(&message.ID, &message.CreatedAt, &message.Content, &message.AgentID, &message.SessionID)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

// Update updates an existing message
func (r *MessageRepository) Update(message *models.Message) error {
	_, err := db.DB.Exec(
		"UPDATE messages SET content = ? WHERE id = ?",
		message.Content, message.ID,
	)
	return err
}

// Delete removes a message from the database
func (r *MessageRepository) Delete(id string) error {
	_, err := db.DB.Exec("DELETE FROM messages WHERE id = ?", id)
	return err
}

// GetBySessionID retrieves all messages for a specific session
func (r *MessageRepository) GetBySessionID(sessionID string) ([]*models.Message, error) {
	rows, err := db.DB.Query(
		"SELECT id, created_at, content, agent_id, session_id FROM messages WHERE session_id = ? ORDER BY created_at",
		sessionID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.ID, &message.CreatedAt, &message.Content, &message.AgentID, &message.SessionID); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	return messages, nil
}

// GetByAgentID retrieves all messages for a specific agent
func (r *MessageRepository) GetByAgentID(agentID string) ([]*models.Message, error) {
	rows, err := db.DB.Query(
		"SELECT id, created_at, content, agent_id, session_id FROM messages WHERE agent_id = ? ORDER BY created_at",
		agentID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.ID, &message.CreatedAt, &message.Content, &message.AgentID, &message.SessionID); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	return messages, nil
}

// GetMessagesAfter retrieves all messages created after a specific time
func (r *MessageRepository) GetMessagesAfter(sessionID string, after time.Time) ([]*models.Message, error) {
	rows, err := db.DB.Query(
		"SELECT id, created_at, content, agent_id, session_id FROM messages WHERE session_id = ? AND created_at > ? ORDER BY created_at",
		sessionID, after,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		var message models.Message
		if err := rows.Scan(&message.ID, &message.CreatedAt, &message.Content, &message.AgentID, &message.SessionID); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}

	return messages, nil
}