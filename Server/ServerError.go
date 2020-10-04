package Server

type problem struct {
	detail string
}

func serverError(description string) error {
	return &problem{description}
}

func (e *problem) Error() string {
	return e.detail
}
