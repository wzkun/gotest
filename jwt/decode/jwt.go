package decode

import (
	"time"

	"auth.server/config"
	"auth.server/errstring"
	pbauth "auth.server/proto/golang/auth"
	"code.aliyun.com/bim_backend/zoogoer/gun/errors"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT 生成jwt
func GenerateJWT(platform, account string, deviceInfo *pbauth.DeviceInfo) (string, error) {
	secret := config.SharePrivateConfigInstance().LogonTokenSettings.Secret
	expiration := config.SharePrivateConfigInstance().LogonTokenSettings.Expiration

	claims := make(jwt.MapClaims)
	claims["account"] = account
	claims["platform"] = platform
	claims["device"] = deviceInfo
	claims["exp"] = time.Now().Unix() + expiration
	claims["iat"] = time.Now().Unix()

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	loginToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return loginToken, nil
}

// VerifyJWT 校验jwt
func VerifyJWT(loginToken, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(loginToken, secretFunc(secret))
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.NewClientErr(nil, errstring.LoginTokenValid, "", "token.VerifyJWT", nil)
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims, nil
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}
