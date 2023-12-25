package external

import (
	"salei/app/internal/usecase"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AuthExternal struct {
	log *logrus.Logger
	*usecase.AuthUsecase
	db *sqlx.DB
}

func NewAuthExternal(log *logrus.Logger, db *sqlx.DB, u *usecase.AuthUsecase) *AuthExternal {
	return &AuthExternal{
		log:         log,
		AuthUsecase: u,
		db:          db,
	}
}

func (e *AuthExternal) Login()
