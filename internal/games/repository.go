package games

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

// Repository loads data from a JSON file and provides read/filter methods
type Repository struct {
	store []Game
}

// NewRepositoryFromFile loads games from the provided JSON file path
func NewRepositoryFromFile(path string) (*Repository, error) {
	b, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	var list []Game
	if err := json.Unmarshal(b, &list); err != nil {
		return nil, err
	}
	return &Repository{store: list}, nil
}

func (r *Repository) ListAll() ([]Game, error) {
	out := make([]Game, len(r.store))
	copy(out, r.store)
	sortByRatingDesc(out)
	return out, nil
}

// Filtered: name => substring case-insensitive; others exact case-insensitive
func (r *Repository) Filtered(name, platform, gender, subGender string) ([]Game, error) {
	out := make([]Game, 0, len(r.store))
	for _, g := range r.store {
		if name != "" {
			if !strings.Contains(strings.ToLower(g.Name), strings.ToLower(name)) {
				continue
			}
		}
		if platform != "" {
			if !strings.EqualFold(g.Platform, platform) {
				continue
			}
		}
		if gender != "" {
			if !strings.EqualFold(g.Gender, gender) {
				continue
			}
		}
		if subGender != "" {
			if !strings.EqualFold(g.SubGender, subGender) {
				continue
			}
		}
		out = append(out, g)
	}
	// apply default ordering by rating (desc) before returning
	sortByRatingDesc(out)
	return out, nil
}

// sortByRatingDesc sorts the slice in-place: highest rating first
func sortByRatingDesc(games []Game) {
	sort.SliceStable(games, func(i, j int) bool {
		return games[i].Rating > games[j].Rating
	})
}
