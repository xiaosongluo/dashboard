package models

import (
	"time"
	"crypto/md5"
	"fmt"
)

//GenerateAPIKey generate api key
func GenerateAPIKey() string {
	t := time.Now().String()
	sum := md5.Sum([]byte(t))
	return fmt.Sprintf("%x", sum)
}
