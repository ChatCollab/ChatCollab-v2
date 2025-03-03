# ChatCollab API

A Golang CRUD API for managing chat sessions, agents, and messages with SQLite storage.

## Data Model

- **Session**: Represents a chat session with a last heartbeat timestamp and relationships to agents and messages.
- **Agent**: Represents an AI agent with properties like name, role, prompt, model, online status, and reasoning log.
- **Message**: Represents a chat message with content, creation timestamp, and relationships to the author (agent) and session.

## Getting Started

### Prerequisites

- Go 1.24 or higher
- SQLite3

### Installation

1. Clone the repository:
```bash
git clone https://github.com/chatcollab/ChatCollab-v2.git
cd ChatCollab-v2
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`.

## API Endpoints

### Agents

- `GET /api/agents` - Get all agents
- `GET /api/agents/:id` - Get agent by ID
- `POST /api/agents` - Create a new agent
- `PUT /api/agents/:id` - Update an agent
- `DELETE /api/agents/:id` - Delete an agent

### Sessions

- `GET /api/sessions` - Get all sessions
- `GET /api/sessions/:id` - Get session by ID
- `POST /api/sessions` - Create a new session
- `PUT /api/sessions/:id/heartbeat` - Update session heartbeat
- `DELETE /api/sessions/:id` - Delete a session
- `GET /api/sessions/:id/agents` - Get all agents for a session
- `GET /api/sessions/:id/messages` - Get all messages for a session

### Messages

- `GET /api/messages/:id` - Get message by ID
- `POST /api/messages` - Create a new message
- `PUT /api/messages/:id` - Update a message
- `DELETE /api/messages/:id` - Delete a message

## Example Usage

### Create a Session

```bash
curl -X POST http://localhost:8080/api/sessions
```

### Create an Agent

```bash
curl -X POST http://localhost:8080/api/agents \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Assistant",
    "role": "helper",
    "prompt": "You are a helpful assistant",
    "model": "gpt-4",
    "is_online": true,
    "session_id": "YOUR_SESSION_ID"
  }'
```

### Create a Message

```bash
curl -X POST http://localhost:8080/api/messages \
  -H "Content-Type: application/json" \
  -d '{
    "content": "Hello, how can I help you today?",
    "author_id": "YOUR_AGENT_ID",
    "session_id": "YOUR_SESSION_ID"
  }'
```

## Database

The application uses SQLite for data storage. The database file is created at `./data/chatcollab.db`.

## License

This project is licensed under the MIT License - see the LICENSE file for details.