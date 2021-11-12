package encry

import (
	"encoding/base64"
	"log"

	"golang.org/x/crypto/scrypt"
)

//密码加密
func ScryptPasswd(password string) string {
	const keyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 22, 66, 222, 111}

	hashPasswd, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}

	realPwd := base64.StdEncoding.EncodeToString(hashPasswd)
	return realPwd
}
