package postgres

import (
	"database/sql"
	"salei/app/internal/entity/auth"
	"salei/app/internal/entity/global"
	"salei/app/internal/repository"

	"github.com/jmoiron/sqlx"
)

type authRepo struct{}

func NewAuthRepo() repository.Auth {
	return &authRepo{}
}

func (*authRepo) FindByLogin(tx sqlx.Tx, login string) (data auth.Profile, err error) {
	err = tx.Get(&data, `
	 select p.prof_id, p.login, p.password
	 from profile
	 where login = $1
	 `, login)
	if err == sql.ErrNoRows {
		err = global.ErrNoData
	}

	return
}
