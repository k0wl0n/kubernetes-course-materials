// Contains all product related API handlers

package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/k0wl0n/go-ecommerce-product/internal/database"
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

// API for creating product record in DB
func (apiCfg *APIConfig) CreateProduct(c *gin.Context) {
	type Parameters struct {
		Name        string `json:"name" binding:"required"`
		Price       int32  `json:"price" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	var params Parameters
	err := c.ShouldBindJSON(&params)

	if err != nil {
		log.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error while parsing the request data"})
		return
	}

	// Fetch userID from the context
	userID, err := getUserID(c)
	if err != nil {
		log.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	dbProduct, err := CreateProductDB(apiCfg, c, database.CreateProductParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		ModifiedAt:  time.Now().UTC(),
		Name:        params.Name,
		Price:       params.Price,
		Description: params.Description,
		CreatorID:   userID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	product := DatabaseProductToProduct(dbProduct)

	// Call goroutine and save product into cache
	go StoreProductToCache(apiCfg.Cache, c, product)

	c.JSON(http.StatusCreated, gin.H{"message": "Product Created Successfully!", "data": product})
}

// API for getting details of a specific product by ID
func (apiCfg *APIConfig) GetProductDetails(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := uuid.Parse(productIDStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID format"})
		return
	}

	// Retrive product details from cache
	cachedProduct, err := RetriveProductFromCache(apiCfg.Cache, c, productIDStr)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"data": cachedProduct})
		return
	}

	log.Errorln("Error caught while fetching the product details from cache: ", err)

	dbProduct, err := GetProductDetailDB(apiCfg, c, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	product := DatabaseProductToProduct(dbProduct)

	// Call goroutine and save product into cache
	go StoreProductToCache(apiCfg.Cache, c, product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// API for updating product details
func (apiCfg *APIConfig) UpdateProductDetails(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := uuid.Parse(productIDStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
		return
	}

	// Fetch userID from the context
	userID, err := getUserID(c)
	if err != nil {
		log.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	dbProduct, err := GetProductDetailDB(apiCfg, c, productID)
	if err != nil {
		log.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	// Check if the current user is the product creator or not
	if dbProduct.CreatorID != userID {
		c.JSON(http.StatusForbidden, gin.H{"message": "You cannot update this product details"})
		return
	}

	type Parameters struct {
		Name        string `json:"name"`
		Price       int32  `json:"price"`
		Description string `json:"description"`
	}

	var params Parameters
	err = c.ShouldBindJSON(&params)

	if err != nil {
		log.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error while parsing the request data"})
		return
	}

	if params.Name == "" && params.Price == 0 && params.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No request data is passed for update"})
		return
	}

	// Pre-fill data if some key is empty
	if params.Name == "" {
		params.Name = dbProduct.Name
	}

	if params.Description == "" {
		params.Description = dbProduct.Description
	}

	if params.Price == 0 {
		params.Price = dbProduct.Price
	}

	dbProduct, err = UpdateProductDetailDB(apiCfg, c, database.UpdateProductDetailsParams{
		ModifiedAt:  time.Now().UTC(),
		ID:          productID,
		Name:        params.Name,
		Price:       params.Price,
		Description: params.Description,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	product := DatabaseProductToProduct(dbProduct)

	// Call goroutine and save new product details into cache
	go StoreProductToCache(apiCfg.Cache, c, product)

	c.JSON(http.StatusOK, gin.H{"message": "Product details updated Successfully!", "data": product})
}

// API for deleting product details
func (apiCfg *APIConfig) DeleteProduct(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := uuid.Parse(productIDStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID"})
		return
	}

	// Fetch userID from the context
	userID, err := getUserID(c)
	if err != nil {
		log.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	dbProduct, err := GetProductDetailDB(apiCfg, c, productID)
	if err != nil {
		log.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	// Check if the current user is the product creator or not
	if dbProduct.CreatorID != userID {
		c.JSON(http.StatusForbidden, gin.H{"message": "You cannot delete this product details"})
		return
	}

	err = DeleteProductDetailDB(apiCfg, c, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Call goroutine and delete product details from cache
	go DeleteProductFromCache(apiCfg.Cache, c, productIDStr)

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully!"})
}

// API for getting list of products
func (apiCfg *APIConfig) GetProducts(c *gin.Context) {
	// Parse string offset to integer
	offsetStr := c.Query("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	dbProducts, err := GetProductListDB(apiCfg, c, database.GetProductsParams{
		Limit:  10,
		Offset: int32(offset),
	})

	if err != nil {
		log.Errorln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	products := DatabaseProductToProductList(dbProducts)

	if len(products) == 0 {
		c.JSON(http.StatusOK, gin.H{"results": []string{}})
		return
	}
	c.JSON(http.StatusOK, gin.H{"results": DatabaseProductToProductList(dbProducts)})
}
