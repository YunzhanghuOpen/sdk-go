package errorx

import "fmt"

const (
	OK = "0000"
)

// Error err
type Error struct {
	Code    string
	Message string
	Data    interface{}
}

// Error 实现error接口
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// New 新建Err
func New(code, msg string, data interface{}) error {
	return &Error{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

// FromError 解析Err
func FromError(err error) (*Error, bool) {
	if err == nil {
		return nil, false
	}

	if e, ok := err.(*Error); ok {
		return e, true
	}
	return nil, false
}
