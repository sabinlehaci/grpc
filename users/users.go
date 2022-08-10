package users

import (
	"context"
	"sync"

	usersv1 "github.com/sabinlehaci/grpc-test/gen/go/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	usersv1.UnimplementedUserServiceServer
	mutex *sync.Mutex
	users map[string]*usersv1.User
}

func NewUserService() *UserService {
	return &UserService{
		mutex: &sync.Mutex{},
		users: make(map[string]*usersv1.User),
	}
}

func (u *UserService) CreateUser(ctx context.Context, req *usersv1.CreateUserRequest) (*usersv1.CreateUserResponse, error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	if _, ok := u.users[req.GetName()]; ok {
		return nil, status.Errorf(codes.AlreadyExists, "a user named %q already exists", req.GetName())
	}
	u.users[req.GetName()] = &usersv1.User{
		Name: req.GetName(),
		Age:  req.GetAge(),
	}
	return &usersv1.CreateUserResponse{
		User: u.users[req.GetName()],
	}, nil
}

func (u *UserService) GetUser(ctx context.Context, req *usersv1.GetUserRequest) (*usersv1.GetUserResponse, error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	if _, ok := u.users[req.GetName()]; !ok {
		return nil, status.Errorf(codes.NotFound, "a user named %q could not be found", req.GetName())
	}
	return &usersv1.GetUserResponse{
		User: u.users[req.GetName()],
	}, nil
}

