// Custom user model for converting the raw data into a desired JSON data
//
// With keys as formatted as snake_case rather than TitleCase

package api

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/k0wl0n/go-ecommerce-user/internal/database"
)

type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func DatabaseUserToUser(dbUser database.User) User {
	return User{
		ID:    dbUser.ID,
		Name:  dbUser.Name.String,
		Email: dbUser.Email,
	}
}

func UserStructToByte(user User) ([]byte, error) {
	return json.Marshal(user)
}

func ByteToUserStruct(userByte []byte) (User, error) {
	var user User

	err := json.Unmarshal(userByte, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
