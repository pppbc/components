package errors

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestError(t *testing.T) {
	var base *Error
	err := New(http.StatusBadRequest, "message")
	err2 := New(http.StatusBadRequest, "message")
	err.WithMetadata(map[string]string{
		"foo": "bar",
	})
	werr := fmt.Errorf("wrap %w", err)

	if errors.Is(err, new(Error)) {
		t.Errorf("should not be equal: %v", err)
	}
	if !errors.Is(werr, err) {
		t.Errorf("should be equal: %v", err)
	}
	if !errors.Is(werr, err2) {
		t.Errorf("should be equal: %v", err)
	}

	if !errors.As(err, &base) {
		t.Errorf("should be matchs: %v", err)
	}
	if !IsBadRequest(err) {
		t.Errorf("should be matchs: %v", err)
	}

	if err.MetaData["foo"] != "bar" {
		t.Error("not expected metadata")
	}
}
