package common

type StatusError struct {
	Status int
	Err    error
}

func (e *StatusError) Error() string {
	return e.Error()
}
