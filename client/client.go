package main

import (
	"context"
	"fmt"
	"log"

	pb "assignment-totality-corp/proto/userService"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:8080" // Address of the gRPC server
)

// getUserById retrieves a user by their ID.
func RetrieveUserById(client pb.UserServiceClient, id int32) (*pb.UserResponse, error) {
	req := &pb.GetUserRequest{Id: id}
	resp, err := client.RetrieveUserById(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}
	return resp, nil
}

// SearchUsersRequest searches for users based on filters
func SearchUsersRequest(client pb.UserServiceClient, fname, city string, phone int64, minHeight, maxHeight float64, married *bool) (*pb.SearchUsersResponse, error) {
	req := &pb.SearchUsersRequest{
		FullName:  fname,
		City:      city,
		Phone:     phone,
		MinHeight: minHeight,
		MaxHeight: maxHeight,
	}

	if married != nil {
		req.IsMarried = *married
	}

	resp, err := client.SearchUsers(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("error while search users: %v", err)
	}
	return resp, nil
}

// getUsersByIds retrieves multiple users by their IDs.
func RetrieveUsersByIds(client pb.UserServiceClient, ids []int32) (*pb.GetUsersResponse, error) {
	req := &pb.GetUsersRequest{Ids: ids}
	resp, err := client.RetrieveUsersByIds(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("could not get users by IDs: %v", err)
	}
	return resp, nil
}

func main() {

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	client := pb.NewUserServiceClient(conn)
	defer conn.Close()

	// execution for user listing
	filteredUsers, err := SearchUsersRequest(client, "Ben", "", 0, 0, 0, nil)
	if err != nil {
		log.Fatalf("Error searching users: %v", err)
	}
	log.Println("filteredUsers-->", filteredUsers)
	for _, u := range filteredUsers.Users {
		fmt.Printf("ID=%d, Name=%s, City=%s, Phone=%d, Height=%f, Married=%t\n",
			u.Id, u.FullName, u.City, u.Phone, u.Height, u.IsMarried)
	}

	// userDetails, err := RetrieveUserById(client, 2)
	// if err != nil {
	// 	log.Fatalf("Error while getting the details of user: %v", err)
	// }
	// log.Println("userDetails-->", userDetails)

}
