package app

import (
	"fmt"
	"net"

	grpcapi "github.com/PeremyshlevPR/AuthService/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type App struct {
	config Config

	grpcServer *grpc.Server
	listener   net.Listener
}

func NewApp(config Config) *App {
	server, listener, err := initServer(config.GRPCServer)
	if err != nil {
		panic("failed to initialize gRPC server: " + err.Error())
	}
	authServer := grpcapi.NewAuthServiceServer()
	grpcapi.RegisterService(server, authServer)

	return &App{
		config:     config,
		grpcServer: server,
		listener:   listener,
	}
}

func initServer(cfg grpcapi.GRPCServerConfig) (*grpc.Server, net.Listener, error) {
	var opts []grpc.ServerOption

	if cfg.MaxRecvMsgSize > 0 {
		opts = append(opts, grpc.MaxRecvMsgSize(cfg.MaxRecvMsgSize))
	}
	if cfg.MaxSendMsgSize > 0 {
		opts = append(opts, grpc.MaxSendMsgSize(cfg.MaxSendMsgSize))
	}

	if cfg.EnableTLS {
		creds, err := credentials.NewServerTLSFromFile(cfg.TLSCertFile, cfg.TLSKeyFile)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to load TLS credentials: %w", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	server := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", cfg.Address())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to listen: %w", err)
	}

	return server, lis, nil
}

func (a *App) Run() error {
	return a.grpcServer.Serve(a.listener)
}

func (a *App) Stop() {
	a.grpcServer.Stop()
}
