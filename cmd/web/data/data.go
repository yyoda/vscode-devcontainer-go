package data

import (
	"github.com/go-redis/redis/v7"
	"github.com/jmoiron/sqlx"
	"os"
)

var db *sqlx.DB
var kvs *redis.Client

func init() {
	// mysql
	dbConnStr := getEnv("MYSQL_CONNECTION", "dev:dev@tcp(mysql:3306)/main?parseTime=true")
	sqlxdb, err := sqlx.Connect("mysql", dbConnStr)
	if err != nil {
		panic(err)
	}
	db = sqlxdb

	// redis
	redisConnStr := getEnv("REDIS_CONNECTION", "redis:6379")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisConnStr,
		Password: "",
		DB:       0,
	})
	kvs = redisClient
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}
