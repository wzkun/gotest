package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/wzkun/aurora/utils/decode"
)

const Secret = "eb407da3-2d60-49ae-9130-2e3893352362"
const Expiration = 60

var Token string

type DeviceInfo struct {
	DeviceUniqueId string
}

type JwtCustomClaims struct {
	jwt.StandardClaims
	// 追加自己需要的信息
	Account string `json:"account"`
	Device  *DeviceInfo
}

// GenerateJWT 生成jwt
func GenerateJWT(platform, account string, deviceInfo *DeviceInfo) string {
	claim := &JwtCustomClaims{}
	claim.Audience = account
	claim.IssuedAt = time.Now().Unix()
	claim.ExpiresAt = time.Now().Unix() + Expiration
	claim.Account = account
	claim.Device = &DeviceInfo{}
	claim.Device.DeviceUniqueId = deviceInfo.DeviceUniqueId

	fmt.Println("time.Now().Unix()==", time.Now().Unix())
	fmt.Println("claim==", claim)
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claim

	loginToken, _ := token.SignedString([]byte(Secret))

	return loginToken
}

// VerifyJWT 校验jwt
func VerifyJWT(loginToken, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(loginToken, secretFunc(secret))
	fmt.Println("======err===", err)
	if err != nil {
		return nil, err
	}
	claim := token.Claims.(jwt.MapClaims)
	account := claim["account"].(string)
	fmt.Println("====account===", account)
	fmt.Println("====claim===", claim)
	return claim, nil
}

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

func Test() {
	platform := "mobile"
	account := "wzkun"

	device := &DeviceInfo{}
	device.DeviceUniqueId = "699c5f56-474f-4643-8654-c888f5e0de8b"

	loginToken := GenerateJWT(platform, account, device)
	fmt.Println("====loginToken====", loginToken)

	VerifyJWT(loginToken, Secret)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ3emt1biIsImV4cCI6MTU5MTUxNzU1MSwiaWF0IjoxNTkxNTE3NDkxLCJhY2NvdW50Ijoid3prdW4iLCJEZXZpY2UiOnsiRGV2aWNlVW5pcXVlSWQiOiI2OTljNWY1Ni00NzRmLTQ2NDMtODY1NC1jODg4ZjVlMGRlOGIifX0.eRsqJrZVuGNhXo_jHDNj25OVZIoN0kM2nyjaXoMGxNI"
	VerifyJWT(token, Secret)

}

type EncodedLoginToken struct {
	Key   string
	Iv    string
	Token string
}

// CFBDecrypter 解密
func CFBDecrypter(key, iv, code string) string {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	newkey := []byte(key)
	newiv := []byte(iv)
	ciphertext, _ := base64.StdEncoding.DecodeString(code)
	// key, _ := hex.DecodeString("6368616e676520746869732070617373")
	// ciphertext, _ := hex.DecodeString("7dd015f06bec7f1b8f6559dad89f4131da62261786845100056b353194ad")
	// ciphertext, _ := hex.DecodeString(code)

	block, err := aes.NewCipher(newkey)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	// iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	fmt.Println("==iv===", iv)
	stream := cipher.NewCFBDecrypter(block, newiv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	fmt.Printf("ciphertext====%s\n", ciphertext)
	fmt.Println("==== string(ciphertext)====", string(ciphertext))

	return string(ciphertext)
	// Output: some plaintext
}

// CFBEncrypter 加密
func CFBEncrypter(code string) (*EncodedLoginToken, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// key, _ := hex.DecodeString("6368616e676520746869732070617373")
	key := decode.RandomString(16)
	newkey := []byte(key)
	iv := decode.RandomString(16)
	newiv := []byte(iv)
	fmt.Println("===key==", key)
	fmt.Println("===newkey==", newkey)

	plaintext := []byte(code)

	block, err := aes.NewCipher(newkey)
	if err != nil {
		return nil, err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	stream := cipher.NewCFBEncrypter(block, newiv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	dst := base64.StdEncoding.EncodeToString(ciphertext)

	resp := &EncodedLoginToken{}
	resp.Key = key
	resp.Iv = iv
	resp.Token = dst
	return resp, nil
}

func Test1() {
	plaintext := "some plaintext"
	EncodedToken, _ := CFBEncrypter(plaintext)
	fmt.Println("EncodedTokens===", EncodedToken)
	decoded := CFBDecrypter(EncodedToken.Key, EncodedToken.Iv, EncodedToken.Token)
	fmt.Println("===decoded====", decoded)

	if decoded == plaintext {
		fmt.Println("===decoded====")
	}
}

func EncodeToken(loginToken, publicKey string) *EncodedLoginToken {
	encodedToken, _ := CFBEncrypter(loginToken)
	newkey, _ := decode.EncodeRSAWithKeyFile("config/pub.pem", []byte(encodedToken.Key))
	newiv, _ := decode.EncodeRSAWithKeyFile("config/pub.pem", []byte(encodedToken.Iv))
	encodedToken.Key = newkey
	encodedToken.Iv = newiv

	fmt.Println("===encodedToken===", encodedToken)
	return encodedToken
}

func VerifyScodiToken(token, key, iv string) {
	newkey, _ := decode.DecodeRSAWithKeyFile("config/pri.pem", key)
	newiv, _ := decode.DecodeRSAWithKeyFile("config/pri.pem", iv)

	decodedToken := CFBDecrypter(string(newkey), string(newiv), token)
	decodedJWT, _ := VerifyJWT(decodedToken, Secret)
	fmt.Println("====decodedJWT==", decodedJWT)

	account := decodedJWT["account"].(string)
	fmt.Println("====account==", account)
	// return this.player.verifyToken(decodedToken, decodedJWT)
}

func Test2() {
	platform := "mobile"
	account := "wzkun"

	device := &DeviceInfo{}
	device.DeviceUniqueId = "699c5f56-474f-4643-8654-c888f5e0de8b"

	loginToken := GenerateJWT(platform, account, device)
	fmt.Println("====loginToken====", loginToken)

	encodeToken := EncodeToken(loginToken, "")
	fmt.Println("===encodeToken===", encodeToken)

	VerifyScodiToken(encodeToken.Token, encodeToken.Key, encodeToken.Iv)
	// CFBDecrypter(encodeToken.Key, encodeToken.Iv, encodeToken.Token)

	// VerifyJWT(loginToken, Secret)
	// token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ3emt1biIsImV4cCI6MTU5MTUxNzU1MSwiaWF0IjoxNTkxNTE3NDkxLCJhY2NvdW50Ijoid3prdW4iLCJEZXZpY2UiOnsiRGV2aWNlVW5pcXVlSWQiOiI2OTljNWY1Ni00NzRmLTQ2NDMtODY1NC1jODg4ZjVlMGRlOGIifX0.eRsqJrZVuGNhXo_jHDNj25OVZIoN0kM2nyjaXoMGxNI"
	// VerifyJWT(token, Secret)
}
func main() {
	// Test()
	Test2()
}
