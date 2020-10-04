package Database

type problem struct {
	detail string
}

func dbError(description string) error {
	return &problem{description}
}

func (e *problem) Error() string {
	return e.detail
}
