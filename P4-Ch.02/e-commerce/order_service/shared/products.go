package shared

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"github.com/k0wl0n/go-ecommerce-order/models"
)

type productResponse struct {
	Data models.Product `json:"data"`
}

// Call product details API
func getProductDetails(ctx *gin.Context, payload map[string]interface{}) (productResponse, error) {
	apiSecretKey, err := GetAPISecretKey()
	if err != nil {
		return productResponse{}, err
	}

	productServiceHost, err := GetProductServiceHost()
	if err != nil {
		return productResponse{}, err
	}

	requestURL := GetBaseURL(ctx.Request, productServiceHost, "product-details")

	clinet := resty.New()

	response := &productResponse{}

	rawResp, err := clinet.R().
		SetResult(response).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Secret", apiSecretKey).
		SetBody(payload).
		Post(requestURL)

	if rawResp.StatusCode() != http.StatusOK || err != nil {
		return productResponse{}, fmt.Errorf("error caught while calling product details API")
	}

	return *response, nil
}

// Fetch product details from the given ID
func GetProductIDToDetails(client *redis.Client, ctx *gin.Context, productID string) (models.Product, error) {
	// Check if key is available in the cache or not
	productByte, err := client.Get(ctx, productID).Bytes()
	if err == nil && productByte != nil {
		product, err := models.ByteToProductStruct(productByte)
		if err != nil {
			log.Errorln("Error while converting product byte to product struct obj: ", err)
		} else {
			return product, nil
		}
	}

	payload := map[string]interface{}{
		"id": productID,
	}

	response, err := getProductDetails(ctx, payload)
	if err != nil {
		log.Errorln("Error while fetching product details from product service: ", err)
		return models.Product{}, err
	}

	productByte, err = models.ProductStructToByte(response.Data)
	if err != nil {
		log.Errorln("Error while converting product details to bytes", err)
	}

	// Save product details into cache
	err = client.Set(ctx, productID, productByte, 1*time.Hour).Err()
	if err != nil {
		log.Errorln("Error while saving product details into cache: ", err)
	}

	return response.Data, nil
}
