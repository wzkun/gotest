package decode

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"github.com/wzkun/aurora/utils/decode"
)

// CFBDecrypter 解密
func CFBDecrypter(key, iv, code string) string {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	newkey := []byte(key)
	newiv := []byte(iv)
	ciphertext, _ := base64.StdEncoding.DecodeString(code)

	block, err := aes.NewCipher(newkey)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, newiv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	fmt.Printf("%s", ciphertext)

	return string(ciphertext)
}

type EncodedLoginToken struct {
	Key   string
	Iv    string
	Token string
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

	plaintext := []byte(code)

	block, err := aes.NewCipher([]byte(newkey))
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	stream := cipher.NewCFBEncrypter(block, newiv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	dst := decode.StdEncodeToString(ciphertext)

	resp := &EncodedLoginToken{}
	resp.Key = key
	resp.Iv = iv
	resp.Token = dst
	return resp, nil
}
