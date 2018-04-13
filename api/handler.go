package api

import (
	"log"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct{}

// CreateUser generates responce id from DB
func (s *Server) CreateUser(ctx context.Context, in *User) (*UserId, error) {
	log.Printf("Recieve message User Id %d", in.Id)
	log.Printf("Recieve message User Name %s", in.Name)

	var uid int32
	uid = 123

	return &UserId{Id: uid}, nil
}

// SelectUser responce selected User
func (s *Server) SelectUser(ctx context.Context, in *UserId) (*User, error) {
	log.Printf("Recieve message User Id %d", in.Id)
	var uid, balance int32
	uid = 555
	name := "Testuser"
	utype := "client"
	balance = 10

	return &User{Id: uid, Name: name, Utype: utype, Balance: balance}, nil
}

// ListUsers responce list of Users
func (s *Server) ListUsers(ctx context.Context, in *User) (*User, error) {
	log.Printf("Recieve message User Id %d", in.Id)
	var uid, balance int32
	uid = 666
	name := "Testuser666"
	utype := "client"
	balance = 66

	return &User{Id: uid, Name: name, Utype: utype, Balance: balance}, nil
}

// UpdateUser responce updating of User
func (s *Server) UpdateUser(ctx context.Context, in *User) (*User, error) {
	log.Printf("Recieve message User Id %d", in.Id)
	var uid, balance int32
	uid = 777
	name := "Testuser777"
	utype := "client"
	balance = 77

	return &User{Id: uid, Name: name, Utype: utype, Balance: balance}, nil
}
