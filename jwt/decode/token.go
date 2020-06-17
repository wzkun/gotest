package decode

import (
	"fmt"

	"auth.server/config"
	"github.com/wzkun/aurora/utils/decode"
)

// EncodeToken token加密
func EncodeToken(loginToken, publicKey string) (*EncodedLoginToken, error) {
	encodedToken, err := CFBEncrypter(loginToken)
	if err != nil {
		return nil, err
	}

	newkey, _ := decode.EncodeRSAWithKeyFile("config/pub.pem", []byte(encodedToken.Key))
	newiv, _ := decode.EncodeRSAWithKeyFile("config/pub.pem", []byte(encodedToken.Iv))
	encodedToken.Key = newkey
	encodedToken.Iv = newiv

	return encodedToken, nil
}

// VerifyScodiToken 校验内部token
func VerifyScodiToken(token, key, iv string) {
	newkey, _ := decode.DecodeRSAWithKeyFile("config/pri.pem", key)
	newiv, _ := decode.DecodeRSAWithKeyFile("config/pri.pem", iv)

	decodedToken := CFBDecrypter(string(newkey), string(newiv), token)

	secret := config.SharePrivateConfigInstance().LogonTokenSettings.Secret

	decodedJWT, _ := VerifyJWT(decodedToken, secret)

	account := decodedJWT["account"].(string)
	fmt.Println("====account==", account)

	// return this.player.verifyToken(decodedToken, decodedJWT)
}
