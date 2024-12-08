package grpc

import (
	"context"

	"github.com/dzhordano/132market/services/sso/pkg/pb/sso_v1"
)

type ValidationController struct {
	sso_v1.UnimplementedValidationV1Server
}

func NewValidationController() *ValidationController {
	return &ValidationController{}
}

func (c *ValidationController) ValidateToken(ctx context.Context, request *sso_v1.ValidateTokenRequest) (*sso_v1.ValidateTokenResponse, error) {
	return &sso_v1.ValidateTokenResponse{}, nil
}
