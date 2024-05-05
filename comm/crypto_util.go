package comm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// EncryptorDecryptor 是一个包含加密和解密方法的结构体
type EncryptorDecryptor struct {
	key []byte
}

// NewEncryptorDecryptor 创建一个新的EncryptorDecryptor实例
func NewEncryptorDecryptor(key string) *EncryptorDecryptor {
	return &EncryptorDecryptor{
		key: []byte(key),
	}
}

// Encrypt 方法用于加密字符串
func (ed *EncryptorDecryptor) Encrypt(text string) (string, error) {
	plaintext := []byte(text)

	block, err := aes.NewCipher(ed.key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 方法用于解密字符串
func (ed *EncryptorDecryptor) Decrypt(encodedText string) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(encodedText)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	block, err := aes.NewCipher(ed.key)
	if err != nil {
		return "", err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

func Encode(text, key string) string {
	//key := CfgLoader.GetString("key.aes_32_code")
	ed := NewEncryptorDecryptor(key)

	encrypted, err := ed.Encrypt(text)
	if err != nil {
		panic(err)
	}
	return encrypted
}

func Decode(text, key string) string {
	ed := NewEncryptorDecryptor(key)
	decrypted, err := ed.Decrypt(text)
	if err != nil {
		panic(err)
	}
	return decrypted
}
