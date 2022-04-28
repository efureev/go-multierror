package multierror

import "fmt"

type ErrorCollection interface {
	Errors() []error
	SetErrors([]error)
}

type Collection struct {
	errors    []error
	Formatter FormatterFn
}

func (c *Collection) Error() string {
	fn := c.Formatter
	if fn == nil {
		fn = DefaultFormatFunc
	}

	return fn(c.errors)
}

func (c *Collection) GoString() string {
	return fmt.Sprintf("*%#v", *c)
}

func (c *Collection) Append(errs ...error) {
	c.errors = append(c.errors, errs...)
}

func (c *Collection) SetFormatter(fn FormatterFn) *Collection {
	c.Formatter = fn
	return c
}

func (c *Collection) SetErrors(list []error) {
	c.errors = list
}

func (c *Collection) UpdateError(idx int, err error) {
	c.errors[idx] = err
}

func (c *Collection) Errors() []error {
	return c.errors
}

func (c Collection) Count() int {
	return len(c.errors)
}

func (c Collection) HasErrors() bool {
	return c.Count() > 0
}

func New(err ...error) *Collection {
	col := &Collection{}
	col.Append(err...)

	return col
}
