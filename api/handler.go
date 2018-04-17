package api

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"

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

// CreateBilling generates responce id from DB
func (s *Server) CreateBilling(ctx context.Context, in *Billing) (*BillingId, error) {
	var b model.Billing
	b.Sender = in.Sender
	b.Reciever = in.Reciever
	b.Amount = int32(in.Amount)
	b.BillingType = in.Btype
	b.TaskID = in.TaskId
	b.TimeBill = time.Now()

	bid, err := s.db.CreateBilling(b)
	if err != nil {
		return nil, err
	}

	return &BillingId{Id: bid.ID}, nil
}

// SelectBilling responce selected Billing
func (s *Server) SelectBilling(ctx context.Context, in *BillingId) (*Billing, error) {
	var bid int64
	bid = in.Id

	b, err := s.db.SelectBilling(bid)
	if err != nil {
		return nil, err
	}

	btime, _ := ptypes.TimestampProto(b.TimeBill)
	return &Billing{b.ID, b.Sender, b.Reciever, b.Amount, btime, b.TaskID, b.BillingType}, nil
}

// ListBillings responce list of Billings
func (s *Server) ListBillings(in *Billing, stream DoBillings_ListBillingsServer) error {
	u, _ := s.db.ListBillings()

	for _, b := range u {
		btime, _ := ptypes.TimestampProto(b.TimeBill)
		if err := stream.Send(&Billing{b.ID, b.Sender, b.Reciever, b.Amount, btime, b.TaskID, b.BillingType}); err != nil {
			return err
		}
	}
	return nil
}

// UpdateBilling responce updating of Billing
func (s *Server) UpdateBilling(ctx context.Context, in *Billing) (*Billing, error) {
	var b model.Billing
	b.Sender = in.Sender
	b.Reciever = in.Reciever
	b.Amount = int32(in.Amount)
	b.BillingType = in.Btype
	b.TaskID = in.TaskId
	b.TimeBill = time.Now()

	k, err := s.db.UpdateBilling(b)
	if err != nil {
		return nil, err
	}
	ktime, _ := ptypes.TimestampProto(k.TimeBill)
	return &Billing{k.ID, k.Sender, k.Reciever, k.Amount, ktime, k.TaskID, k.BillingType}, nil
}

// DeleteBilling ...
func (s *Server) DeleteBilling(ctx context.Context, in *BillingId) (*Billing, error) {
	bid := in.Id
	err := s.db.DeleteBilling(bid)
	if err != nil {
		return nil, err
	}
	return &Billing{Id: bid}, nil
}
