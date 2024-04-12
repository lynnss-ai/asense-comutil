package encryptutil

import (
	"bytes"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword bcrypt根据密码和盐值生成密码
// password: 密码
// salt:     盐值
//
// returns:
//
//	hash: 生成的密码
//	err:  生成密码错误信息
func GeneratePassword(password, salt string) (string, error) {
	var buf bytes.Buffer
	buf.WriteString(password)
	buf.WriteString(salt)
	hash, err := bcrypt.GenerateFromPassword([]byte(buf.String()), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("Created Password Error")
	}
	return string(hash), err
}

// ComparePassword 根据密文和密码+盐值验证密码是否正确
// sourcePassword: 密文
// password:       密码
// salt:           盐值
//
// returns:
//
//	true:  密码正确
//	false: 密码错误
func ComparePassword(sourcePassword, password, salt string) bool {
	var buf bytes.Buffer
	buf.WriteString(password)
	buf.WriteString(salt)
	err := bcrypt.CompareHashAndPassword([]byte(sourcePassword), buf.Bytes())
	if err != nil {
		return false
	}
	return true
}
