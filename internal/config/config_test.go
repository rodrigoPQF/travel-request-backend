package config_test

import (
	"os"
	"testing"

	"github.com/rodrigoPQF/travel-request-backend/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitAppConfig(t *testing.T) {

	t.Run("successfully load .env", func(t *testing.T) {

		os.Setenv("POSTGRES_USER", "testuser")
		os.Setenv("POSTGRES_PASSWORD", "testpass")
		os.Setenv("POSTGRES_DB", "testdb")
		os.Setenv("POSTGRES_HOST", "localhost")
		os.Setenv("PORT", "8080")

		err := config.InitAppConfig()

		require.NoError(t, err)
	})

	t.Run("missing required environment variable", func(t *testing.T) {
		os.Unsetenv("POSTGRES_USER")

		err := config.InitAppConfig()

		require.Error(t, err)
		assert.Contains(t, err.Error(), "error empty environment POSTGRES_USER")
	})
}

func TestGetEnvs(t *testing.T) {
	os.Setenv("POSTGRES_USER", "testuser")
	os.Setenv("POSTGRES_PASSWORD", "testpass")
	os.Setenv("POSTGRES_DB", "testdb")
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("PORT", "8080")

	envs := config.GetEnvs()

	assert.Equal(t, "testuser", envs.POSTGRES_USER)
	assert.Equal(t, "testpass", envs.POSTGRES_PASSWORD)
	assert.Equal(t, "testdb", envs.POSTGRES_DB)
	assert.Equal(t, "localhost", envs.POSTGRES_HOST)
	assert.Equal(t, "8080", envs.PORT)
}

func TestMain(m *testing.M) {
	os.Clearenv()
	exitVal := m.Run()

	os.Clearenv()
	os.Exit(exitVal)
}
