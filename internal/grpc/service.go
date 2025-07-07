package grpc

import (
	"context"

	authservicepb "github.com/PeremyshlevPR/AuthService/proto/authservice/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServiceServer struct {
	authservicepb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer() *AuthServiceServer {
	return &AuthServiceServer{}
}

func RegisterService(grpcServer *grpc.Server, authService *AuthServiceServer) {
	authservicepb.RegisterAuthServiceServer(grpcServer, authService)
}

func (as *AuthServiceServer) Register(context.Context, *authservicepb.RegisterRequest) (*authservicepb.RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func (as *AuthServiceServer) Login(context.Context, *authservicepb.LoginRequest) (*authservicepb.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func (as *AuthServiceServer) Authenticate(context.Context, *authservicepb.AuthenticateRequest) (*authservicepb.AuthenticateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}

func (as *AuthServiceServer) Authorize(context.Context, *authservicepb.AuthorizeRequest) (*authservicepb.AuthorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
