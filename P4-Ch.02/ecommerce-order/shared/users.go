package shared

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/redis/go-redis/v9"
)

type userResponse struct {
	Data struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"data"`
}

// Call user token API to fetch user details
func getUserDetails(ctx *gin.Context, payload map[string]interface{}) (userResponse, error) {
	apiSecretKey, err := GetAPISecretKey()
	if err != nil {
		return userResponse{}, err
	}

	userServiceHost, err := GetUserServiceHost()
	if err != nil {
		return userResponse{}, err
	}

	requestURL := GetBaseURL(ctx.Request, userServiceHost, "token")

	clinet := resty.New()

	response := &userResponse{}

	rawResp, err := clinet.R().
		SetResult(response).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Secret", apiSecretKey).
		SetBody(payload).
		Post(requestURL)

	if rawResp.StatusCode() != http.StatusOK || err != nil {
		return userResponse{}, fmt.Errorf("error caught while calling user details API")
	}

	return *response, nil
}

// Fetch user ID from cache, Else call the user microserivce to get the user details
func GetUserFromToken(client *redis.Client, ctx *gin.Context, token string) (string, error) {
	// Check if key is available in the cache or not
	userID, err := client.Get(ctx, token).Result()
	if err == nil && userID != "" {
		return userID, nil
	}

	payload := map[string]interface{}{
		"token": token,
	}

	response, err := getUserDetails(ctx, payload)
	if err != nil {
		log.Errorln("Error while fetching user details from user service: ", err)
		return "", err
	}

	// Set userID into cache
	err = client.Set(ctx, "token", response.Data.ID, 1*time.Hour).Err()
	if err != nil {
		log.Errorln("Error while saving userID into cache: ", err)
		return "", err
	}

	return response.Data.ID, nil
}
