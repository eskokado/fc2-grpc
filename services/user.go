package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/eskokado/fc2-grpc/pb"
)

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	mustEmbedUnimplementedUserServiceServer()
// }
// func (c *userServiceClient) AddUserVerbose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVerboseClient, error) {


type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	// Insert - Database
	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.Name,
		Email: req.Email,
	}, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User: &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User: &pb.User{},
	})

	time.Sleep(time.Second * 3)
	stream.Send(&pb.UserResultStream{
		Status: "Use has been inserted",
		User: &pb.User{
			Id: "123",
			Name: req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	time.Sleep(time.Second * 3)
	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id: "123",
			Name: req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	users := []*pb.User{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users {
				User: users, 
			})
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}
		users = append(users, &pb.User{
			Id: req.GetId(),
			Name: req.GetName(),
			Email: req.GetEmail(),
		})
		fmt.Println("Adding", req.GetName())
	}

	return nil
}

func (*UserService) AddUserStremBoth(stream pb.UserService_AddUserStremBothServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving stream from the client: %v", err)
		}
		err = stream.Send(&pb.UserResultStream {
			Status: "Added",
			User: req,
		})
		if err != nil {
			log.Fatalf("Error sending stream to the client: %v", err)
		}
	}
}