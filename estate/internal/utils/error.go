package utils

import (
	"errors"
	"fmt"
	"github.com/alserov/restate/estate/internal/log"
	"google.golang.org/grpc/codes"
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

func FromError(l log.Logger, in error) (string, codes.Code) {
	var e *err
	if !errors.As(in, &e) {
		return fmt.Sprintf("unknown error: %v", in), codes.Internal
	}

	switch e.status {
	case InvalidData:
		return e.msg, codes.InvalidArgument
	case NotFound:
		return e.msg, codes.NotFound
	case Internal:
		l.Error(in.Error(), nil)
		return "internal error", codes.Internal
	default:
		l.Error(fmt.Sprintf("unknown status: %s", in.Error()), nil)
		return "internal error", codes.Internal
	}
}

type ErrorStatus uint

const (
	Internal ErrorStatus = iota
	InvalidData
	NotFound
)
