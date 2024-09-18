// Contains cache util functions to store or retrive product details

package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func StoreProductToCache(client *redis.Client, ctx *gin.Context, product ProductDetail) {
	key := product.ID.String()
	value, err := ProductStructToByte(product)
	if err != nil {
		log.Errorln("Error caught while converting product struct to byte: ", err)
		return
	}

	err = client.Set(ctx, key, value, 1*time.Hour).Err()
	if err != nil {
		log.Errorln("Error caught while saving product details into cache: ", err)
	}
}

func RetriveProductFromCache(client *redis.Client, ctx *gin.Context, productID string) (ProductDetail, error) {
	productByte, err := client.Get(ctx, productID).Bytes()
	if err != nil {
		return ProductDetail{}, err
	}

	product, err := ByteToProductStruct(productByte)
	if err != nil {
		return ProductDetail{}, err
	}

	return product, nil
}

func DeleteProductFromCache(client *redis.Client, ctx *gin.Context, productID string) {
	err := client.Del(ctx, productID).Err()
	log.Errorln("Error caught while deleting product details from cache: ", err)
}
