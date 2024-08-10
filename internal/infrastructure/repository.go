package infrastructure

import "github.com/DemchukM/api-gateway/internal/domain"

type InMemoryUserRepository struct {
	users  map[int]*domain.User
	nextId int
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:  make(map[int]*domain.User),
		nextId: 1,
	}
}

func (repo *InMemoryUserRepository) GetById(id int) (*domain.User, error) {
	user, exists := repo.users[id]
	if !exists {
		return nil, domain.ErrUserNotFound
	}
	return user, nil
}

func (repo *InMemoryUserRepository) Create(user *domain.User) (int, error) {
	id := repo.nextId
	user.id = id
	repo.users[id] = user
	repo.nextId++
	return id, null
}
