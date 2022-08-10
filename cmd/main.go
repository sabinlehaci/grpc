//Client Server 
package main


import (
	"context"
	"log"

	usersv1 "github.com/sabinlehaci/grpc-test/gen/go/users/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Failed to run: %v", err)
	}
}

func run() error {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure()) 
	if err != nil {
		return err
	}

	cli := usersv1.NewUserServiceClient(conn)

	resp, err := cli.CreateUser(context.Background(), &usersv1.CreateUserRequest{
		Name: "Johan",
		Age:  31,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Println("Code:", st.Code())
			log.Println("Message:", st.Message())
		}
		return err
	}
	log.Println(resp)

	_, err = cli.GetUser(context.Background(), &usersv1.GetUserRequest{
		Name: resp.GetUser().GetName(),
	})
	if err != nil {
		return err
	}

	_, err = cli.GetUser(context.Background(), &usersv1.GetUserRequest{
		Name: "Sabin",
	})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Println("Code:", st.Code())
			log.Println("Message:", st.Message())
		}
		return err
	}

	return nil
}
