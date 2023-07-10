package errors

import (
	"fmt"
	"runtime"

	"github.com/pkg/errors"
)

type Error struct {
	ErrMsg       string      `json:"err_msg"`
	ErrDetail    interface{} `json:"err_detail"`
	ErrMsgParams interface{} `json:"err_msg_params"`
	UserMsg      string      `json:"user_msg"`      //用户测外显
	UserErrCode  int32       `json:"user_err_code"` //用户侧外显
	HttpCode     int         `json:"-"`
	Stack        string      `json:"-"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("stack=%s|err_msg=%s|err_detail=%v|err_msg_params=%s|user_msg=%s|user_error_code=%d",
		e.Stack, e.ErrMsg, e.ErrDetail, e.ErrMsgParams, e.UserMsg, e.UserErrCode)
}

func Wrap(e error) error {
	if e == nil {
		return nil
	}
	if err, ok := e.(*Error); ok {
		return &Error{
			ErrMsg:       err.ErrMsg,
			ErrDetail:    err.ErrDetail,
			ErrMsgParams: err.ErrMsgParams,
			UserMsg:      err.UserMsg,
			UserErrCode:  err.UserErrCode,
			//HttpCode:     err.HttpCode,
			Stack: GetStack(),
		}
	}
	return errors.Wrap(e, fmt.Sprintf("%v - %v", e.Error(), GetStack()))
}

func New(msg string) error      { return errors.New(msg) }
func Is(err, target error) bool { return errors.Is(err, target) }

var (
	ErrServer = &Error{ErrMsg: "error_server", UserErrCode: 200102}
	ErrParam  = &Error{ErrMsg: "error_param", UserErrCode: 200101}
	//ErrCallFreq = &Error{ErrMsg: "error_call_limit_over", HttpCode: http.StatusBadRequest}
	//ErrAuth     = &Error{ErrMsg: "error_auth", HttpCode: http.StatusForbidden}
	//ErrNotLogin = &Error{ErrMsg: "error_not_login", HttpCode: http.StatusForbidden}

	//ConcurrentlyOperateErr = &Error{ErrMsg: "error_concurrently_operate_exception", HttpCode: http.StatusBadRequest}
)

func GetStack() string {
	buf := make([]byte, 10240)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}

func GetErrCause(e error) string {
	cause := errors.Cause(e)
	return cause.Error()
}
