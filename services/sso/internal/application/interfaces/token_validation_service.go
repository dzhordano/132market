package interfaces

import "context"

// FIXME token are redundant due to ctx utilization??? (WHEN I MAKE REQUEST INTERCEPTOR WILL VALIDATE AND TOKEN WILL REMAIN IN CTX?)
type TokenValidationService interface {
	ValidateToken(ctx context.Context, token string) (bool, error)
}
