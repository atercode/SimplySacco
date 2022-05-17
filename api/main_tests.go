package api

import (
	"database/sql"
	"os"
	"testing"
	"time"

	db "github.com/atercode/SimplySacco/db/sqlc"
	"github.com/atercode/SimplySacco/utils"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func newTestStore(t *testing.T, config *utils.Config) *db.Store {
	conn, err := sql.Open(config.DBDriver, config.DBSourceSting)
	require.NoError(t, err)

	store := db.NewStore(conn)
	return &store
}

func newTestServer(t *testing.T) *Server {
	config, err := utils.LoadConfig("../.")
	require.NoError(t, err)
	config.TokenSymmetricKey = gofakeit.DigitN(32)
	config.AccessTokenDuration = time.Minute

	conn, err := sql.Open(config.DBDriver, config.DBSourceSting)
	require.NoError(t, err)

	store := db.NewStore(conn)

	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func newTestServerWithCustomStore(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymmetricKey:   gofakeit.DigitN(32),
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
