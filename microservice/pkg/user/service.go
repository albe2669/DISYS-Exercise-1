package user

import (
	"context"
)

type UserService interface {
	GetUser(ctx context.Context, userId uint64) (User, error)
	GetUsers(ctx context.Context) ([]User, error)
	GetUsersIn(ctx context.Context, userIds []uint64) ([]User, error)
	CreateUser(ctx context.Context, user CreateUser) (User, error)
	UpdateUser(ctx context.Context, userId uint64, user CreateUser) (User, error)
	DeleteUser(ctx context.Context, userId uint64) error // FIXME: Not quite sure about what this should return
}
