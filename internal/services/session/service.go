package session

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

func (s *Service) Create(name string) (*models.Session, error) {
    session := &models.Session{
        ID:         uuid.New().String(),
        Name:       name,
        CreatedAt:  time.Now(),
        LastActive: time.Now(),
    }

    query := `INSERT INTO sessions (id, name, created_at, last_active) VALUES (?, ?, ?, ?)`
    _, err := s.db.Exec(query, session.ID, session.Name, session.CreatedAt, session.LastActive)
    if err != nil {
        return nil, err
    }

    return session, nil
}

func (s *Service) Get(id string) (*models.Session, error) {
    session := &models.Session{}
    query := `SELECT id, name, created_at, last_active FROM sessions WHERE id = ?`
    err := s.db.QueryRow(query, id).Scan(&session.ID, &session.Name, &session.CreatedAt, &session.LastActive)
    if err != nil {
        return nil, err
    }
    return session, nil
}