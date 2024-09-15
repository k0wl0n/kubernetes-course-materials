// Contains cache util functions to facilitate storing and retrieving order details

package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"github.com/k0wl0n/go-ecommerce-order/models"
)

func StoreOrderToCache(client *redis.Client, ctx *gin.Context, order models.OrderDetail) {
	key := order.ID.String()
	value, err := models.OrderStructToByte(order)

	if err != nil {
		log.Errorln("Error caught while converting order details to bytes: ", err)
		return
	}

	err = client.Set(ctx, key, value, 1*time.Hour).Err()

	if err != nil {
		log.Errorln("Error caught while saving order details into cache: ", err)
	}
}

func RetriveOrderFromCache(client *redis.Client, ctx *gin.Context, orderID string) (models.OrderDetail, error) {
	orderByte, err := client.Get(ctx, orderID).Bytes()
	if err != nil {
		return models.OrderDetail{}, err
	}

	order, err := models.ByteToOrderStruct(orderByte)
	if err != nil {
		return models.OrderDetail{}, err
	}

	return order, nil
}

func DeleteOrderFromCache(client *redis.Client, ctx *gin.Context, orderID string) {
	err := client.Del(ctx, orderID).Err()
	log.Errorln("Error caught while deleting order from cache: ", err)
}
