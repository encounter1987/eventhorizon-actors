package taskboard

import "testing"

// This test will fail if docker is not running or if the MySQL server is not reachable.
func TestInitMySQLClient(t *testing.T) {
	// Get the MySQL client instance
	client := GetMySQLClient()

	// Check if the client is not nil
	if client == nil {
		t.Fatal("Expected MySQL client to be initialized, but got nil")
	}

	// Test the connection
	if err := client.Ping(); err != nil {
		t.Fatalf("Expected MySQL client to be able to ping the database, but got error: %v", err)
	}
}
