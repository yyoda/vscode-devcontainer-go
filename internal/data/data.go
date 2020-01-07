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
	redisClient := redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_CONNECTION", "redis:6379"),
		Password: "",
		DB:       0,
	})
	kvs = redisClient
}

func getEnv(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultValue
	}
	return val
}
