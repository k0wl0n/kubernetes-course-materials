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
	"github.com/k0wl0n/go-ecommerce-product/internal/database"
)

var (
	apiCfg    *APIConfig
	productID uuid.UUID
)

func TestMain(m *testing.M) {
	// Connect to testing DB
	conn := GetDBConn("postgres://db_user:1234@localhost:5432/products_test_db?sslmode=disable")

	apiCfg = &APIConfig{
		DB:      conn,
		Queries: database.New(conn),
	}

	exitCode := m.Run()

	conn.Close()
	os.Exit(exitCode)
}

func TestCreateProduct(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	// Create a testing product
	product, err := CreateProductDB(apiCfg, ctx, database.CreateProductParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		ModifiedAt:  time.Now().UTC(),
		Name:        "Testing product",
		Price:       10,
		Description: "Random text",
		CreatorID:   uuid.New(),
	})

	if err != nil {
		t.Errorf("Error caught while creating a product: %s", err.Error())
	}

	productID = product.ID
}

func TestConcurrentProductDeletion(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	// Counter to track successful product deletion
	var successCounter int
	var mu sync.Mutex // Mutex to protect the counter

	var wg sync.WaitGroup

	// Number of concurrent signups
	numGoroutines := 5

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Delete the product
			err := DeleteProductDetailDB(apiCfg, ctx, productID)

			if err != nil {
				// Increment the counter if product deletion is successful
				mu.Lock()
				successCounter++
				mu.Unlock()
			}
		}()
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	if successCounter > 1 {
		t.Errorf("Expected only one product to be deleted, but got %d", successCounter)
	}
}
