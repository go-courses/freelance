package db

import (
	"database/sql"
	"testing"

	"github.com/go-courses/freelance/config"
	"github.com/go-courses/freelance/model"
	"github.com/stretchr/testify/assert"
)

func TestCRUDLUser(t *testing.T) {
	c, err := config.GetConfig()
	assert.NoError(t, err)
	m, err := NewPgSQL(c)
	assert.NoError(t, err)
	s, err := m.CreateUser(model.User{
		Name:     "John Doe",
		UserType: "user",
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

}
