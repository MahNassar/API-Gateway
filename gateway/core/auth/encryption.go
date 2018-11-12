package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
	"encoding/base64"
	AppLogger "api_gateway/gateway/core/logger"
)

const (
	keyFile = "core/auth/pem.key"
)

var IV = []byte(os.Getenv("IV"))

func encryption(plain_string string) (string, error) {
	logger := AppLogger.GetLogInstance()

	keyStr := os.Getenv("ENC_SECRET")
	key := []byte(keyStr)

	plaintext := []byte(plain_string)

	block, err := aes.NewCipher(key)
	if err != nil {
		logger.AddStep("encryption", err.Error())
		return "", err
	}

	stream := cipher.NewCTR(block, IV)
	stream.XORKeyStream(plaintext, plaintext)
	base64encoded := base64.StdEncoding.EncodeToString(plaintext)
	logger.AddStep("encryption", "")
	return base64encoded, nil

}

func decryption(ss string) []byte {

	keyStr := os.Getenv("ENC_SECRET")
	key := []byte(keyStr)

	data, err := base64.StdEncoding.DecodeString(ss)

	if err != nil {

		fmt.Println("error:", err)

		return nil

	}
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(data, data)
	fmt.Println(string(data))
	return data
}
