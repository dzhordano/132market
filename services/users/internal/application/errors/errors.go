package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserNotFound = status.Error(codes.NotFound, "user not found")
	ErrUnauthorized = status.Error(codes.Unauthenticated, "unauthorized")
	ErrInternal     = status.Error(codes.Internal, "internal error")
)

func FromRepoErrorToGRPC(err error) error {
	switch err {
	case ErrUserNotFound:
		return ErrUserNotFound
	case ErrUnauthorized:
		return ErrUnauthorized
	default:
		return ErrInternal
	}
}
