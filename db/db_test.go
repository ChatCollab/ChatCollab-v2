package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseInitialization(t *testing.T) {
	testDBPath := "./test.db"
	
	// Clean up from previous tests
	_ = os.Remove(testDBPath)
	
	// Initialize the database
	err := Initialize(testDBPath)
	assert.NoError(t, err)
	
	// Check if database connection is working
	err = DB.Ping()
	assert.NoError(t, err)
	
	// Query to check if tables were created
	var tableCount int
	tables := []string{"sessions", "agents", "messages"}
	
	for _, table := range tables {
		query := `SELECT count(name) FROM sqlite_master WHERE type='table' AND name=?`
		row := DB.QueryRow(query, table)
		err = row.Scan(&tableCount)
		assert.NoError(t, err)
		assert.Equal(t, 1, tableCount, "Table "+table+" should exist")
	}
	
	// Close the database connection
	err = Close()
	assert.NoError(t, err)
	
	// Clean up
	_ = os.Remove(testDBPath)
}