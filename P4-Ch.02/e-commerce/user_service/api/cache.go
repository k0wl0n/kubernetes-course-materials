// Contains cache util functions to store or retrive user details

package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func StoreUserToCache(client *redis.Client, ctx *gin.Context, user User) {
	key := user.ID.String()
	value, err := UserStructToByte(user)
	if err != nil {
		log.Errorln("Error caught while converting user struct to byte: ", err)
		return
	}

	err = client.Set(ctx, key, value, 1*time.Hour).Err()
	if err != nil {
		log.Errorln("Error caught while saving user details into cache: ", err)
	}
}

func RetriveUserFromCache(client *redis.Client, ctx *gin.Context, userID string) (User, error) {
	userByte, err := client.Get(ctx, userID).Bytes()
	if err != nil {
		return User{}, err
	}

	user, err := ByteToUserStruct(userByte)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func DeleteUserFromCache(client *redis.Client, ctx *gin.Context, userID string) {
	err := client.Del(ctx, userID).Err()
	log.Errorln("Error caught while deleting user details from cache: ", err)
}
