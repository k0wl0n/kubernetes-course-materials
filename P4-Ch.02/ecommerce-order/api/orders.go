// Contains order related API handlers

package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/k0wl0n/go-ecommerce-order/internal/database"
	"github.com/k0wl0n/go-ecommerce-order/models"
	"github.com/k0wl0n/go-ecommerce-order/shared"
)

func getUserID(c *gin.Context) (uuid.UUID, error) {
	userIDStr, isExists := c.Get("userID")
	if !isExists {
		return uuid.Nil, fmt.Errorf("authentication failed")
	}

	userID, err := uuid.Parse(fmt.Sprintf("%v", userIDStr))
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

// API for creating an order record
func (apiCfg *APIConfig) CreateOrder(c *gin.Context) {
	type Parameters struct {
		ProductID string `json:"product_id" binding:"required"`
	}
	var params Parameters

	err := c.ShouldBindJSON(&params)
	if err != nil {
		log.Errorln("Error caught while parsing order creation request data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error while parsing the request"})
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Errorln("Error caught while parsing user id: ", err)
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid credentails"})
		return
	}

	_, err = uuid.Parse(params.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product id format"})
		return
	}

	product, err := shared.GetProductIDToDetails(apiCfg.Cache, c, params.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product id"})
		return
	}

	dbOrder, err := CreateOrderDB(apiCfg, c, database.CreateOrderParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
		UserID:     userID,
		ProductID:  product.ID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	order := models.DatabaseOrderToOrder(dbOrder, product)

	// Call goroutine to save order details in cache
	go StoreOrderToCache(apiCfg.Cache, c, order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// API for getting order list
func (apiCfg *APIConfig) GetOrders(c *gin.Context) {
	// Parse string offset to integer
	offsetStr := c.Query("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Errorln("Error caught while parsing user id: ", err)
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid credentails"})
		return
	}

	dbOrders, err := GetOrderListDB(apiCfg, c, database.GetOrdersParams{
		UserID: userID,
		Limit:  10,
		Offset: int32(offset),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	orders := models.DatabaseOrderToOrderList(dbOrders)

	if len(orders) == 0 {
		c.JSON(http.StatusOK, gin.H{"results": []string{}})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": orders})
}

// API for getting details of a specific order by its ID
func (apiCfg *APIConfig) GetOrderDetail(c *gin.Context) {
	orderIDStr := c.Param("id")
	orderID, err := uuid.Parse(orderIDStr)

	if err != nil {
		log.Errorln("Error caught while parsing the order ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid order id format"})
		return
	}

	// Check and retrive order from cache
	cachedOrder, err := RetriveOrderFromCache(apiCfg.Cache, c, orderIDStr)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"data": cachedOrder})
		return
	}

	log.Errorln("Error caught while fetching the order details from cache: ", err)

	userID, err := getUserID(c)
	if err != nil {
		log.Errorln("Error caught while parsing user id: ", err)
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid credentails"})
		return
	}

	dbOrder, err := GetOrderDetailDB(apiCfg, c, database.GetOrderByIdParams{
		ID:     orderID,
		UserID: userID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Fetch product details
	product, err := shared.GetProductIDToDetails(apiCfg.Cache, c, dbOrder.ProductID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	order := models.OrderDetail{
		ID:         dbOrder.ID,
		CreatedAt:  dbOrder.CreatedAt,
		ModifiedAt: dbOrder.ModifiedAt,
		Product:    product,
	}

	// Call goroutine to save order details in cache
	go StoreOrderToCache(apiCfg.Cache, c, order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}

// API for deleting an order by its ID
func (apiCfg *APIConfig) DeleteOrder(c *gin.Context) {
	orderIDStr := c.Param("id")
	orderID, err := uuid.Parse(orderIDStr)

	if err != nil {
		log.Errorln("Error caught while parsing the order ID: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid order id format"})
		return
	}

	userID, err := getUserID(c)
	if err != nil {
		log.Errorln("Error caught while parsing user id: ", err)
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid credentails"})
		return
	}

	_, err = GetOrderDetailDB(apiCfg, c, database.GetOrderByIdParams{
		ID:     orderID,
		UserID: userID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err = DeleteOrderDB(apiCfg, c, database.DeleteOrderParams{
		ID:     orderID,
		UserID: userID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Call goroutine to remove order details from cache
	go DeleteOrderFromCache(apiCfg.Cache, c, orderIDStr)

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully!"})
}
