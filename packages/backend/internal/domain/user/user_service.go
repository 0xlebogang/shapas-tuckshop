package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(user *User) (*User, error) {
	return s.repo.CreateUser(user)
}
