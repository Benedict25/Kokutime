package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
)

var jwtKey = []byte("!@#123abc")
var tokenName = "token"
var onlineId = -1

// ctx (global) for redis
var ctx = context.Background()

// GoRedis - Set & Get from Redis
func SetRedis(rdb *redis.Client, key string, value string, expiration int) {
	err := rdb.Set(ctx, key, value, 0).Err()
	CheckError(err)
}

func GetRedis(rdb *redis.Client, key string) string {
	val, err := rdb.Get(ctx, key).Result()

	CheckError(err)
	return val
}

type Claims struct {
	ID       int    `json:id`
	Name     string `json:"name"`
	UserType int    `json:user_type`
	jwt.StandardClaims
}

func generateToken(w http.ResponseWriter, id int, name string, userType int) {
	tokenExpiryTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		ID:       id,
		Name:     name,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}

	// Set the now-logged-in-user's id to onlineId for global use
	onlineId = id
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	SetRedis(rdb, "onlineId", strconv.Itoa(id), 0)
	onlineId, _ = strconv.Atoi(GetRedis(rdb, "onlineId"))
	// Can do onlineId = id -> but not safe

	// Continue Generate Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)

	if err != nil {
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: true,
	})
}

func resetToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	})
}

func Authenticate(next http.HandlerFunc, accessType int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidToken := validateUserToken(r, accessType)
		if !isValidToken {
			PrintError(400, "Unathorized Access!", w)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func validateUserToken(r *http.Request, accessType int) bool {
	isAccessTokenValid, id, name, userType := validateTokenFromCookies(r)
	fmt.Println("Using: ", id, " ", name)

	if isAccessTokenValid {
		isUserValid := userType == accessType
		if isUserValid {
			return true
		}
	}
	return false
}

func validateTokenFromCookies(r *http.Request) (bool, int, string, int) {
	if cookie, err := r.Cookie(tokenName); err == nil {
		accessToken := cookie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.ID, accessClaims.Name, accessClaims.UserType
		}
	}
	return false, -1, "", -1
}
