package exception

type BadRequestError struct {
	Message string
}

func (e *BadRequestError) Error() string {
	return e.Message
}
