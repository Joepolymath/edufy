package auth

import "errors"

var (
	ErrDuplicate      = errors.New("User Account Exists Already")
	ErrRecordNotFound = errors.New("User Account Does Not Exist")
	ErrCredentials    = errors.New("Invalid User Credentials Provided")
	ErrPassword       = errors.New("Incorrect Password")
	ErrUnverified     = errors.New("User Account Not Verified")
	ErrOTP            = errors.New("Invalid OTP Provided")
	ErrID             = errors.New("Invalid User ID")
)
