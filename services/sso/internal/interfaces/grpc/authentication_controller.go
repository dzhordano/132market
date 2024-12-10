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
	err := c.authenticationService.Register(ctx, request.GetName(), request.GetEmail(), request.GetPassword())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (c *AuthenticationController) Login(ctx context.Context, request *sso_v1.LoginRequest) (*sso_v1.LoginResponse, error) {
	tokens, err := c.authenticationService.Login(ctx, request.GetEmail(), request.GetPassword())
	if err != nil {
		return nil, err
	}

	return &sso_v1.LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (c *AuthenticationController) RefreshTokens(ctx context.Context, request *sso_v1.RefreshTokensRequest) (*sso_v1.RefreshTokensResponse, error) {
	tokens, err := c.authenticationService.RefreshTokens(ctx, request.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &sso_v1.RefreshTokensResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
