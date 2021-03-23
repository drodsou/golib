package aestool

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// Cipher a text with a key
func Cipher(clearStr, key string) string {
	text := []byte(clearStr)

	gcm := getGCM(key)
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure random sequence
	if _, err := io.ReadFull(crand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	ciphered := gcm.Seal(nonce, nonce, text, nil)
	cipheredB64 := base64.StdEncoding.EncodeToString(ciphered)
	return cipheredB64
}

// Decipher text with key
func Decipher(cipheredB64, key string) string {
	gcm := getGCM(key)

	ciphered, _ := base64.StdEncoding.DecodeString(cipheredB64)
	nsize := gcm.NonceSize()
	nonce, cipheredText := ciphered[:nsize], ciphered[nsize:]
	plainText, err := gcm.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		fmt.Println(err)
	}
	return string(plainText)
}

// SafeChars chars without quotes and backslashes
const SafeChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-*/=<>{}[]()&%$#@!"

// RandString32 returns random key of 32 bytes
// pass a string to pick chars from it, or no argument makes cripto random one
func RandString32(charsA ...string) string {
	bytes := make([]byte, 32)

	if len(charsA) == 0 {
		_, _ = io.ReadFull(crand.Reader, bytes)
	} else {
		chars := charsA[0]
		rand.Seed(time.Now().UTC().UnixNano())
		for i := range bytes {
			bytes[i] = chars[rand.Intn(len(chars))]
		}
	}
	return string(bytes)
}

// -------- HELPER

func getGCM(keyStr string) cipher.AEAD {
	if len(keyStr) != 32 {
		panic("CreateAESTool: key must be 32 chars long")
	}

	key := []byte(keyStr)
	ciph, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(ciph)
	if err != nil {
		fmt.Println(err)
	}

	return gcm
}
