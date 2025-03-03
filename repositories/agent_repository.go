package repositories

import (
	"github.com/chatcollab/chatcollab/db"
	"github.com/chatcollab/chatcollab/models"
)

// AgentRepository handles database operations for agents
type AgentRepository struct{}

// Create inserts a new agent into the database
func (r *AgentRepository) Create(agent *models.Agent) error {
	_, err := db.DB.Exec(
		"INSERT INTO agents (id, is_online, name, role, prompt, model, reasoning_log, session_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		agent.ID, agent.IsOnline, agent.Name, agent.Role, agent.Prompt, agent.Model, agent.ReasoningLog, agent.SessionID,
	)
	return err
}

// GetByID retrieves an agent by its ID
func (r *AgentRepository) GetByID(id string) (*models.Agent, error) {
	var agent models.Agent
	err := db.DB.QueryRow(
		"SELECT id, is_online, name, role, prompt, model, reasoning_log, session_id FROM agents WHERE id = ?",
		id,
	).Scan(&agent.ID, &agent.IsOnline, &agent.Name, &agent.Role, &agent.Prompt, &agent.Model, &agent.ReasoningLog, &agent.SessionID)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

// Update updates an existing agent
func (r *AgentRepository) Update(agent *models.Agent) error {
	_, err := db.DB.Exec(
		"UPDATE agents SET is_online = ?, name = ?, role = ?, prompt = ?, model = ?, reasoning_log = ?, session_id = ? WHERE id = ?",
		agent.IsOnline, agent.Name, agent.Role, agent.Prompt, agent.Model, agent.ReasoningLog, agent.SessionID, agent.ID,
	)
	return err
}

// Delete removes an agent from the database
func (r *AgentRepository) Delete(id string) error {
	_, err := db.DB.Exec("DELETE FROM agents WHERE id = ?", id)
	return err
}

// ListAll retrieves all agents
func (r *AgentRepository) ListAll() ([]*models.Agent, error) {
	rows, err := db.DB.Query("SELECT id, is_online, name, role, prompt, model, reasoning_log, session_id FROM agents")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agents []*models.Agent
	for rows.Next() {
		var agent models.Agent
		if err := rows.Scan(&agent.ID, &agent.IsOnline, &agent.Name, &agent.Role, &agent.Prompt, &agent.Model, &agent.ReasoningLog, &agent.SessionID); err != nil {
			return nil, err
		}
		agents = append(agents, &agent)
	}

	return agents, nil
}

// GetBySessionID retrieves all agents for a specific session
func (r *AgentRepository) GetBySessionID(sessionID string) ([]*models.Agent, error) {
	rows, err := db.DB.Query(
		"SELECT id, is_online, name, role, prompt, model, reasoning_log, session_id FROM agents WHERE session_id = ?",
		sessionID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agents []*models.Agent
	for rows.Next() {
		var agent models.Agent
		if err := rows.Scan(&agent.ID, &agent.IsOnline, &agent.Name, &agent.Role, &agent.Prompt, &agent.Model, &agent.ReasoningLog, &agent.SessionID); err != nil {
			return nil, err
		}
		agents = append(agents, &agent)
	}

	return agents, nil
}