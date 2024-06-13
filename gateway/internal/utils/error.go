package utils

import (
	"errors"
	"fmt"
	"github.com/alserov/restate/gateway/internal/log"
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

type ErrorStatus uint

const (
	Internal ErrorStatus = iota
	InvalidData
	NotFound
)
