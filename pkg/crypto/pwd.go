package crypto

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Argon2 constants struct
// for now they are going to be constants thinking about env vars tho
const argon2Time uint32 = 3
const argon2Memory uint32 = 64 * 1024 // 64 MB
const argon2Threads uint8 = 4
const argon2KeyLen uint32 = 32
const saltLen uint32 = 16

func HashPassword(password string) (string, error) {
	salt, err := generateSalt(saltLen)
	if err != nil {
		return "", err
	}

	hash, err := generateHash([]byte(password), salt)
	if err != nil {
		return "", err
	}

	b64Hash := base64.StdEncoding.EncodeToString(hash)
	b64Salt := base64.StdEncoding.EncodeToString(salt)

	return b64Hash + "$" + b64Salt, nil
}

func VerifyPassword(userHash string, password string) (bool, error) {
	contents := strings.Split(userHash, "$")
	if len(contents) != 2 {
		return false, fmt.Errorf("provided hash split was not correct length")
	}

	hash, err := base64.StdEncoding.DecodeString(contents[0])
	if err != nil {
		return false, nil
	}

	salt, err := base64.StdEncoding.DecodeString(contents[1])
	if err != nil {
		return false, nil
	}

	check, err := generateHash([]byte(password), salt)
	if err != nil {
		return false, nil
	}

	return bytes.Equal(hash, check), nil
}

func generateHash(plaintext, salt []byte) ([]byte, error) {
	if len(plaintext) == 0 {
		return nil, fmt.Errorf("plaintext has len of 0")
	}

	if len(salt) == 0 {
		return nil, fmt.Errorf("salt has len of 0")
	}

	hash := argon2.IDKey(plaintext, salt, argon2Time, argon2Memory, argon2Threads, argon2KeyLen)
	return hash, nil
}

func generateSalt(length uint32) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}
