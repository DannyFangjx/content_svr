package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	buf := make([]rune, n)
	for i := range buf {
		buf[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(buf)
}

func UUID() string {
	uid, _ := uuid.NewRandom()
	return strings.ReplaceAll(uid.String(), "-", "")
}
