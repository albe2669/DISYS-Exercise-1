package user

import (
	"context"
)

type userService struct {
	repo UserRepository
}

func NewService(repo UserRepository) userService {
	return userService{repo}
}

func (s userService) GetUser(ctx context.Context, userId uint64) (user User, err error) {
	user, err = s.repo.GetById(userId)

	return user, err
}

func (s userService) GetUsers(ctx context.Context) (users []User, err error) {
	users, err = s.repo.GetAll()

	return users, err
}

func (s userService) GetUsersIn(ctx context.Context, userIds []uint64) (users []User, err error) {
	users, err = s.repo.GetIn(userIds)

	return users, err
}

func (s userService) CreateUser(ctx context.Context, user CreateUser) (User, error) {
	realUser := CreateToReal(user)
	realUser, err := s.repo.Create(realUser)

	return realUser, err
}

func (s userService) UpdateUser(ctx context.Context, userId uint64, user CreateUser) (User, error) {
	realUser := CreateToReal(user)
	realUser.ID = userId
	realUser, err := s.repo.Update(realUser)

	return realUser, err
}

func (s userService) DeleteUser(ctx context.Context, userId uint64) error {
	err := s.repo.Delete(userId)

	return err
}
