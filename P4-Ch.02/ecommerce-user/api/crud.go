// Contains Product CRUD Queries related function

package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/k0wl0n/go-ecommerce-user/internal/database"
)

// Get user by email from DB
func GetUserByEmailDB(apiCfg *APIConfig, ctx *gin.Context, userEmail string) (database.User, error) {
	user, err := apiCfg.Queries.GetUserByEmail(ctx, userEmail)
	if err != nil {
		log.Errorln("Error while getting user by email from DB: ", err)
		return database.User{}, fmt.Errorf("something went wrong")
	}
	return user, nil
}

// Get user by ID from DB
func GetUserByIDFromDB(apiCfg *APIConfig, ctx *gin.Context, userID uuid.UUID) (database.User, error) {
	user, err := apiCfg.Queries.GetUserById(ctx, userID)
	if err != nil {
		log.Errorln("Error while getting user by ID from DB: ", err)
		return database.User{}, fmt.Errorf("something went wrong")
	}
	return user, nil
}

// Update user details in DB
func UpdateUserDetailDB(apiCfg *APIConfig, ctx *gin.Context, params database.UpdateUserDetailsParams) (database.User, error) {
	// Begin DB transaction
	tx, err := apiCfg.DB.Begin()
	if err != nil {
		log.Fatal("Error caught while starting a transaction: ", err)
		return database.User{}, fmt.Errorf("something went wrong")
	}
	defer tx.Rollback()
	qtx := apiCfg.Queries.WithTx(tx)

	// Update user details
	user, err := qtx.UpdateUserDetails(ctx, params)

	if err != nil {
		log.Errorln("Error caught while updating user details in DB: ", err)
		return database.User{}, fmt.Errorf("something went wrong")
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal("Error caught while closing a transaction: ", err)
		return database.User{}, fmt.Errorf("something went wrong")
	}

	return user, nil
}

// Update user password in DB
func UpdateUserPasswordDB(apiCfg *APIConfig, ctx *gin.Context, params database.UpdateUserPasswordParams) error {
	// Begin DB transaction
	tx, err := apiCfg.DB.Begin()
	if err != nil {
		log.Fatal("Error caught while starting a transaction: ", err)
		return fmt.Errorf("something went wrong")
	}
	defer tx.Rollback()
	qtx := apiCfg.Queries.WithTx(tx)

	// Update user password
	err = qtx.UpdateUserPassword(ctx, params)

	if err != nil {
		log.Errorln("Error caught while updating user password in DB: ", err)
		return fmt.Errorf("something went wrong")
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal("Error caught while closing a transaction: ", err)
		return fmt.Errorf("something went wrong")
	}

	return nil
}

// Delete user from DB
func DeleteUserDB(apiCfg *APIConfig, ctx *gin.Context, userID uuid.UUID) error {
	// Begin DB transaction
	tx, err := apiCfg.DB.Begin()
	if err != nil {
		log.Fatal("Error caught while starting a transaction: ", err)
		return fmt.Errorf("something went wrong")
	}
	defer tx.Rollback()
	qtx := apiCfg.Queries.WithTx(tx)

	err = qtx.DeleteUser(ctx, userID)

	if err != nil {
		log.Errorln("Error caught while deleting user from DB: ", err)
		return fmt.Errorf("something went wrong")
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal("Error caught while closing a transaction: ", err)
		return fmt.Errorf("something went wrong")
	}

	return nil
}
