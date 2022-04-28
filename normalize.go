package multierror

func Normalize(err error) *Collection {
	if err == nil {
		return new(Collection)
	}
	e, ok := err.(*Collection)
	if ok {
		if e != nil {
			return e
		}

		return new(Collection)
	}

	if e, ok := err.(ErrorCollection); ok {
		return New(e.Errors()...)
	}

	return New(err)
}
