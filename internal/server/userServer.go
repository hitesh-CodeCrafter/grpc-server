package server

import (
	"context"
	"log"
	"strings"

	model "assignment-totality-corp/internal/models"
	service "assignment-totality-corp/internal/services"
	pb "assignment-totality-corp/proto/userService"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService struct {
	userService service.InterfaceUserService
	pb.UnimplementedUserServiceServer
}

func NewUserService(us service.InterfaceUserService) pb.UserServiceServer {
	return &userService{
		userService: us,
	}
}

func (sv *userService) RetrieveUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, err := sv.userService.GetUserById(req.Id)
	if err != nil {
		// Check the type of error and return appropriate gRPC status code
		if strings.ToLower(err.Error()) == "user not found" {
			return nil, status.Errorf(codes.NotFound, "user with ID %s not found", string(req.Id))
		}
		// Return internal server error for other cases
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return &pb.UserResponse{
		Id:        user.ID,
		FullName:  user.FullName,
		City:      user.City,
		Phone:     user.Phone,
		Height:    user.Height,
		IsMarried: user.Married,
	}, nil
}

func (sv *userService) RetrieveUsersByIds(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {

	users, err := sv.userService.GetUserByIds(req.Ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	var usersRes []*pb.UserResponse
	for _, user := range users {
		usersRes = append(usersRes, &pb.UserResponse{
			Id:        user.ID,
			FullName:  user.FullName,
			City:      user.City,
			Phone:     user.Phone,
			Height:    user.Height,
			IsMarried: user.Married,
		})
	}
	return &pb.GetUsersResponse{Users: usersRes}, nil
}

func (sv *userService) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {

	SearchUsersRequest := model.SearchUsersRequest{
		Fname:     req.FullName,
		City:      req.City,
		Phone:     req.Phone,
		MinHeight: req.MinHeight,
		MaxHeight: req.MaxHeight,
	}

	log.Println(SearchUsersRequest, "-->")

	// SearchUsersRequest.Married = &req.IsMarried

	users, err := sv.userService.SearchUsers(SearchUsersRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	var usersRes []*pb.UserResponse
	for _, user := range users {
		usersRes = append(usersRes, &pb.UserResponse{
			Id:        user.ID,
			FullName:  user.FullName,
			City:      user.City,
			Phone:     user.Phone,
			Height:    user.Height,
			IsMarried: user.Married,
		})
	}

	return &pb.SearchUsersResponse{Users: usersRes}, nil
}
