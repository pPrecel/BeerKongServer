package programerrors

import "fmt"

const (
	CodeInternal                 = 1
	CodeNotFound                 = 2
	CodeAlreadyExists            = 3
	CodeWrongInput               = 4
	CodeUpstreamServerCallFailed = 5
	CodeAuthenticationFailed     = 6
)

type Error interface {
	Append(string, ...interface{}) Error
	Code() int
	Error() string
}

type error struct {
	code    int
	message string
}

func errorf(code int, format string, a ...interface{}) Error {
	return error{code: code, message: fmt.Sprintf(format, a...)}
}

func Internal(format string, a ...interface{}) Error {
	return errorf(CodeInternal, format, a...)
}

func NotFound(format string, a ...interface{}) Error {
	return errorf(CodeNotFound, format, a...)
}

func AlreadyExists(format string, a ...interface{}) Error {
	return errorf(CodeAlreadyExists, format, a...)
}

func WrongInput(format string, a ...interface{}) Error {
	return errorf(CodeWrongInput, format, a...)
}

func UpstreamServerCallFailed(format string, a ...interface{}) Error {
	return errorf(CodeUpstreamServerCallFailed, format, a...)
}

func AuthenticationFailed(format string, a ...interface{}) Error {
	return errorf(CodeAuthenticationFailed, format, a...)
}

func (s error) Error() string {
	return s.message
}

func (s error) Code() int {
	return s.code
}

func (s error) Append(additionalFormat string, a ...interface{}) Error {
	format := additionalFormat + ", " + s.message
	return errorf(s.code, format, a...)
}








