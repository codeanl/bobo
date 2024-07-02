package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

type _encrypt struct{}

var Encryptor = new(_encrypt)

// ValidateCode 验证码
func (*_encrypt) ValidateCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

// MD5 加密
func (*_encrypt) MD5(str string, b ...byte) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(b))
}

// 使用 bcrypt 对密码进行加密生成一个哈希值
func (*_encrypt) BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// 使用 bcrypt 对比 明文密码 和 数据库中哈希值
func (*_encrypt) BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
