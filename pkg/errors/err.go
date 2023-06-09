package errors

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"runtime"
)

// Error an error with caller stack information
// Error 一个能调用方堆栈信息的错误
type Error interface {
	error
	t()
}

var _ Error = (*item)(nil)
var _ fmt.Formatter = (*item)(nil)

type item struct {
	msg   string
	stack []uintptr
}

func (i *item) Error() string {
	return i.msg
}

func (i *item) t() {}

func callers() []uintptr {
	var pcs [32]uintptr
	l := runtime.Callers(3, pcs[:])
	return pcs[:l]
}

// Format used by go.uber.org/zap in Verbose
func (i *item) Format(s fmt.State, verb rune) {
	io.WriteString(s, i.msg)
	io.WriteString(s, "\n")

	for _, pc := range i.stack {
		fmt.Fprintf(s, "%+v\n", errors.Frame(pc))
	}
}

// New create a new error
// New 创建一个新错误
func New(msg string) Error {
	return &item{
		msg:   msg,
		stack: callers(),
	}
}

// Errorf create a new error
// Errorf 创建一个新错误
func Errorf(format string, args ...interface{}) Error {
	return &item{
		msg:   fmt.Sprintf(format, args...),
		stack: callers(),
	}
}

// Wrap with some extra message into err
// Wrap 将额外信息包装进err
func Wrap(err error, msg string) Error {
	if err == nil {
		return nil
	}

	if e, ok := err.(*item); !ok {
		return &item{msg: fmt.Sprintf("%s; %s", msg, err.Error()), stack: callers()}
	} else {
		e.msg = fmt.Sprintf("%s; %s", msg, e.msg)
		return e
	}
}

// Wrapf with some extra message into err
func Wrapf(err error, format string, args ...interface{}) Error {
	if err == nil {
		return nil
	}

	msg := fmt.Sprintf(format, args...)

	if e, ok := err.(*item); !ok {
		return &item{msg: fmt.Sprintf("%s; %s", msg, err.Error()), stack: callers()}
	} else {
		e.msg = fmt.Sprintf("%s; %s", msg, e.msg)
		return e
	}
}

// WithStack add caller stack information
func WithStack(err error) Error {
	if err == nil {
		return nil
	}

	if e, ok := err.(*item); ok {
		return e
	}

	return &item{msg: err.Error(), stack: callers()}
}
