package controllers

import (
	"context"

	"github.com/go-redis/redis/v8"
	gomail "gopkg.in/mail.v2"
)

var ctx = context.Background()

func SetRedis(rdb *redis.Client, key string, value string, expiration int) {
	err := rdb.Set(ctx, key, value, 0).Err()
	CheckError(err)
}

func GetRedis(rdb *redis.Client, key string) string {
	val, err := rdb.Get(ctx, key).Result()

	CheckError(err)
	return val
}

func sendMail(to string, subject string, text string) {
	gmail := gomail.NewMessage()

	gmail.SetHeader("From", "cobapbp@gmail.com")
	gmail.SetHeader("To", to)
	gmail.SetHeader("Subject", subject)
	gmail.SetBody("text/plain", text)

	gm := gomail.NewDialer("smtp.gmail.com", 587, "cobapbp@gmail.com", "CobaPBP5656")

	err := gm.DialAndSend(gmail)
	CheckError(err)
}
