package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	os.Setenv("DATABASE_URL", "value")
	os.Setenv("DB_TYPE", "value2")
	os.Setenv("DO_MIGRATION", "value3")

	c, err := GetConfig()
	assert.NoError(t, err)
	assert.Equal(t, "value", c.DatabaseURL)
	assert.Equal(t, "value2", c.DbType)
	assert.Equal(t, "value3", c.DoMigration)

}
