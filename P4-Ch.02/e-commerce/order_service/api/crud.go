// Contains Order CRUD queries related functions

package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/k0wl0n/go-ecommerce-order/internal/database"
)

// Create an order record in DB
func CreateOrderDB(apiCfg *APIConfig, ctx *gin.Context, params database.CreateOrderParams) (database.Order, error) {
	// Begin DB transaction
	tx, err := apiCfg.DB.Begin()
	if err != nil {
		log.Fatal("Error caught while initiating the transaction: ", err)
		return database.Order{}, fmt.Errorf("something went wrong")
	}
	defer tx.Rollback()
	qtx := apiCfg.Queries.WithTx(tx)

	// Create the order
	dbOrder, err := qtx.CreateOrder(ctx, params)

	if err != nil {
		log.Errorln("Error caught while creating the order in DB", err)
		return database.Order{}, fmt.Errorf("something went wrong")
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal("Error caught while committing the transaction; ", err)
		return database.Order{}, fmt.Errorf("something went wrong")
	}

	return dbOrder, nil
}

// Fetch order list from DB
func GetOrderListDB(apiCfg *APIConfig, ctx *gin.Context, params database.GetOrdersParams) ([]database.GetOrdersRow, error) {
	dbOrders, err := apiCfg.Queries.GetOrders(ctx, params)
	if err != nil {
		log.Errorln("Error caught while fetching order list: ", err)
		return []database.GetOrdersRow{}, fmt.Errorf("something went wrong")
	}
	return dbOrders, nil
}

// Fetch order details from DB
func GetOrderDetailDB(apiCfg *APIConfig, ctx *gin.Context, params database.GetOrderByIdParams) (database.GetOrderByIdRow, error) {
	dbOrder, err := apiCfg.Queries.GetOrderById(ctx, params)
	if err != nil {
		log.Errorln("Error caught while fetching order details: ", err)
		return database.GetOrderByIdRow{}, fmt.Errorf("something went wrong")
	}
	return dbOrder, nil
}

func DeleteOrderDB(apiCfg *APIConfig, ctx *gin.Context, params database.DeleteOrderParams) error {
	// Begin DB transaction
	tx, err := apiCfg.DB.Begin()
	if err != nil {
		log.Fatal("Error caught while initiating the transaction: ", err)
		return fmt.Errorf("something went wrong")
	}
	defer tx.Rollback()
	qtx := apiCfg.Queries.WithTx(tx)

	// Delete the order
	err = qtx.DeleteOrder(ctx, params)
	if err != nil {
		log.Errorln("Error caught while deleting order details: ", err)
		return fmt.Errorf("something went wrong")
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal("Error caught while committing the transaction: ", err)
		return fmt.Errorf("something went wrong")
	}

	return nil
}
