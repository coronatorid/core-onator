package entity

// ApplicationError cover all application error
type ApplicationError struct {
	Err        []error
	HTTPStatus int
}

func (a *ApplicationError) Error() string {
	errString := ""

	for _, e := range a.Err {
		errString = errString + e.Error() + ". "
	}

	return errString
}
