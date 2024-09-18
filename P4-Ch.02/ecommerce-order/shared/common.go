package shared

import (
	"fmt"
	"net/http"
	"os"
)

func GetBaseURL(request *http.Request, host, endpoint string) string {
	protocol := "http"
	if request.TLS != nil {
		protocol += "s"
	}

	return fmt.Sprintf("%s://%s/internal/v1/%s/", protocol, host, endpoint)
}

func GetAPISecretKey() (string, error) {
	apiSecretKey := os.Getenv("INTERNAL_API_SECRET")

	if apiSecretKey == "" {
		return "", fmt.Errorf("API secret is not configured")
	}

	return apiSecretKey, nil
}

func GetUserServiceHost() (string, error) {
	userServiceHost := os.Getenv("USER_SERVICE_HOST")
	if userServiceHost == "" {
		return "", fmt.Errorf("user service host is not configured")
	}
	return userServiceHost, nil
}

func GetProductServiceHost() (string, error) {
	productServiceHost := os.Getenv("PRODUCT_SERVICE_HOST")
	if productServiceHost == "" {
		return "", fmt.Errorf("product service host is not configured")
	}
	return productServiceHost, nil
}
