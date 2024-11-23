package errors

import (
	"dzhordano/132market/services/users/internal/infrastructure/db/postgres"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrInternalFailure = errors.New("internal failure")
	ErrBadRequest      = errors.New("bad request")
	ErrAlreadyExists   = errors.New("already exists")
)

func ToGRPCError(err error) error {
	switch {
	case errors.Is(err, postgres.ErrNotFound):
		return status.Error(codes.NotFound, "not found")

	case errors.Is(err, ErrBadRequest):
		return status.Error(codes.InvalidArgument, "bad request")

	case errors.Is(err, postgres.ErrAlreadyExists):
		return status.Error(codes.AlreadyExists, "already exists")

	default:
		return status.Error(codes.Internal, "internal error")
	}
}
