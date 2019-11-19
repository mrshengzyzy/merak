package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AesEncryptWithBase64(data, key string) (string, error) {

	plaintext := []byte(data)
	keyIv := []byte(key)

	// key长度必须是 16, 24, 32, 对应 128, 192, 256 位加密
	block, err := aes.NewCipher(keyIv)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()
	plaintext = PKCS5Padding(plaintext, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, keyIv)
	cipherText := make([]byte, len(plaintext))
	blockMode.CryptBlocks(cipherText, plaintext)
	return base64.RawURLEncoding.EncodeToString(cipherText), nil
}

func AesDecryptWithBase64(cipherText, key string) (string, error) {

	keyIv := []byte(key)

	// 先使用Base64解码码
	originCipherText, _ := base64.RawURLEncoding.DecodeString(cipherText)

	block, err := aes.NewCipher(keyIv)
	if err != nil {
		panic(err)
	}
	blockModel := cipher.NewCBCDecrypter(block, keyIv)
	plaintext := make([]byte, len(originCipherText))
	blockModel.CryptBlocks(plaintext, originCipherText)
	plaintext = PKCS5UnPadding(plaintext)
	return string(plaintext), nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])
	return src[:(length - unPadding)]
}
