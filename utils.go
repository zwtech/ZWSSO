package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func encryptString(str string) string {
	h := sha256.New()
	h.Write([]byte(str + salt))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func getCurrentTime() string {
	current := time.Now()
	timestamp := current.Unix()
	return strconv.FormatInt(timestamp, 10)
}
