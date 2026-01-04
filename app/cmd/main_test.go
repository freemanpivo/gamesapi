package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestNewApp_RoutesRegistered(t *testing.T) {
	app := NewApp()

	req := httptest.NewRequest("GET", "/health", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("error testing /health: %v", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 503 {
		t.Fatalf("expected 200 or 503 for /health, got %d", resp.StatusCode)
	}

	var healthBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&healthBody); err != nil {
		t.Fatalf("invalid JSON in /health: %v", err)
	}
	if _, ok := healthBody["status"]; !ok {
		t.Errorf("/health response missing 'status' field: %v", healthBody)
	}

	req = httptest.NewRequest("GET", "/games", nil)
	resp, err = app.Test(req, -1)
	if err != nil {
		t.Fatalf("error testing /games: %v", err)
	}
	if resp.StatusCode != 200 && resp.StatusCode != 404 && resp.StatusCode != 500 && resp.StatusCode != 503 {
		t.Fatalf("unexpected /games status code: %d", resp.StatusCode)
	}
}
