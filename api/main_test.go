package api

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	db "go-bday-reminders/db/sqlc"
	"go-bday-reminders/utils"
	"os"
	"testing"
	"time"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomStringAlphabet(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
