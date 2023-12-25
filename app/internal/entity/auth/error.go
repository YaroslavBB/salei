package auth

import "errors"

var (
	// ErrIncorrectLoginOrPassword неверно введен логин или пароль
	ErrIncorrectLoginOrPassword = errors.New("неверно введен логин или пароль")
)
