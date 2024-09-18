// Contains Product CRUD Queries related function

package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/k0wl0n/go-ecommerce-product/internal/database"
)

// Create product in DB
func CreateProductDB(apiCfg *APIConfig, ctx *gin.Context, params database.CreateProductParams) (database.Product, error) {
	// Begin DB transaction
	tx, err := apiCfg.DB.Begin()
	if err != nil {
		log.Fatal(err)
		return database.Product{}, fmt.Errorf("something went wrong")
	}
	defer tx.Rollback()
	qtx := apiCfg.Queries.WithTx(tx)

	// Create product
	dbProduct, err := qtx.CreateProduct(ctx, params)

	if err != nil {
		log.Errorln(err)
		return database.Product{}, fmt.Errorf("something went wrong")
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return database.Product{}, fmt.Errorf("something went wrong")
	}

	return dbProduct, nil
}

// Get list of products from DB
func GetProductListDB(apiCfg *APIConfig, ctx *gin.Context, params database.GetProductsParams) ([]database.GetProductsRow, error) {
	products, err := apiCfg.Queries.GetProducts(ctx, params)
	if err != nil {
		return []database.GetProductsRow{}, fmt.Errorf("something went wrong")
	}
	return products, nil
}

// Get details of a specific product from DB
func GetProductDetailDB(apiCfg *APIConfig, ctx *gin.Context, productID uuid.UUID) (database.Product, error) {
	product, err := apiCfg.Queries.GetProductById(ctx, productID)
	if err != nil {
		log.Errorln(err)
		return database.Product{}, fmt.Errorf("something went wrong")
	}

	return product, nil
}

// Update product details in DB
func UpdateProductDetailDB(apiCfg *APIConfig, ctx *gin.Context, params database.UpdateProductDetailsParams) (database.Product, error) {
	// Begin DB transaction
	tx, err := apiCfg.DB.Begin()
	if err != nil {
		log.Fatal(err)
		return database.Product{}, fmt.Errorf("something went wrong")
	}
	defer tx.Rollback()
	qtx := apiCfg.Queries.WithTx(tx)

	// Update product
	product, err := qtx.UpdateProductDetails(ctx, params)

	if err != nil {
		log.Errorln(err)
		return database.Product{}, fmt.Errorf("something went wrong")
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return database.Product{}, fmt.Errorf("something went wrong")
	}

	return product, nil
}

// Delete product from DB
func DeleteProductDetailDB(apiCfg *APIConfig, ctx *gin.Context, productID uuid.UUID) error {
	// Begin DB transaction
	tx, err := apiCfg.DB.Begin()
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("something went wrong")
	}
	defer tx.Rollback()
	qtx := apiCfg.Queries.WithTx(tx)

	// Delete product
	err = qtx.DeleteProduct(ctx, productID)
	if err != nil {
		log.Errorln(err)
		return fmt.Errorf("something went wrong")
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("something went wrong")
	}

	return nil
}
