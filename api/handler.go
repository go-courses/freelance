package api

import (
	"fmt"
	"log"

	"github.com/go-courses/freelance/config"
	"github.com/go-courses/freelance/db"
	"github.com/go-courses/freelance/model"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct{}

// CreateUser generates responce id from DB
func (s *Server) CreateUser(ctx context.Context, in *User) (*UserId, error) {
	var u model.User
	u.Name = in.Name
	u.UserType = in.Utype
	u.Balance = int32(in.Balance)

	// Read config from system environment
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal(err, "could not get env conf parms")
	}

	// подключение для PostgreSQL or MySQL, расскоментить нужное
	m, err := db.NewMySQL(c)
	//m, err := db.NewPgSQL(c)
	if err != nil {
		fmt.Println(err, " could not create database connection")
	}

	uid, err := m.CreateUser(u)
	if err != nil {
		return nil, err
	}

	return &UserId{Id: uid.ID}, nil
}

// SelectUser responce selected User
func (s *Server) SelectUser(ctx context.Context, in *UserId) (*User, error) {
	log.Printf("Recieve message User Id %d", in.Id)
	var uid int64
	var balance int32
	uid = 555
	name := "Testuser555"
	utype := "client"
	balance = 10

	return &User{Id: uid, Name: name, Utype: utype, Balance: balance}, nil
}

// ListUsers responce list of Users
func (s *Server) ListUsers(ctx context.Context, in *User) (*User, error) {
	log.Printf("Recieve message User Id %d", in.Id)
	var uid int64
	var balance int32
	uid = 666
	name := "Testuser666"
	utype := "client"
	balance = 66

	return &User{Id: uid, Name: name, Utype: utype, Balance: balance}, nil
}

// UpdateUser responce updating of User
func (s *Server) UpdateUser(ctx context.Context, in *User) (*User, error) {
	log.Printf("Recieve message User Id %d", in.Id)
	var uid int64
	var balance int32
	uid = 777
	name := "Testuser777"
	utype := "client"
	balance = 77

	return &User{Id: uid, Name: name, Utype: utype, Balance: balance}, nil
}
