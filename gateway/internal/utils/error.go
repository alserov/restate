package utils

import (
	"errors"
	"fmt"
	"github.com/alserov/restate/gateway/internal/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type err struct {
	msg    string
	status ErrorStatus
}

func (e *err) Error() string {
	return e.msg
}

func NewError(msg string, status ErrorStatus) error {
	return &err{msg: msg, status: status}
}

func FromError(l log.Logger, in error) (string, int) {
	var e *err
	if !errors.As(in, &e) {
		return fmt.Sprintf("unknown error: %v", in), http.StatusInternalServerError
	}

	switch e.status {
	case InvalidData:
		return e.msg, http.StatusBadRequest
	case NotFound:
		return e.msg, http.StatusNotFound
	case Internal:
		l.Error(in.Error(), nil)
		return "internal error", http.StatusInternalServerError
	default:
		l.Error(fmt.Sprintf("unknown status: %s", in.Error()), nil)
		return "internal error", http.StatusInternalServerError
	}
}

func ExtractStatus(in error) (ErrorStatus, bool) {
	var e *err
	if !errors.As(in, &e) {
		return 0, false
	}

	return e.status, true
}

func FromGRPCError(in error) (error, ErrorStatus) {
	e, _ := status.FromError(in)

	switch e.Code() {
	case codes.InvalidArgument:
		return NewError(e.Message(), InvalidData), InvalidData
	case codes.NotFound:
		return NewError(e.Message(), NotFound), NotFound
	case codes.Internal:
		return NewError(e.Message(), Internal), Internal
	default:
		return NewError(e.Message(), Internal), Internal
	}
}

type ErrorStatus uint

const (
	Internal ErrorStatus = iota
	InvalidData
	NotFound
)
