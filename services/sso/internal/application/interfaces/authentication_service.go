package interfaces

import (
	"context"

	"github.com/dzhordano/132market/services/sso/internal/domain/entities"
)

type AuthenticationService interface {
	Register(ctx context.Context, name, email, password string) error
	Login(ctx context.Context, email, password string) (*entities.Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (*entities.Tokens, error)
}
