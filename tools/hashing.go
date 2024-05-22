package tools

import "math/rand"

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func CreateUrlHash() string {
	hash := ""
	for i := 0; i < 8; i++ {
		hash += string(chars[rand.Intn(62)])
	}
	return hash
}
