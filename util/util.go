package util

import (
	"math/rand"
	"time"
)

// RandomString 随机生成名称
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range result {
		result[i] = letters[r.Intn(len(letters))]
	}
	return string(result)
}
