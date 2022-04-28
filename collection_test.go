package multierror

import (
	"errors"
	"testing"
)

func TestError_Impl(t *testing.T) {
	var _ error = new(Collection)
}

func TestErrorError_custom(t *testing.T) {
	list := []error{
		errors.New("foo"),
		errors.New("bar"),
	}

	fn := func(es []error) string {
		return "foo"
	}

	multi := New(list...).SetFormatter(fn)
	if multi.Error() != "foo" {
		t.Fatalf("bad: %s", multi.Error())
	}
}

func TestErrorError_default(t *testing.T) {
	expected := `2 errors occurred:
	* foo
	* bar

`

	multi := New(errors.New("foo"), errors.New("bar"))

	if multi.Error() != expected {
		t.Fatalf("bad: %s", multi.Error())
	}
}

func TestErrorError_Append(t *testing.T) {
	multi := New()
	multi.Append(errors.New("bar"))
	multi.Append(errors.New("foo"))

	expected := `2 errors occurred:
	* bar
	* foo

`
	if multi.Count() != 2 {
		t.Fatal("should be 2")
	}

	if multi.Error() != expected {
		t.Fatalf("bad: %s", multi.Error())
	}
}
