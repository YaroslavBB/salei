package authtool

import (
	"salei/app/internal/entity/auth"
	"salei/app/internal/entity/global"

	"github.com/GehirnInc/crypt"
	"github.com/sirupsen/logrus"
)

func CheckMD5CryptPassword(srcHash, rawPassword string, log *logrus.Logger) (err error) {
	cryptv := crypt.MD5.New()
	if err = cryptv.Verify(srcHash, []byte(rawPassword)); err != nil {
		if err == crypt.ErrKeyMismatch {
			err = auth.ErrIncorrectLoginOrPassword
			return
		}
		log.Errorf("method CheckMD5CryptPassword; error: %v", err)
		err = global.ErrInternalError
		return
	}
	return
}
