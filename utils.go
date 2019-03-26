package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func encryptString(str string) string {
	h := sha256.New()
	h.Write([]byte(str + salt))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func getCurrentTime() int64 {
	current := time.Now()
	timestamp := current.UnixNano()
	return timestamp
}
