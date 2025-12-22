package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"github.com/misakacoder/kagome/errs"
)

func EncryptBase64(original string, key string) string {
	encrypted, err := Encrypt([]byte(original), []byte(key))
	errs.Panic(err)
	return base64.StdEncoding.EncodeToString(encrypted)
}

func DecryptBase64(encrypted string, key string) string {
	enc, err := base64.StdEncoding.DecodeString(encrypted)
	errs.Panic(err)
	original, err := Decrypt(enc, []byte(key))
	errs.Panic(err)
	return string(original)
}

func EncryptHex(original string, key string) string {
	encrypted, err := Encrypt([]byte(original), []byte(key))
	errs.Panic(err)
	return hex.EncodeToString(encrypted)
}

func DecryptHex(encrypted string, key string) string {
	enc, err := hex.DecodeString(encrypted)
	errs.Panic(err)
	original, err := Decrypt(enc, []byte(key))
	errs.Panic(err)
	return string(original)
}

func Encrypt(original, key []byte) ([]byte, error) {
	validKey(key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	original = pkcs7Padding(original, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encrypted := make([]byte, len(original))
	blockMode.CryptBlocks(encrypted, original)
	return encrypted, nil
}

func Decrypt(encrypted, key []byte) ([]byte, error) {
	validKey(key)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	original := make([]byte, len(encrypted))
	blockMode.CryptBlocks(original, encrypted)
	original = pkcs7UnPadding(original)
	return original, nil
}

func pkcs7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pkcs7UnPadding(original []byte) []byte {
	length := len(original)
	unPadding := int(original[length-1])
	return original[:(length - unPadding)]
}

func validKey(key []byte) {
	if (len(key) & 0xF) != 0 {
		panic("the key length must be a multiple of 16")
	}
}
