package global

import "errors"

type SystemError struct {
	errCode int
	errMsg  string
}

func (se *SystemError) ErrCode() int {
	return se.errCode
}

func (se *SystemError) ErrMsg() string {
	return se.errMsg
}

func (se *SystemError) Error() error {
	if se.errMsg == "" {
		return nil
	}
	return errors.New(se.errMsg)
}

var ()
