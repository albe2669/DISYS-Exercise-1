package endpoints

import (
	"context"

	"github.com/Arneproductions/DISYS-Exercise-1/microservices/pkg/user"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetUserEndpoint    endpoint.Endpoint
	GetUsersEndpoint   endpoint.Endpoint
	GetUsersInEndpoint endpoint.Endpoint
	CreateUserEndpoint endpoint.Endpoint
	UpdateUserEndpoint endpoint.Endpoint
	DeleteUserEndpoint endpoint.Endpoint
}

func NewEndpoints(svc user.UserService) Endpoints {
	return Endpoints{
		GetUserEndpoint:    makeGetUserEndpoint(svc),
		GetUsersEndpoint:   makeGetUsersEndpoint(svc),
		GetUsersInEndpoint: makeGetUsersInEndpoint(svc),
		CreateUserEndpoint: makeCreateUserEndpoint(svc),
		UpdateUserEndpoint: makeUpdateUserEndpoint(svc),
		DeleteUserEndpoint: makeDeleteUserEndpoint(svc),
	}
}

func makeGetUserEndpoint(svc user.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		retrieved, err := svc.GetUser(ctx, req.UserId)
		if err != nil {
			return GetUserResonse{user.User{}, err.Error()}, err
		}

		return GetUserResonse{retrieved, ""}, err
	}
}

func makeGetUsersEndpoint(svc user.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		retrieved, err := svc.GetUsers(ctx)
		if err != nil {
			return GetUsersResponse{nil, err.Error()}, err
		}

		return GetUsersResponse{retrieved, ""}, err
	}
}

func makeGetUsersInEndpoint(svc user.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUsersInRequest)
		retrieved, err := svc.GetUsersIn(ctx, req.UserIds)
		if err != nil {
			return GetUsersInResponse{nil, err.Error()}, err
		}

		return GetUsersInResponse{retrieved, ""}, err
	}
}

func makeCreateUserEndpoint(svc user.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(user.CreateUser)
		created, err := svc.CreateUser(ctx, req)
		if err != nil {
			return CreateUserResponse{user.User{}, err.Error()}, err
		}

		return CreateUserResponse{created, ""}, err
	}
}

func makeUpdateUserEndpoint(svc user.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		updated, err := svc.UpdateUser(ctx, req.UserId, req.User)
		if err != nil {
			return UpdateUserResponse{user.User{}, err.Error()}, err
		}

		return UpdateUserResponse{updated, ""}, err
	}
}

func makeDeleteUserEndpoint(svc user.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		err := svc.DeleteUser(ctx, req.UserId)
		if err != nil {
			return DeleteUserResonse{"", err.Error()}, err
		}

		return DeleteUserResonse{"User deleted", ""}, err
	}
}
