package exception

type NotFoundError struct {
	Entity string
}

func (e *NotFoundError) Error() string {
	return e.Entity + " doesn't exist"
}
