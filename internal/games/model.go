package games

// Game represents the domain model for a game
type Game struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	ReleaseDate string  `json:"releaseDate"` // YYYY-MM-DD
	Platform    string  `json:"platform"`
	Gender      string  `json:"gender"`
	SubGender   string  `json:"subGender"`
	Rating      float32 `json:"rating"`
}
