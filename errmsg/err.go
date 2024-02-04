package errmsg

import "errors"

var (
	DBUseSelf = errors.New("please use DropDBSelf function")
)
