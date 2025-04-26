package errs

type Errs struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status"`
}

func (e Errs) Error() string {
	return e.Message
}
func (e Errs) Status() int {
	return e.StatusCode
}

func NewBadRequestError(message string) Errs {
	return Errs{Message: message, StatusCode: 400}
}

func NewUnauthorizedError(message string) Errs {
	return Errs{Message: message, StatusCode: 401}
}

func NewForbiddenError(message string) Errs {
	return Errs{Message: message, StatusCode: 403}
}

func NewNotFoundError(message string) Errs {
	return Errs{Message: message, StatusCode: 404}
}

func NewConflictError(message string) Errs {
	return Errs{Message: message, StatusCode: 409}
}

func NewInternalServerError(message string) Errs {
	return Errs{Message: message, StatusCode: 500}
}
