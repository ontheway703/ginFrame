package gerror

import "fmt"

type Error interface {
	StatusCode() int32
	Error() string
	IsStable() bool
	WithCustomMessage(message string) Error
}

type gError struct {
	statusCode int32
	message    string
	stable     bool

	err error
}

var (
	errorMap map[int32]*gError
)

func (g *gError) Error() string {
	return g.message
}

func (g *gError) StatusCode() int32 {
	return g.statusCode
}

func (g *gError) IsStable() bool {
	return g.stable
}

func (g *gError) clone() *gError {
	return &gError{
		statusCode: g.statusCode,
		message:    g.message,
		stable:     g.stable,
		err:        g.err,
	}
}

func (g *gError) WithCustomMessage(message string) Error {
	e := g.clone()
	e.message = message
	return e
}

func buildError(statusCode int32, msg string) *gError {
	if errorMap == nil {
		errorMap = make(map[int32]*gError)
	}

	_, ok := errorMap[statusCode]
	if ok {
		panic(fmt.Sprintf("duplicated error code definition. status=%+v", statusCode))
	}
	err := &gError{
		statusCode: statusCode,
		message:    msg,
	}

	errorMap[statusCode] = err

	return err
}
