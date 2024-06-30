package utils

import (
	"fmt"
	"math/rand"
	"time"
)

type _encrypt struct{}

var Encryptor = new(_encrypt)

// ValidateCode 验证码
func (*_encrypt) ValidateCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
