package errors

import (
	"errors"

	"github.com/dzhordano/132market/services/users/internal/infrastructure/db/postgres"

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

// FIXME костыльная тема как-будто, но робит...
func ToGRPCErrors(mainErr error, detailErrs []error) (err error) {
	switch {
	case errors.Is(mainErr, postgres.ErrNotFound):
		err = status.Error(codes.NotFound, "not found:\n"+errors.Join(detailErrs...).Error())

	case errors.Is(mainErr, ErrBadRequest):
		err = status.Error(codes.InvalidArgument, "bad request:\n"+errors.Join(detailErrs...).Error())

	case errors.Is(mainErr, postgres.ErrAlreadyExists):
		err = status.Error(codes.AlreadyExists, "already exists:\n"+errors.Join(detailErrs...).Error())

	default:
		err = status.Error(codes.Internal, "internal error:\n"+errors.Join(detailErrs...).Error())
	}

	return
}
