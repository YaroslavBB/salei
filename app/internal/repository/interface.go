package repository

import (
	"salei/app/internal/entity/auth"

	"github.com/jmoiron/sqlx"
)

type Auth interface {
	FindByLogin(tx sqlx.Tx, login string) (data auth.Profile, err error)
}
