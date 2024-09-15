// This module contains all the API which is to be used by other micro-services

package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/k0wl0n/go-ecommerce-user/utils"
)

func (apiCfg *APIConfig) GetUserFromToken(c *gin.Context) {
	type Parameters struct {
		Token string `json:"token" binding:"required"`
	}
	var params Parameters
	err := c.ShouldBindJSON(&params)

	if err != nil {
		log.Errorln("Error caught while parsing internal API request data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error while parsing the request data"})
		return
	}

	// Check token claim
	claims, err := utils.VerifyToken(params.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid authentication token"})
		return
	}

	// Check the validity of the token
	if !time.Unix(claims.ExpiresAt.Unix(), 0).After(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Authentication token is expired"})
		return
	}

	// Convert the userID string to UUID
	userID, err := uuid.Parse(claims.Data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid authentication token"})
		return
	}

	// Fetch the user by the ID
	dbUser, err := GetUserByIDFromDB(apiCfg, c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": DatabaseUserToUser(dbUser)})
}
