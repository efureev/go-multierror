package multierror

func Append(err error, errs ...error) *Collection {

	errCollection := Normalize(err)

	for _, e := range errs {
		switch e := e.(type) {
		case ErrorCollection:
			if e != nil {
				errCollection.Append(e.Errors()...)
			}
		default:
			if e != nil {
				errCollection.Append(e)
			}
		}
	}

	return errCollection
}
