package dom

type Error struct {
	code  int
	error string
}

func NewError(code int, error string) Error {
	return Error{
		code:  code,
		error: error,
	}
}

func (e Error) Error() string {
	return e.error
}

func (e Error) Code() int {
	return e.code
}