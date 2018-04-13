package db

import (
	"database/sql"
	"os"
	"testing"

	"github.com/go-courses/freelance/config"
	"github.com/go-courses/freelance/model"
	"github.com/stretchr/testify/assert"
)

func TestCRUDLUser(t *testing.T) {
	lastenv := os.Getenv("DATABASE_URL")
	os.Setenv("DATABASE_URL", "postgres://dbuser_f:dbpass_f@localhost:5432/freelance?query")
	c, err := config.GetConfig()
	assert.NoError(t, err)
	m, err := NewPgSQL(c)
	assert.NoError(t, err)
	s, err := m.CreateUser(model.User{
		Name:     "John Doe",
		UserType: "client",
		Balance:  0,
	})
	assert.NoError(t, err)
	assert.Equal(t, s.Balance, 0)

	s.Balance = 100
	s, err = m.UpdateUser(s)
	assert.NoError(t, err)
	assert.Equal(t, s.Balance, 100)
	items, err := m.ListUsers()
	assert.NoError(t, err)
	assert.Equal(t, items[0], s)
	selected, err := m.SelectUser(s.ID)
	assert.NoError(t, err)
	assert.Equal(t, s, selected)
	assert.NoError(t, m.DeleteUser(s.ID))
	s, err = m.SelectUser(s.ID)
	assert.Equal(t, err, sql.ErrNoRows)
	assert.Equal(t, s.ID, int64(0))

	os.Setenv("DATABASE_URL", lastenv)
}
