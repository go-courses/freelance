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
	u.Balance = model.Money(in.Balance)
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

	return &User{Id: u.ID, Name: u.Name, Utype: u.UserType, Balance: int32(u.Balance)}, nil
}

// ListUsers responce list of Users
func (s *Server) ListUsers(ctx context.Context, in *User) (*ManyUsers, error) {
	u, _ := s.db.ListUsers()
	mu := &ManyUsers{}
	var ms []*User
	for _, k := range u {
		kl := &User{Id: k.ID, Name: k.Name, Utype: k.UserType, Balance: int32(k.Balance)}
		ms = append(ms, kl)
	}

	mu.Users = ms

	return mu, nil
}

// UpdateUser responce updating of User
func (s *Server) UpdateUser(ctx context.Context, in *User) (*User, error) {
	var u model.User
	u.ID = in.Id
	u.Name = in.Name
	u.UserType = in.Utype
	u.Balance = model.Money(in.Balance)

	upd, err := s.db.UpdateUser(u)
	if err != nil {
		fmt.Println("ErrMySQL:", u.UserType, u.Name, u.ID, err)
		return nil, err
	}

	return &User{Id: upd.ID, Name: upd.Name, Utype: upd.UserType, Balance: int32(upd.Balance)}, nil
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
	u.Price = model.Money(in.Price)
	u.Status = in.Status
	u.Creator = in.Creator
	u.Executor = in.Executor

	uid, err := s.db.CreateTask(u)
	if err != nil {
		return nil, err
	}

	return &TaskId{Id: uid.ID}, nil
}

// SelectTask responce selected Task
func (s *Server) SelectTask(ctx context.Context, in *TaskId) (*Task, error) {
	var uid int64
	uid = in.Id

	u, err := s.db.SelectTask(uid)
	if err != nil {
		return nil, err
	}

	return &Task{Id: u.ID, Description: u.Description, Creator: u.Creator, Executor: u.Executor, Price: int32(u.Price), Status: u.Status}, nil
}

// ListTasks responce list of Tasks
func (s *Server) ListTasks(ctx context.Context, in *Task) (*ManyTasks, error) {
	u, _ := s.db.ListTasks()

	mt := &ManyTasks{}
	var ms []*Task
	for _, k := range u {
		kl := &Task{Id: k.ID, Description: k.Description, Creator: k.Creator, Executor: k.Executor, Price: int32(k.Price), Status: k.Status}
		ms = append(ms, kl)
	}

	mt.Tasks = ms

	return mt, nil
}

// UpdateTask responce updating of Task
func (s *Server) UpdateTask(ctx context.Context, in *Task) (*Task, error) {
	var u model.Task
	u.ID = in.Id
	u.Description = in.Description
	u.Price = model.Money(in.Price)
	u.Status = in.Status

	upd, err := s.db.UpdateTask(u)
	if err != nil {
		return nil, err
	}

	return &Task{Id: upd.ID, Description: upd.Description, Creator: upd.Creator, Executor: upd.Executor, Price: int32(upd.Price), Status: upd.Status}, nil
}

// DeleteTask ...
func (s *Server) DeleteTask(ctx context.Context, in *TaskId) (*Task, error) {
	uid := in.Id
	err := s.db.DeleteTask(uid)
	if err != nil {
		return nil, err
	}
	return &Task{Id: uid}, nil
}

// CreateBilling generates responce id from DB
func (s *Server) CreateBilling(ctx context.Context, in *Billing) (*BillingId, error) {
	var b model.Billing
	b.Sender = in.Sender
	b.Reciever = in.Reciever
	b.Amount = model.Money(in.Amount)
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
	return &Billing{b.ID, b.Sender, b.Reciever, int32(b.Amount), btime, b.TaskID, b.BillingType}, nil
}

// ListBillings responce list of Billings
func (s *Server) ListBillings(ctx context.Context, in *Billing) (*ManyBillings, error) {
	u, _ := s.db.ListBillings()

	mb := &ManyBillings{}
	var ms []*Billing
	for _, b := range u {
		btime, _ := ptypes.TimestampProto(b.TimeBill)
		kl := &Billing{b.ID, b.Sender, b.Reciever, int32(b.Amount), btime, b.TaskID, b.BillingType}
		ms = append(ms, kl)
	}

	mb.Billings = ms

	return mb, nil
}

// UpdateBilling responce updating of Billing
func (s *Server) UpdateBilling(ctx context.Context, in *Billing) (*Billing, error) {
	var b model.Billing
	b.ID = in.Id
	b.Sender = in.Sender
	b.Reciever = in.Reciever
	b.Amount = model.Money(in.Amount)
	b.BillingType = in.Btype
	b.TaskID = in.TaskId
	b.TimeBill = time.Now()

	k, err := s.db.UpdateBilling(b)
	if err != nil {
		fmt.Println("UPDATE ERROR", err)
		return nil, err
	}
	ktime, err := ptypes.TimestampProto(k.TimeBill)
	if err != nil {
		fmt.Println("Error of Time", err)
	}
	return &Billing{k.ID, k.Sender, k.Reciever, int32(k.Amount), ktime, k.TaskID, k.BillingType}, nil
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
