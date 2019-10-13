package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	todopb "github.com/istsh/todo-server/app/proto/v1/todo"
	userpb "github.com/istsh/todo-server/app/proto/v1/user"
	"github.com/istsh/todo-server/app/repository"
	"github.com/istsh/todo-server/app/service/todoservice"
	"github.com/istsh/todo-server/app/service/userservice"
)

func main() {
	// TODO: port should be configurable.
	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	// TODO: implement repository
	server := NewServer(nil)
	server.Serve(listenPort)
}

func NewServer(r repository.Repository) *grpc.Server {
	var ops []grpc.ServerOption
	//ops = append(ops, grpc.UnaryInterceptor(auth.AuthenticationUnaryInterceptor()))
	s := grpc.NewServer(ops...)

	todoService := todoservice.New(r)
	todopb.RegisterTodoServiceServer(s, todoService)

	userService := userservice.New(r)
	userpb.RegisterUserServiceServer(s, userService)

	return s
}
