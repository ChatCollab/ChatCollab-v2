package services

import (
	"time"

	"github.com/chatcollab/chatcollab/models"
	"github.com/chatcollab/chatcollab/repositories"
)

// SessionService handles business logic for sessions
type SessionService struct {
	repo repositories.SessionRepository
}

// NewSessionService creates a new SessionService
func NewSessionService() *SessionService {
	return &SessionService{
		repo: repositories.SessionRepository{},
	}
}

// CreateSession creates a new session
func (s *SessionService) CreateSession() (*models.Session, error) {
	session := models.NewSession()
	err := s.repo.Create(session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// GetSession retrieves a session by ID
func (s *SessionService) GetSession(id string) (*models.Session, error) {
	return s.repo.GetByID(id)
}

// UpdateHeartbeat updates a session's heartbeat
func (s *SessionService) UpdateHeartbeat(id string) error {
	session, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	
	session.UpdateHeartbeat()
	return s.repo.Update(session)
}

// DeleteSession deletes a session
func (s *SessionService) DeleteSession(id string) error {
	return s.repo.Delete(id)
}

// ListSessions lists all sessions
func (s *SessionService) ListSessions() ([]*models.Session, error) {
	return s.repo.ListAll()
}

// ListActiveSessions lists all active sessions
func (s *SessionService) ListActiveSessions(timeout time.Duration) ([]*models.Session, error) {
	return s.repo.GetActiveSessions(timeout)
}

// IsSessionActive checks if a session is active
func (s *SessionService) IsSessionActive(id string, timeout time.Duration) (bool, error) {
	session, err := s.repo.GetByID(id)
	if err != nil {
		return false, err
	}
	
	return session.IsActive(timeout), nil
}