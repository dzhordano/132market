package grpc

import (
	"context"

	"github.com/dzhordano/132market/services/sso/internal/application/interfaces"
	"github.com/dzhordano/132market/services/sso/pkg/pb/validation_v1"
)

type ValidationController struct {
	tokenValidationService interfaces.TokenValidationService
	validation_v1.UnimplementedValidationV1Server
}

func NewValidationController(tvs interfaces.TokenValidationService) *ValidationController {
	return &ValidationController{
		tokenValidationService: tvs,
	}
}

func (c *ValidationController) ValidateToken(ctx context.Context, request *validation_v1.ValidateTokenRequest) (*validation_v1.ValidateTokenResponse, error) {
	return &validation_v1.ValidateTokenResponse{}, nil
}
