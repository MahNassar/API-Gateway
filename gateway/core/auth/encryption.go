package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"encoding/base64"
)

const (
	keyFile = "core/auth/pem.key"
)

var IV = []byte("1234567812345678")

func readKey(filename string) ([]byte, error) {
	key, err := ioutil.ReadFile(filename)
	if err != nil {
		return key, err
	}
	block, _ := pem.Decode(key)
	return block.Bytes, nil
}

func aesKey() []byte {
	file := fmt.Sprintf(keyFile)
	key, _ := readKey(file)

	return key
}

func createCipher() cipher.Block {
	c, err := aes.NewCipher(aesKey())
	if err != nil {
		log.Fatalf("Failed to create the AES cipher: %s", err)
	}
	return c
}

func encryption(plainText string) string {
	bytes := []byte(plainText)
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(bytes, bytes)
	base64encoded := base64.StdEncoding.EncodeToString(bytes)
	return base64encoded

}

func decryption(ss string) []byte {

	data, err := base64.StdEncoding.DecodeString(ss)

	if err != nil {

		fmt.Println("error:", err)

		return nil

	}
	blockCipher := createCipher()
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(data, data)
	fmt.Println(string(data))
	return data
}
