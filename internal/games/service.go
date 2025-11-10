package games

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetGames(name, platform, genre, subGenre string) ([]Game, error) {
	if name == "" && platform == "" && genre == "" && subGenre == "" {
		return s.repo.ListAll()
	}
	return s.repo.Filtered(name, platform, genre, subGenre)
}
