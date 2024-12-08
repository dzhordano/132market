package interfaces

import "context"

type TokenValidationService interface {
	ValidateToken(ctx context.Context, token string) error
}
