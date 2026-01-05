package health_test

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"net/http/httptest"

	"github.com/freemanpivo/games-api/internal/health"
	"github.com/gofiber/fiber/v2"
)

func TestHealth_UP_WhenFileExists(t *testing.T) {
	// setup temp dir with data/games_seed.json
	tmp, err := os.MkdirTemp("", "games-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmp)

	dataDir := filepath.Join(tmp, "data")
	if err := os.MkdirAll(dataDir, fs.ModePerm); err != nil {
		t.Fatalf("failed to create data dir: %v", err)
	}
	seedPath := filepath.Join(dataDir, "games_seed.json")
	if err := os.WriteFile(seedPath, []byte(`[{"id":"1","title":"Test Game"}]`), 0o644); err != nil {
		t.Fatalf("failed to write seed file: %v", err)
	}

	// change working dir to tmp so relative path "data/games_seed.json" resolves
	wd, _ := os.Getwd()
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("failed to chdir: %v", err)
	}
	defer os.Chdir(wd)

	app := fiber.New()
	health.RegisterRoutes(app)

	req := httptest.NewRequest("GET", "/health", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if status, ok := body["status"].(string); !ok || status != "UP" {
		t.Fatalf("expected status UP, got %v", body["status"])
	}
}

func TestHealth_DOWN_WhenFileMissing(t *testing.T) {
	// setup temp dir without data/games_seed.json
	tmp, err := os.MkdirTemp("", "games-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmp)

	// ensure no data dir present
	wd, _ := os.Getwd()
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("failed to chdir: %v", err)
	}
	defer os.Chdir(wd)

	app := fiber.New()
	health.RegisterRoutes(app)

	req := httptest.NewRequest("GET", "/health", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode != 503 {
		t.Fatalf("expected status 503, got %d", resp.StatusCode)
	}

	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if status, ok := body["status"].(string); !ok || status != "DOWN" {
		t.Fatalf("expected status DOWN, got %v", body["status"])
	}
	if _, ok := body["error"]; !ok {
		t.Fatalf("expected error field when DOWN")
	}
}
