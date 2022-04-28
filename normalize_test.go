package multierror

import (
	"errors"
	"reflect"
	"testing"
)

func TestNormalizeNil(t *testing.T) {
	if Normalize(nil) == nil {
		t.Fatalf("bad: %#v", Normalize(nil))
	}
}

func TestNormalizeNativeError(t *testing.T) {
	err := errors.New(`test`)
	multi := Normalize(err)

	isComparable := reflect.TypeOf(multi).Comparable()
	if isComparable && err == multi {
		t.Fatal("should be true")
	}

	errStr := `1 error occurred:
	* test

`
	if errStr != multi.Error() {
		t.Fatal("should be true")
	}

	expected := New(err)
	if expected.Error() != multi.Error() {
		t.Fatalf("should be true: `%s` = `%s`", expected.Error(), multi.Error())
	}
	if expected.Count() != multi.Count() {
		t.Fatalf("should be true: `%d` = `%d`", expected.Count(), multi.Count())
	}

	if !multi.HasErrors() {
		t.Fatal("should be true")
	}
}

func TestNormalizePtr(t *testing.T) {
	var e *Collection
	multi := Normalize(e)

	if multi == nil {
		t.Fatal("should not be nil")
	}

	if reflect.ValueOf(multi).IsNil() {
		t.Fatal("should not be nil")
	}

	if multi.Count() != 0 {
		t.Fatal("should be 0")
	}

	if multi.HasErrors() {
		t.Fatal("should be true")
	}

}
