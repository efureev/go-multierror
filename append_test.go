package multierror

import (
	"errors"
	"testing"
)

func TestAppend_Error(t *testing.T) {
	original := &Collection{
		errors: []error{errors.New("foo")},
	}

	result := Append(original, errors.New("bar"))
	if len(result.Errors()) != 2 {
		t.Fatalf("wrong len: %d", len(result.Errors()))
	}

	original = &Collection{}
	result = Append(original, errors.New("bar"))
	if len(result.Errors()) != 1 {
		t.Fatalf("wrong len: %d", len(result.Errors()))
	}

	// Test when a typed nil is passed
	var e *Collection
	result = Append(e, errors.New("baz"))
	if len(result.Errors()) != 1 {
		t.Fatalf("wrong len: %d", len(result.Errors()))
	}

	// Test flattening
	original = &Collection{
		errors: []error{errors.New("foo")},
	}

	result = Append(original, Append(nil, errors.New("foo"), errors.New("bar")))
	if len(result.Errors()) != 3 {
		t.Fatalf("wrong len: %d", len(result.Errors()))
	}
}
