package entity

// ApplicationError cover all application error
type ApplicationError struct {
	Err        []error
	HTTPStatus int
}

// ErrorString return error string as an array of string
func (a *ApplicationError) ErrorString() []string {
	appendedErrorString := []string{}

	for _, e := range a.Err {
		appendedErrorString = append(appendedErrorString, e.Error())
	}

	return appendedErrorString
}

func (a *ApplicationError) Error() string {
	errString := ""

	for _, e := range a.Err {
		errString = errString + e.Error() + ". "
	}

	return errString
}
