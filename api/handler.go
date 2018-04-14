package api

import (
	"errors"
	"log"

	"github.com/go-courses/freelance/config"
	"github.com/go-courses/freelance/db"
	"github.com/go-courses/freelance/model"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
	db db.DB
	c  *config.FreelanceConfig
}

func NewServer(c *config.FreelanceConfig) (*Server, error) {
	s := &Server{c: c}
	switch c.DbType {
	case "mysql":
		conn, err := db.NewMySQL(c)
		if err != nil {
			return nil, err
		}
		s.db = conn
	case "postgres":
		conn, err := db.NewPgSQL(c)
		if err != nil {
			return nil, err
		}
		s.db = conn
	default:
		return nil, errors.New("unknown database type")
	}
	return s, nil
}

// CreateUser generates responce id from DB
func (s *Server) CreateUser(ctx context.Context, in *User) (*UserId, error) {
	var u model.User
	u.Name = in.Name
	u.UserType = in.Utype
	u.Balance = int32(in.Balance)

	uid, err := s.db.CreateUser(u)
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
