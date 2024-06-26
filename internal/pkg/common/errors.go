package common

import "errors"

var (
	ErrDuplicate      = errors.New("Entity Exists Already")
	ErrRecordNotFound = errors.New("Entity Does Not Exist")
)
