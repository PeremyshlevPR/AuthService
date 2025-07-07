package acceptance

import (
	"net"
	"testing"
	"time"

	"github.com/PeremyshlevPR/AuthService/internal/app"
	"github.com/PeremyshlevPR/AuthService/internal/db"
	grpcapi "github.com/PeremyshlevPR/AuthService/internal/grpc"
	"github.com/stretchr/testify/require"
)

func testConfig() app.Config {
	return app.Config{
		DB: db.PostgresConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "test",
			Password: "test",
			DBName:   "test",
			SSLMode:  "disable",
		},
		GRPCServer: grpcapi.GRPCServerConfig{
			Host:           "127.0.0.1",
			Port:           25003,
			MaxRecvMsgSize: 1 << 20,
			MaxSendMsgSize: 1 << 20,
			Timeout:        1 * time.Second,
			EnableTLS:      false,
		},
	}
}

func TestApp_StartAndStop(t *testing.T) {
	cfg := testConfig()
	a := app.NewApp(cfg)

	go func() {
		_ = a.Run()
	}()
	defer a.Stop()
	time.Sleep(150 * time.Millisecond)

	conn, err := net.DialTimeout("tcp", cfg.GRPCServer.Address(), 200*time.Millisecond)
	require.NoError(t, err)
	_ = conn.Close()
}
