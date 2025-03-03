package services

import (
	"github.com/chatcollab/chatcollab/models"
	"github.com/chatcollab/chatcollab/repositories"
)

// AgentService handles business logic for agents
type AgentService struct {
	repo repositories.AgentRepository
}

// NewAgentService creates a new AgentService
func NewAgentService() *AgentService {
	return &AgentService{
		repo: repositories.AgentRepository{},
	}
}

// CreateAgent creates a new agent
func (s *AgentService) CreateAgent(name, role, prompt, model, sessionID string) (*models.Agent, error) {
	agent := models.NewAgent(name, role, prompt, model, sessionID)
	err := s.repo.Create(agent)
	if err != nil {
		return nil, err
	}
	return agent, nil
}

// GetAgent retrieves an agent by ID
func (s *AgentService) GetAgent(id string) (*models.Agent, error) {
	return s.repo.GetByID(id)
}

// UpdateAgent updates an agent
func (s *AgentService) UpdateAgent(agent *models.Agent) error {
	return s.repo.Update(agent)
}

// DeleteAgent deletes an agent
func (s *AgentService) DeleteAgent(id string) error {
	return s.repo.Delete(id)
}

// ListAgents lists all agents
func (s *AgentService) ListAgents() ([]*models.Agent, error) {
	return s.repo.ListAll()
}

// ListSessionAgents lists all agents for a session
func (s *AgentService) ListSessionAgents(sessionID string) ([]*models.Agent, error) {
	return s.repo.GetBySessionID(sessionID)
}

// SetAgentOnlineStatus updates an agent's online status
func (s *AgentService) SetAgentOnlineStatus(id string, isOnline bool) error {
	agent, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	
	agent.SetOnline(isOnline)
	return s.repo.Update(agent)
}

// AppendAgentReasoningLog adds to an agent's reasoning log
func (s *AgentService) AppendAgentReasoningLog(id string, log string) error {
	agent, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	
	agent.AppendReasoningLog(log)
	return s.repo.Update(agent)
}