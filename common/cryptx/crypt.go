package cryptx

import (
	"fmt"

	"golang.org/x/crypto/scrypt"
)

const salt = "tiktok"

func PasswordEncrypt(password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}
func PasswordVerify(password, hash string) bool {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk)) == hash
}
