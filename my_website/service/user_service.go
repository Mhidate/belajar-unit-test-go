package service

import "github.com/Mhidate/belajar-unit-test-go/my_website/repository"

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) GetUserName(id int) (string, error) {
	user, err := s.Repo.GetByID(id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}
