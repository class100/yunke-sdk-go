package yunke

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/storezhang/gox"
)

type (
	UserClaims struct {
		gox.BaseUser
		jwt.StandardClaims
	}
)

func token(method jwt.SigningMethod, secret string) (token string, err error) {
	claims := jwt.NewWithClaims(method, UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
		BaseUser: gox.BaseUser{
			BaseStruct: gox.BaseStruct{
				Id: 1,
			},
		},
	})
	token, err = claims.SignedString([]byte(secret))

	return
}
