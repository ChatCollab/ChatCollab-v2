package services

import (
	"time"

	"github.com/chatcollab/chatcollab/models"
	"github.com/chatcollab/chatcollab/repositories"
)

// MessageService handles business logic for messages
type MessageService struct {
	repo repositories.MessageRepository
}

// NewMessageService creates a new MessageService
func NewMessageService() *MessageService {
	return &MessageService{
		repo: repositories.MessageRepository{},
	}
}

// CreateMessage creates a new message
func (s *MessageService) CreateMessage(content, agentID, sessionID string) (*models.Message, error) {
	message := models.NewMessage(content, agentID, sessionID)
	err := s.repo.Create(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

// GetMessage retrieves a message by ID
func (s *MessageService) GetMessage(id string) (*models.Message, error) {
	return s.repo.GetByID(id)
}

// UpdateMessage updates a message's content
func (s *MessageService) UpdateMessage(id, content string) error {
	message, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	
	message.Content = content
	return s.repo.Update(message)
}

// DeleteMessage deletes a message
func (s *MessageService) DeleteMessage(id string) error {
	return s.repo.Delete(id)
}

// GetSessionMessages retrieves all messages for a session
func (s *MessageService) GetSessionMessages(sessionID string) ([]*models.Message, error) {
	return s.repo.GetBySessionID(sessionID)
}

// GetAgentMessages retrieves all messages for an agent
func (s *MessageService) GetAgentMessages(agentID string) ([]*models.Message, error) {
	return s.repo.GetByAgentID(agentID)
}

// GetNewMessages retrieves all messages after a specific time
func (s *MessageService) GetNewMessages(sessionID string, after time.Time) ([]*models.Message, error) {
	return s.repo.GetMessagesAfter(sessionID, after)
}