package shared

import (
	"fmt"
	"net/http"
	"os"
	"time"

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

func getBaseURL(request *http.Request, host, endpoint string) string {
	protocol := "http"
	if request.TLS != nil {
		protocol += "s"
	}

	return fmt.Sprintf("%s://%s/internal/v1/%s/", protocol, host, endpoint)
}

func getAPISecretKey() (string, error) {
	apiSecretKey := os.Getenv("INTERNAL_API_SECRET")

	if apiSecretKey == "" {
		return "", fmt.Errorf("API secret is not configured")
	}

	return apiSecretKey, nil
}

// Call user token API to fetch user details
func getUserDetails(ctx *gin.Context, payload map[string]interface{}) (userResponse, error) {
	apiSecretKey, err := getAPISecretKey()
	if err != nil {
		return userResponse{}, err
	}

	userServiceHost := os.Getenv("USER_SERVICE_HOST")

	requestURL := getBaseURL(ctx.Request, userServiceHost, "token")

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
		return userResponse{}, fmt.Errorf("error caught while fetching user details API")
	}

	return *response, nil
}

// Fetch user ID from cache, Else call the user microserivce to get the user details
func GetUserFromToken(client *redis.Client, ctx *gin.Context, token string) (string, error) {
	// Check if key is available in the cache or not
	userID, err := client.Get(ctx, token).Result()
	if err != nil && userID != "" {
		return userID, nil
	}

	payload := map[string]interface{}{
		"token": token,
	}

	response, err := getUserDetails(ctx, payload)
	if err != nil {
		return "", err
	}

	// Set userID into cache
	err = client.Set(ctx, "token", response.Data.ID, 1*time.Hour).Err()
	if err != nil {
		return "", err
	}

	return response.Data.ID, nil
}
