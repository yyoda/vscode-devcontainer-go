package data

import (
	"log"
	"time"
)

// Set .
func Set(key string, value interface{}, expire time.Duration) {
	err := kvs.Set(key, value, expire).Err()
	if err != nil {
		log.Fatal(err)
		return
	}
}

// Get .
func Get(key string) string {
	val, err := kvs.Get(key).Result()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return val
}
