package common

// StatusError represents an HTTP Status with a corresponding error object.
type StatusError struct {
	Status int
	Err    error
}

// Gets the Error string from a StatusError object.
func (e *StatusError) Error() string {
	return e.Error()
}
