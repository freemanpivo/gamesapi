package games

type Game struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	ReleaseDate string  `json:"releaseDate"` // In format YYYY-MM-DD
	Platform    string  `json:"platform"`
	Gender      string  `json:"gender"`
	SubGender   string  `json:"subGender"`
	Rating      float32 `json:"rating"`
}
