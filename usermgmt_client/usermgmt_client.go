package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Adan1996/go_usermgmt_grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50021"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)
	new_users["Syahdan"] = 43
	new_users["Enden"] = 30
	for name, age := range new_users {
		r, err := client.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf(`Users details:
		Name: %s
		Age: %d
		ID: %d`, r.GetName(), r.GetAge(), r.GetId())
	}
}
