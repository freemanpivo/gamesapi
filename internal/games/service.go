package games

// Service orchestrates business logic
type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetGames(name, platform, gender, subGender string) ([]Game, error) {
	// If no filters provided, return all
	if name == "" && platform == "" && gender == "" && subGender == "" {
		return s.repo.ListAll()
	}
	return s.repo.Filtered(name, platform, gender, subGender)
}
