package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/k0wl0n/go-ecommerce-user/internal/database"
	"github.com/k0wl0n/go-ecommerce-user/utils"
	"github.com/k0wl0n/go-ecommerce-user/validators"
)

// Common function for getting user object from the context, For all handlers in this module
func getUserFromCtx(c *gin.Context) (User, error) {
	value, exists := c.Get("user")

	if !exists {
		return User{}, fmt.Errorf("authentication required")
	}

	user, ok := value.(User)

	if !ok {
		return User{}, fmt.Errorf("invalid user")
	}

	return user, nil
}

// Fetch user profile details
func (apiCfg *APIConfig) GetUserProfile(c *gin.Context) {
	user, err := getUserFromCtx(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Update user profile details
func (apiCfg *APIConfig) UpdateUserProfile(c *gin.Context) {
	user, err := getUserFromCtx(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}

	type Parameters struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	var params Parameters
	err = c.ShouldBindJSON(&params)

	if err != nil {
		log.Errorln("Error caught while parsing update user detail API request data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error while parsing the request data"})
		return
	}

	// check if request body is empty
	if params.Name == "" && params.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
		return
	}

	// Pre-fill name and email address with existing data from DB
	if params.Name == "" {
		params.Name = user.Name
	}

	isEmailChanged := true
	if params.Email == "" {
		params.Email = user.Email
		isEmailChanged = false
	}

	email := strings.ToLower(params.Email)

	// Validate the given email address
	if !validators.EmailValidator(params.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email address"})
		return
	}

	// Check if user any exists with the new email address
	if isEmailChanged && email != user.Email {
		_, err = GetUserByEmailDB(apiCfg, c, email)
		if err == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "User with this email address already exists"})
			return
		}
	}

	// Update the profile details
	dbUser, err := UpdateUserDetailDB(apiCfg, c, database.UpdateUserDetailsParams{
		Name: sql.NullString{
			String: params.Name,
			Valid:  true,
		},
		Email:      email,
		ModifiedAt: time.Now().UTC(),
		ID:         user.ID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	// Call goroutine to save user details into cache
	user = DatabaseUserToUser(dbUser)
	go StoreUserToCache(apiCfg.Cache, c, user)

	c.JSON(http.StatusOK, gin.H{"message": "Profile details updated successfully!", "data": user})
}

// Delete user profile
func (apiCfg *APIConfig) DeleteUserProfile(c *gin.Context) {
	user, err := getUserFromCtx(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}

	err = DeleteUserDB(apiCfg, c, user.ID)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Something went wrong"})
		return
	}

	// Delete user details from the cache
	go DeleteUserFromCache(apiCfg.Cache, c, user.ID.String())

	c.JSON(http.StatusOK, gin.H{"message": "Profile deleted successfully!"})
}

// Change password API for authenticated user
func (apiCfg *APIConfig) ChangePassword(c *gin.Context) {
	user, err := getUserFromCtx(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}

	type Parameters struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	var params Parameters
	err = c.ShouldBindJSON(&params)

	if err != nil {
		log.Errorln("Error caught while parsing change password API request data: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error while parsing the request data"})
		return
	}

	// Fetch user by ID from DB
	dbUser, err := GetUserByIDFromDB(apiCfg, c, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	// Check wheather or not old password is correct or not
	match, err := utils.CheckPassowrdValid(params.OldPassword, dbUser.Password)
	if err != nil {
		log.Errorln("Error caught while checking current password: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid old password, Please try again."})
		return
	} else if !match {
		log.Errorln("Error caught while checking current password: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid old password, Please try again."})
		return
	}

	if params.OldPassword == params.NewPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "New password should not be same as old password"})
		return
	}

	// Validate the new password
	if err = validators.PasswordValidator(params.NewPassword, dbUser.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Generate the new hashed password
	hashedPassword, err := utils.GetHashedPassword(params.NewPassword)

	if err != nil {
		log.Errorln("Error caught while generating hashed password: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}

	// Update the password
	err = UpdateUserPasswordDB(apiCfg, c, database.UpdateUserPasswordParams{
		Password:   hashedPassword,
		ModifiedAt: time.Now().UTC(),
		ID:         dbUser.ID,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully!"})
}
