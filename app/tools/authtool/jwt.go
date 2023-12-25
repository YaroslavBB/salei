package authtool

import (
	"os"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
)

var (
	tokenEncodeString   = []byte(os.Getenv("TOKEN_HASH"))
	accessTokenDuration = time.Minute * 60
)

// Claims  claims
type Claims struct {
	OperID        int    `json:"oper_id"`
	OperLogin     string `json:"oper_login"`
	RefreshToken  string `json:"refresh_token"`
	AccessExpired int64  `json:"access_expired"`
	jwt.StandardClaims
}

// NewToken создание нового токена
func NewToken(profileID int, login string, refreshToken string) (string, error) {
	// Create the Claims
	claims := Claims{
		profileID,
		login,
		refreshToken,
		time.Now().Add(accessTokenDuration).Unix(),
		jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(tokenEncodeString)
}

// GenerateRefreshToken генерация refreshtoken в uuid
func GenerateRefreshToken() string {
	uu, _ := uuid.NewV4()
	return uu.String()
}
