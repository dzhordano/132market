package errors

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrInternalFailure = errors.New("internal failure")
	ErrBadRequest      = errors.New("bad request")
	// ErrUserNotFound = status.Error(codes.NotFound, "user not found")
	// ErrUnauthorized = status.Error(codes.Unauthenticated, "unauthorized")
	// ErrInternal     = status.Error(codes.Internal, "internal error")
)

func ToGRPCError(err error) error {
	switch {
	case errors.Is(err, ErrNotFound):
		return status.Error(codes.NotFound, "not found")
	case errors.Is(err, ErrBadRequest):
		return status.Error(codes.InvalidArgument, "bad request")
	default:
		return status.Error(codes.Internal, "internal error")
	}
}
