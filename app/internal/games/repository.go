package games

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

type Repository struct {
	store []Game
}

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
			if !strings.EqualFold(g.Genre, gender) {
				continue
			}
		}
		if subGender != "" {
			if !strings.EqualFold(g.SubGenre, subGender) {
				continue
			}
		}
		out = append(out, g)
	}

	sortByRatingDesc(out)
	return out, nil
}

func sortByRatingDesc(games []Game) {
	sort.SliceStable(games, func(i, j int) bool {
		return games[i].Rating > games[j].Rating
	})
}
