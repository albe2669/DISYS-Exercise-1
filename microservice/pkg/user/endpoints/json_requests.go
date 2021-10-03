package endpoints

import (
	"github.com/Arneproductions/DISYS-Exercise-1/microservices/pkg/user"
)

type GetUserRequest struct {
	UserId uint64 `json:"id"`
}

type GetUserResonse struct {
	User user.User `json:"data,omitempty"`
	Err  string    `json:"err,omitempty"`
}

type GetUsersRequest struct {
}

type GetUsersResponse struct {
	User []user.User `json:"data,omitempty"`
	Err  string      `json:"err,omitempty"`
}

type GetUsersInRequest struct {
	UserIds []uint64 `json:"ids"`
}

type GetUsersInResponse struct {
	User []user.User `json:"data,omitempty"`
	Err  string      `json:"err,omitempty"`
}

type CreateUserResponse struct {
	User user.User `json:"data,omitempty"`
	Err  string    `json:"err,omitempty"`
}

type UpdateUserRequest struct {
	UserId uint64
	User   user.CreateUser
}

type UpdateUserResponse struct {
	User user.User `json:"data,omitempty"`
	Err  string    `json:"err,omitempty"`
}

type DeleteUserRequest struct {
	UserId uint64 `json:"id"`
}

type DeleteUserResonse struct {
	Msg string `json:"msg,omitempty"`
	Err string `json:"err,omitempty"`
}
