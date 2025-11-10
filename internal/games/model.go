package games

type Game struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	ReleaseDate string  `json:"releaseDate"` // In format YYYY-MM-DD
	Platform    string  `json:"platform"`
	Genre       string  `json:"genre"`
	SubGenre    string  `json:"subGenre"`
	Rating      float32 `json:"rating"`
}
