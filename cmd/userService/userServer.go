package main

import (
	"log"
	"net"

	"assignment-totality-corp/internal/constants"
	"assignment-totality-corp/internal/database"
	"assignment-totality-corp/internal/server"
	service "assignment-totality-corp/internal/services"
	pb "assignment-totality-corp/proto/userService"

	"google.golang.org/grpc"
)

var (
	Db          database.Database
	GrpcServer  *grpc.Server
	UserService service.InterfaceUserService
	Lis         net.Listener
	err         error
)

func init() {
	Lis, err = net.Listen("tcp", constants.ServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	GrpcServer = grpc.NewServer()
	Db = database.DBCreation()

	// create a new user service
	UserService = service.NewUserService(&Db)

}

func main() {

	pb.RegisterUserServiceServer(GrpcServer, server.NewUserService(UserService))

	log.Printf("Server is running on %s", constants.ServerPort)
	if err := GrpcServer.Serve(Lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
