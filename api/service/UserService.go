package service

import (
	"smartgw/api/domain"
	"smartgw/api/repository"
)

type (
	UserService interface {
		Valid(username, password string) bool
		Add(user *domain.User) error
		Update(user *domain.User) error
		Delete(username string) error
		Find(username string) (domain.User, error)
		FindAll() ([]domain.User, error)
	}

	userService struct {
		userRepository repository.UserRepository
	}
)

var _ UserService = (*userService)(nil)

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) Valid(username, password string) bool {
	if user, err := u.userRepository.Find(username); err != nil {
		return false
	} else {
		return user.Username == username && user.Password == password
	}
}

func (u *userService) Add(user *domain.User) error {
	return u.userRepository.Save(user)
}

func (u *userService) Update(user *domain.User) error {
	return u.userRepository.Save(user)
}

func (u *userService) Delete(username string) error {
	return u.userRepository.Delete(username)
}

func (u *userService) Find(username string) (domain.User, error) {
	return u.userRepository.Find(username)
}

func (u *userService) FindAll() ([]domain.User, error) {
	return u.userRepository.FindAll()
}
