package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	Name        string    `json:"name"`
	Price       int32     `json:"price"`
	Description string    `json:"description"`
}

func ProductStructToByte(product Product) ([]byte, error) {
	return json.Marshal(product)
}

func ByteToProductStruct(productByte []byte) (Product, error) {
	var product Product
	err := json.Unmarshal(productByte, &product)
	if err != nil {
		return product, err
	}
	return product, nil
}
