package crypto

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/pbkdf2"
	"io"
)

func SecureRandomBytes(nbytes uint32) (secureBytes []byte, err error) {
	secureBytes = make([]byte, nbytes)
	_, err = rand.Read(secureBytes)
	if err != nil {
		logrus.Fatal(err)
		return
	}
	return
}

func DeriveKey(secret []byte, salt []byte) (key []byte) {
	key = pbkdf2.Key(secret, salt, int(Params.iterations), int(Params.keyLength), crypto.SHA3_512.New)
	return
}

func DeriveUrlSafeKey(secret []byte, salt []byte) (key []byte) {
	byteKey := DeriveKey(secret, salt)
	key = make([]byte, Params.keyLength)
	base64.URLEncoding.Encode(key, byteKey)
	return
}

func Encrypt(key []byte, nonce []byte, text string) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext := []byte(text)
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext
}

func Decrypt(key []byte, nonce []byte, ciphertext []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return string(plaintext)
}
