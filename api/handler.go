package api

import (
	"errors"
	"fmt"

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

// NewServer ...
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
	var uid int64
	uid = in.Id

	u, err := s.db.SelectUser(uid)
	if err != nil {
		return nil, err
	}

	return &User{Id: u.ID, Name: u.Name, Utype: u.UserType, Balance: u.Balance}, nil
}

// ListUsers responce list of Users
func (s *Server) ListUsers(in *User, stream DoUsers_ListUsersServer) error {
	u, _ := s.db.ListUsers()

	for _, k := range u {
		if err := stream.Send(&User{Id: k.ID, Name: k.Name, Utype: k.UserType, Balance: k.Balance}); err != nil {
			return err
		}
	}
	return nil
}

// UpdateUser responce updating of User
func (s *Server) UpdateUser(ctx context.Context, in *User) (*User, error) {
	var u model.User
	u.ID = in.Id
	u.Name = in.Name
	u.UserType = in.Utype
	u.Balance = int32(in.Balance)

	upd, err := s.db.UpdateUser(u)
	if err != nil {
		fmt.Println("ErrMySQL:", u.UserType, u.Name, u.ID, err)
		return nil, err
	}

	return &User{Id: upd.ID, Name: upd.Name, Utype: upd.UserType, Balance: upd.Balance}, nil
}

// DeleteUser ...
func (s *Server) DeleteUser(ctx context.Context, in *UserId) (*User, error) {
	uid := in.Id
	err := s.db.DeleteUser(uid)
	if err != nil {
		return nil, err
	}
	return &User{Id: uid}, nil
}

// CreateTask generates responce id from DB
func (s *Server) CreateTask(ctx context.Context, in *Task) (*TaskId, error) {
	var u model.Task
	u.Description = in.Description
	u.Price = int32(in.Price)
	u.Status = in.Status

	uid, err := s.db.CreateTask(u)
	if err != nil {
		return nil, err
	}

	return &TaskId{Id: uid.ID}, nil
}

// SelectTask responce selected Task
func (s *Server) SelectTask(ctx context.Context, in *TaskId) (*Task, error) {

	return nil, nil
}

// ListTasks responce list of Tasks
func (s *Server) ListTasks(ctx context.Context, in *Task) (*Task, error) {

	return nil, nil
	//return &Task{Id: uid, Name: name, Utype: utype, Balance: balance}, nil
}

// UpdateTask responce updating of Task
func (s *Server) UpdateTask(ctx context.Context, in *Task) (*Task, error) {

	return nil, nil
	//return &Task{Id: uid, Name: name, Utype: utype, Balance: balance}, nil
}
