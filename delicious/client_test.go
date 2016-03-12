package delicious

import "testing"

func TestNewClient(t *testing.T) {
	c, err := NewClient("DUMMY_ACCESS_TOKEN")
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}

	if actual, expected := c.Endpoint, apiEndpoint; actual != expected {
		t.Errorf("NewClient Endpoint is %v, want %v", actual, expected)
	}
}

func TestNewClient_EmptyAccessToken(t *testing.T) {
	_, err := NewClient("")
	if err == nil {
		t.Error("Expected error to be returned")
	}
}
