package usecase

import "github.com/DemchukM/api-gateway/internal/domain"

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) GetUserById(id int) (*domain.User, error) {
	return uc.repo.GetById(id)
}

func (uc *UserUseCase) CreateUser(user *domain.User) (int, error) {
	return uc.repo.Create(user)
}
