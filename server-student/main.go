package main

import (
	"grpc/database"
	"grpc/server"
	"grpc/studentpb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5050")

	if err != nil {
		log.Fatal(err)
	}

	repo, _ := database.NewPostgresRepository("postgres://postgres:postgres@localhost:5455/postgres?sslmode=disable")

	server := server.NewStudentServer(repo)

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}
}
