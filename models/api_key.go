package models

import (
	"time"
	"crypto/md5"
	"fmt"
)

func GenerateAPIKey() string {
	t := time.Now().String()
	sum := md5.Sum([]byte(t))
	return fmt.Sprintf("%x", sum)
}
