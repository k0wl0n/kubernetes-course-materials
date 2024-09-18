package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// API for getting details of a specific product
func (apiCfg *APIConfig) GetProductIDToDetails(c *gin.Context) {
	type Parameters struct {
		ID string `json:"id" binding:"required"`
	}

	var params Parameters
	err := c.ShouldBindJSON(&params)
	if err != nil {
		log.Errorln(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error while parsing the request data"})
		return
	}

	productID, err := uuid.Parse(params.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid product ID format"})
		return
	}

	dbProduct, err := GetProductDetailDB(apiCfg, c, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": DatabaseProductToProduct(dbProduct)})
}
