package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Arneproductions/DISYS-Exercise-1/microservices/pkg/user"
	"github.com/Arneproductions/DISYS-Exercise-1/microservices/pkg/user/endpoints"
	"github.com/go-kit/kit/log"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHttpHandler(e endpoints.Endpoints, logger log.Logger) *mux.Router {
	options := []httpTransport.ServerOption{
		httpTransport.ServerErrorLogger(logger),
		httpTransport.ServerErrorEncoder(encodeErrorResponse),
	}

	r := mux.NewRouter()
	s := r.PathPrefix("/v1/user").Subrouter()

	s.Methods("GET").Path("").Handler(httpTransport.NewServer(
		e.GetUsersEndpoint,
		decodeGetUsersRequest,
		encodeResponse,
		options...,
	))

	s.Methods("GET").Path("/{userId}").Handler(httpTransport.NewServer(
		e.GetUserEndpoint,
		decodeGetUserRequest,
		encodeResponse,
		options...,
	))

	s.Methods("GET").Path("/byIds").Handler(httpTransport.NewServer(
		e.GetUsersInEndpoint,
		decodeGetUsersInRequest,
		encodeResponse,
		options...,
	))

	s.Methods("POST").Path("").Handler(httpTransport.NewServer(
		e.CreateUserEndpoint,
		decodeCreateUserRequest,
		encodeResponse,
		options...,
	))

	s.Methods("PUT").Path("/{userId}").Handler(httpTransport.NewServer(
		e.UpdateUserEndpoint,
		decodeUpdateUserRequest,
		encodeResponse,
		options...,
	))

	s.Methods("DELETE").Path("/{userId}").Handler(httpTransport.NewServer(
		e.DeleteUserEndpoint,
		decodeDeleteUserRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeGetUsersRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func decodeGetUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var (
		request      endpoints.GetUserRequest
		err          error
		userIdString = mux.Vars(r)["userId"]
	)

	if request.UserId, err = strconv.ParseUint(userIdString, 10, 64); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeGetUsersInRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.GetUsersInRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeCreateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request user.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeUpdateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var (
		request      endpoints.UpdateUserRequest
		err          error
		userIdString = mux.Vars(r)["userId"]
	)

	if request.UserId, err = strconv.ParseUint(userIdString, 10, 64); err != nil {
		return nil, err
	}

	if err = json.NewDecoder(r.Body).Decode(&request.User); err != nil {
		return nil, err
	}

	return request, nil
}

func decodeDeleteUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var (
		request      endpoints.DeleteUserRequest
		err          error
		userIdString = mux.Vars(r)["userId"]
	)

	if request.UserId, err = strconv.ParseUint(userIdString, 10, 64); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeErrorResponse(ctx, e, w)
		return nil
	}

	return json.NewEncoder(w).Encode(response)
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(getErrorCode(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

// Useful when using a bunch of defaults
func getErrorCode(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}
