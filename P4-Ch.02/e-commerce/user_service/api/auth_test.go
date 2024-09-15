package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/k0wl0n/go-ecommerce-user/internal/database"
)

var apiCfg *APIConfig

func TestMain(m *testing.M) {
	// Connect to testing DB
	conn := GetDBConn("postgres://db_user:1234@localhost:5432/users_test_db?sslmode=disable")

	apiCfg = &APIConfig{
		DB:      conn,
		Queries: database.New(conn),
	}

	exitCode := m.Run()

	conn.Close()
	os.Exit(exitCode)
}

func TestSignup(t *testing.T) {
	engine := gin.Default()
	engine.POST("/register/", apiCfg.Singup)

	payload := []byte(`{"email": "john.doe@example.com", "password": "12345678Uu@"}`)

	// Create a HTTP request
	req := httptest.NewRequest(http.MethodPost, "/register/", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder
	rr := httptest.NewRecorder()
	engine.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusCreated {
		t.Errorf("Error while creating user: %v", rr.Body.String())
	}
}

func TestConcurrentSingup(t *testing.T) {
	t.Parallel()

	engine := gin.Default()
	engine.POST("/register/", apiCfg.Singup)

	payload := []byte(`{"email": "john.doe@example1.com", "password": "12345678Uu@"}`)

	// Counter to track successful signups
	var successCounter int
	var mu sync.Mutex // Mutex to protect the counter

	var wg sync.WaitGroup

	// Number of concurrent signups
	numGoroutines := 5

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Create a request for each Goroutine
			req := httptest.NewRequest(http.MethodPost, "/register/", bytes.NewBuffer(payload))
			req.Header.Set("Content-Type", "application/json")

			// Create a ResponseRecorder to capture the response
			rr := httptest.NewRecorder()

			// Call the signup handler
			engine.ServeHTTP(rr, req)

			// Check the status code for each Goroutine
			if rr.Code == http.StatusOK {
				// Increment the counter if signup is successful
				mu.Lock()
				successCounter++
				mu.Unlock()
			}
		}()
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	if successCounter > 1 {
		t.Errorf("Expected only one account to be created, but got %d", successCounter)
	}
}
