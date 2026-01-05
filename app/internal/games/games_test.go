package games_test

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"net/http/httptest"

	"github.com/freemanpivo/games-api/internal/games"
	"github.com/gofiber/fiber/v2"
)

func setupTempWithSeed(t *testing.T, content string) (cleanup func()) {
	t.Helper()

	tmp, err := os.MkdirTemp("", "games-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	dataDir := filepath.Join(tmp, "data")
	if err := os.MkdirAll(dataDir, fs.ModePerm); err != nil {
		os.RemoveAll(tmp)
		t.Fatalf("failed to create data dir: %v", err)
	}

	seedPath := filepath.Join(dataDir, "games_seed.json")
	if err := os.WriteFile(seedPath, []byte(content), 0o644); err != nil {
		os.RemoveAll(tmp)
		t.Fatalf("failed to write seed file: %v", err)
	}

	wd, _ := os.Getwd()
	if err := os.Chdir(tmp); err != nil {
		os.RemoveAll(tmp)
		t.Fatalf("failed to chdir: %v", err)
	}

	return func() {
		_ = os.Chdir(wd)
		os.RemoveAll(tmp)
	}
}

func TestGetGames_ReturnsList_WhenSeedExists(t *testing.T) {
	cleanup := setupTempWithSeed(t, `[
		{"id":"1","title":"Test Game 1"},
		{"id":"2","title":"Test Game 2"}
	]`)
	defer cleanup()

	app := fiber.New()
	games.RegisterRoutes(app)

	req := httptest.NewRequest("GET", "/games", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var raw interface{}
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	switch v := raw.(type) {
	case []interface{}:
		if len(v) != 2 {
			t.Fatalf("expected 2 games, got %d", len(v))
		}
		if first, ok := v[0].(map[string]interface{}); ok {
			if id, _ := first["id"].(string); id != "1" {
				t.Fatalf("expected first id '1', got %v", id)
			}
		} else {
			t.Fatalf("unexpected type for first element")
		}
	case map[string]interface{}:
		found := false
		keys := []string{"data", "games", "items"}
		for _, k := range keys {
			if maybe, ok := v[k]; ok {
				if arr, ok := maybe.([]interface{}); ok {
					found = true
					if len(arr) != 2 {
						t.Fatalf("expected 2 games inside %s, got %d", k, len(arr))
					}
					if first, ok := arr[0].(map[string]interface{}); ok {
						if id, _ := first["id"].(string); id != "1" {
							t.Fatalf("expected first id '1' inside %s, got %v", k, id)
						}
					} else {
						t.Fatalf("unexpected type for first element inside %s", k)
					}
					break
				} else {
					t.Fatalf("found key %s but it's not an array", k)
				}
			}
		}
		if !found {
			t.Fatalf("expected response to contain an array under one of keys %v, got object keys %v", keys, getMapKeys(v))
		}
	default:
		t.Fatalf("unexpected response type: %T", raw)
	}
}

func TestGetGames_ReturnsError_WhenSeedMissing(t *testing.T) {
	tmp, err := os.MkdirTemp("", "games-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmp)

	wd, _ := os.Getwd()
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("failed to chdir: %v", err)
	}
	defer os.Chdir(wd)

	app := fiber.New()
	games.RegisterRoutes(app)

	req := httptest.NewRequest("GET", "/games", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode == 200 {
		t.Fatalf("expected non-200 when seed is missing, got 200")
	}
}

func getMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
