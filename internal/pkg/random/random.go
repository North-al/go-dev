package random

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenerateRandomUsername() string {
	// 随机生成6位字符串 + 时间戳
	timestamp := time.Now().UnixNano()
	randomString := fmt.Sprintf("user_%d_%s", timestamp, GenerateRandomString(6))
	return randomString

}
