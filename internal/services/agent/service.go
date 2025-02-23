package agent

import (
    "database/sql"
    "time"

    "chatcollab/internal/models"
    "github.com/google/uuid"
)

type Service struct {
    db *sql.DB
}

func NewService(db *sql.DB) *Service {
    return &Service{db: db}
}

func (s *Service) Create(sessionID, name, model string) (*models.Agent, error) {
    agent := &models.Agent{
        ID:         uuid.New().String(),
        SessionID:  sessionID,
        Name:       name,
        Model:      model,
        Status:     "active",
        CreatedAt:  time.Now(),
        LastActive: time.Now(),
    }

    query := `INSERT INTO agents (id, session_id, name, model, status, created_at, last_active) 
              VALUES (?, ?, ?, ?, ?, ?, ?)`
    _, err := s.db.Exec(query, agent.ID, agent.SessionID, agent.Name, agent.Model,
                       agent.Status, agent.CreatedAt, agent.LastActive)
    if err != nil {
        return nil, err
    }

    return agent, nil
}

func (s *Service) Delete(sessionID, id string) error {
    query := `DELETE FROM agents WHERE session_id = ? AND id = ?`
    _, err := s.db.Exec(query, sessionID, id)
    return err
}

func (s *Service) ListBySession(sessionID string) ([]*models.Agent, error) {
    query := `SELECT id, session_id, name, model, status, created_at, last_active 
              FROM agents WHERE session_id = ?`
    rows, err := s.db.Query(query, sessionID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var agents []*models.Agent
    for rows.Next() {
        agent := &models.Agent{}
        err := rows.Scan(&agent.ID, &agent.SessionID, &agent.Name, &agent.Model,
                        &agent.Status, &agent.CreatedAt, &agent.LastActive)
        if err != nil {
            return nil, err
        }
        agents = append(agents, agent)
    }
    return agents, nil
}

func (s *Service) GetModelByID(sessionID, id string) (string, error) {
    query := `SELECT model FROM agents WHERE session_id = ? AND id = ?`
    var model string
    err := s.db.QueryRow(query, sessionID, id).Scan(&model)
    return model, err
}