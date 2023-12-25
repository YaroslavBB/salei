package usecase

import (
	"salei/app/internal/entity/auth"
	"salei/app/internal/entity/global"
	repository_import "salei/app/internal/repository/import"
	"salei/app/tools/authtool"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AuthUsecase struct {
	*repository_import.PostgresRepo
	log *logrus.Logger
}

func NewAuthUsecase(rp *repository_import.PostgresRepo, log *logrus.Logger) *AuthUsecase {
	return &AuthUsecase{
		PostgresRepo: rp,
		log:          log,
	}
}

func (u *AuthUsecase) Login(tx sqlx.Tx, login, password string) (profile auth.Profile, err error) {
	profile, err = u.Auth.FindByLogin(tx, login)
	switch err {
	case nil:
		//
	case global.ErrNoData:
		err = auth.ErrIncorrectLoginOrPassword
		return
	default:
		u.log.Errorf("method Login: FindByLogin; error: %v", err)
		err = global.ErrInternalError
		return
	}

	if err = authtool.CheckMD5CryptPassword(profile.Password, password, u.log); err != nil {
		return
	}

	if profile.Token, err = authtool.NewToken(profile.ProfID, profile.Login, authtool.GenerateRefreshToken()); err != nil {
		u.log.Errorf("method Login: authtool.NewToken; error: %v", err)
		err = global.ErrInternalError
		return
	}

	return
}
