package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0wl0n/go-ecommerce-user/internal/database"
	"github.com/k0wl0n/go-ecommerce-user/utils"
	"github.com/k0wl0n/go-ecommerce-user/validators"
	log "github.com/sirupsen/logrus"
)

// SignUp API
func (apiCfg *APIConfig) Singup(c *gin.Context) {
	type Parameters struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var params Parameters
	err := c.ShouldBindJSON(&params)

	if err != nil {
		log.Errorln("Error caught while parsing signup request data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error while parsing the request data"})
		return
	}

	// Make the password case sensitive
	email := strings.ToLower(params.Email)

	// Validate Password
	err = validators.PasswordValidator(params.Password, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Generate hashed password
	hashedPassword, err := utils.GetHashedPassword(params.Password)

	if err != nil {
		log.Errorln("Error caught while generating hashed password: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	// Check if user exists with the given email address
	_, err = GetUserByEmailDB(apiCfg, c, email)
	if err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User with this email address already exists"})
		return
	}

	// Begin DB transaction
	tx, err := apiCfg.DB.Begin()
	if err != nil {
		log.Fatal("Error caught while starting a transaction: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}
	defer tx.Rollback()
	qtx := apiCfg.Queries.WithTx(tx)

	// Create user account
	dbUser, err := qtx.CreateUser(c, database.CreateUserParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		ModifiedAt: time.Now().UTC(),
		Email:      email,
		Password:   hashedPassword,
	})

	if err != nil {
		log.Errorln("Error caught while creating a user in DB: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	// Generate auth tokens for the user
	tokens, err := utils.GenerateTokens(dbUser.ID.String())

	if err != nil {
		log.Errorln("Error caught while generating auth tokens during signup: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	// Store user detail in cache
	user := DatabaseUserToUser(dbUser)
	go StoreUserToCache(apiCfg.Cache, c, user)

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal("Error caught while closing a transaction: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Account created successfully!", "data": tokens})
}

// Login API
func (apiCfg *APIConfig) Login(c *gin.Context) {
	type Parameters struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var params Parameters
	err := c.ShouldBindJSON(&params)

	if err != nil {
		log.Errorln("Error caught while parsing login request data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error while parsing the request data"})
		return
	}

	// Check wheather the user exists with the given email or not
	user, err := GetUserByEmailDB(apiCfg, c, strings.ToLower(params.Email))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User does not exists, Please check your credentials"})
		return
	}

	// Check the given password with hashed password stored in DB
	match, err := utils.CheckPassowrdValid(params.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid password !"})
		return
	} else if !match {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid password !"})
		return
	}

	// Generate auth tokens for the user
	tokens, err := utils.GenerateTokens(user.ID.String())

	if err != nil {
		log.Errorln("Error caught while generating auth tokens during login: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged in Successfully!", "data": tokens})
}

// Refresh Token API
//
// Generate new tokens if the given refresh token is valid
func (apiCfg *APIConfig) RefreshAccessToken(c *gin.Context) {
	type Parameters struct {
		RefreshToken string `json:"refresh_token"`
	}

	var params Parameters
	err := c.ShouldBindJSON(&params)

	if err != nil {
		log.Errorln("Error caught while parsing refresh token request data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error while parsing the request data"})
		return
	}

	tokens, err := utils.ReIssueAccessToken(params.RefreshToken)
	if err != nil {
		log.Errorln("Error caught while re-issuing auth tokens: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error while issueing new tokens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tokens re-issued Successfully!", "data": tokens})
}
