package grpc

import (
	"context"

	"github.com/dzhordano/132market/services/sso/internal/application/interfaces"
	"github.com/dzhordano/132market/services/sso/pkg/pb/sso_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthenticationController struct {
	authenticationService interfaces.AuthenticationService
	sso_v1.UnimplementedAuthenticationV1Server
}

func NewAuthenticationController(as interfaces.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		authenticationService: as,
	}
}

func (c *AuthenticationController) Register(ctx context.Context, request *sso_v1.RegisterRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (c *AuthenticationController) Login(ctx context.Context, request *sso_v1.LoginRequest) (*sso_v1.LoginResponse, error) {
	return &sso_v1.LoginResponse{}, nil
}

func (c *AuthenticationController) RefreshTokens(ctx context.Context, request *sso_v1.RefreshTokensRequest) (*sso_v1.RefreshTokensResponse, error) {
	return &sso_v1.RefreshTokensResponse{}, nil
}
