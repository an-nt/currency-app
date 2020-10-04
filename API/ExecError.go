package API

func execError(description string) error {
	return &problem{
		detail: description,
	}
}

type problem struct {
	detail string
}

func (e *problem) Error() string {
	return e.detail
}
