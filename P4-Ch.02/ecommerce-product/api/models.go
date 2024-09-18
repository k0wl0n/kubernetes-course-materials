package api

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/k0wl0n/go-ecommerce-product/internal/database"
)

type ProductDetail struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	Name        string    `json:"name"`
	Price       int32     `json:"price"`
	Description string    `json:"description"`
}

type productList struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Price       int32     `json:"price"`
	Description string    `json:"description"`
}

func DatabaseProductToProduct(dbProduct database.Product) ProductDetail {
	return ProductDetail{
		ID:          dbProduct.ID,
		CreatedAt:   dbProduct.CreatedAt,
		ModifiedAt:  dbProduct.ModifiedAt,
		Name:        dbProduct.Name,
		Price:       dbProduct.Price,
		Description: dbProduct.Description,
	}
}

func DatabaseProductToProductList(dbProducts []database.GetProductsRow) []productList {
	var products []productList

	for _, dbProduct := range dbProducts {
		products = append(products, productList{
			ID:          dbProduct.ID,
			Name:        dbProduct.Name,
			Price:       dbProduct.Price,
			Description: dbProduct.Description,
		})
	}
	return products
}

func ProductStructToByte(product ProductDetail) ([]byte, error) {
	return json.Marshal(product)
}

func ByteToProductStruct(productByte []byte) (ProductDetail, error) {
	var product ProductDetail

	err := json.Unmarshal(productByte, &product)
	if err != nil {
		return product, err
	}

	return product, nil
}
