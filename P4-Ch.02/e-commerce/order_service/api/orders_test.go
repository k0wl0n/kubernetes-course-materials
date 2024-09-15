package api

import (
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/k0wl0n/go-ecommerce-order/internal/database"
)

var (
	apiCfg  *APIConfig
	orderID uuid.UUID
	userID  uuid.UUID
)

func TestMain(m *testing.M) {
	// Connect to testing DB
	conn := GetDBConn("postgres://db_user:1234@localhost:5432/orders_test_db?sslmode=disable")

	apiCfg = &APIConfig{
		DB:      conn,
		Queries: database.New(conn),
	}

	exitCode := m.Run()

	conn.Close()
	os.Exit(exitCode)
}

func TestCreateOrder(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	// Create a testing order
	userID = uuid.New()

	order, err := CreateOrderDB(apiCfg, ctx, database.CreateOrderParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
		ProductID:  uuid.New(),
		UserID:     userID,
	})

	if err != nil {
		t.Errorf("Error caught while creating an order: %s", err.Error())
	}

	orderID = order.ID
}

func TestConcurrentOrderDeletion(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	// Counter to track successful order deletion
	var successCounter int
	var mu sync.Mutex // Mutex to protect the counter

	var wg sync.WaitGroup

	// Number of concurrent order deletion
	numGoroutines := 5

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Delete the order
			err := DeleteOrderDB(apiCfg, ctx, database.DeleteOrderParams{
				ID:     orderID,
				UserID: userID,
			})

			if err != nil {
				// Increment the counter if order deletion is successful
				mu.Lock()
				successCounter++
				mu.Unlock()
			}
		}()
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	if successCounter > 1 {
		t.Errorf("Expected only one order record to be deleted, but got %d", successCounter)
	}
}
