package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/Adan1996/go_usermgmt_grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	port = ":50021"
)

type UserManagementServer struct {
	// implementasi grpc server
	pb.UnimplementedUserManagementServer
}

// buat fungsi receiver
func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName()) // receiver nama
	var user_id int32 = int32(rand.Intn(1000))
	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserManagementServer(server, &UserManagementServer{})
	log.Printf("server listening at: %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
