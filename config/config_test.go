package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	os.Setenv("DATABASE_URL", "value")

	c, err := GetConfig()
	assert.NoError(t, err)
	assert.Equal(t, "value", c.DatabaseURL)

}
