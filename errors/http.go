package _errors

type HttpStatusError struct {
	httpCode int
	mess     string
}

func NewHttpCodeError(status int, mess string) HttpStatusError {
	return HttpStatusError{
		httpCode: status,
		mess:     mess,
	}
}

func BadRequest(mess string) HttpStatusError {
	return NewHttpCodeError(400, mess)
}

func (e HttpStatusError) Error() string {
	return e.mess
}

func (e HttpStatusError) HttpStatus() int {
	return e.httpCode
}
