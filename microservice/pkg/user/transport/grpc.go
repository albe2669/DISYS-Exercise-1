package transport

import (
	"context"
	"errors"
	"time"

	pb "github.com/Arneproductions/DISYS-Exercise-1/microservices/api/v1/pb/user"
	"github.com/Arneproductions/DISYS-Exercise-1/microservices/pkg/user"
	"github.com/Arneproductions/DISYS-Exercise-1/microservices/pkg/user/endpoints"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	// getUser    grpctransport.Handler
	// getUsers   grpctransport.Handler
	// getUsersIn grpctransport.Handler
	createUser grpctransport.Handler
	// updateUser grpctransport.Handler
	// deleteUser grpctransport.Handler
	pb.UnimplementedUserServiceServer
}

func NewGRPCServer(ep endpoints.Endpoints) pb.UserServiceServer {
	return &grpcServer{
		createUser: grpctransport.NewServer(
			ep.CreateUserEndpoint,
			decodeGRPCCreateRequest,
			encodeGRPCCreateResponse,
		),
	}
}

func (g *grpcServer) Create(ctx context.Context, r *pb.CreateUser) (*pb.User, error) {
	_, rep, err := g.createUser.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}

	return rep.(*pb.User), nil
}

func protoCreateUserToLocal(proto *pb.CreateUser) user.CreateUser {
	return user.CreateUser{
		Name: proto.Name,
		Mail: proto.Mail,
		Role: proto.Role,
	}
}

func convertTime(t string) (time.Time, error) {
	return time.Parse(time.RFC3339, t)
}

func localUserToProto(local user.User) *pb.User {
	created := &pb.User{
		Name:      local.Name,
		Mail:      local.Mail,
		Role:      local.Role,
		Id:        local.ID,
		CreatedAt: local.CreatedAt.Format(time.RFC3339),
		UpdatedAt: local.UpdatedAt.Format(time.RFC3339),
	}

	if local.DeletedAt.Time != (time.Time{}) {
		created.DeletedAt = local.DeletedAt.Time.Format(time.RFC3339)
	} else {
		created.DeletedAt = ""
	}

	return created
}

func decodeGRPCCreateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateUser)

	return protoCreateUserToLocal(req), nil
}

func encodeGRPCCreateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	rep := grpcReply.(endpoints.CreateUserResponse)
	if rep.Err != "" {
		return nil, errors.New(rep.Err)
	}

	created := localUserToProto(rep.User)

	return created, nil
}
