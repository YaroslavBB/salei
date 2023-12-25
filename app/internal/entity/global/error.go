package global

import "errors"

var (
	// ErrNoData нет данных
	ErrNoData = errors.New("нет данных")
	// ErrInternalError возникла внутренняя ошибка
	ErrInternalError = errors.New("возникла внутренняя ошибка")
)
