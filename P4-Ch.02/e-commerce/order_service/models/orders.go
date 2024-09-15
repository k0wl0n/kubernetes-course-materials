package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/k0wl0n/go-ecommerce-order/internal/database"
)

type OrderDetail struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Product    Product   `json:"product"`
}

type OrderList struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
}

func DatabaseOrderToOrder(dbOrder database.Order, dbProduct Product) OrderDetail {
	return OrderDetail{
		ID:         dbOrder.ID,
		CreatedAt:  dbOrder.CreatedAt,
		ModifiedAt: dbOrder.ModifiedAt,
		Product:    dbProduct,
	}
}

func DatabaseOrderToOrderList(dbOrders []database.GetOrdersRow) []OrderList {
	var orders []OrderList

	for _, dbOrder := range dbOrders {
		orders = append(orders, OrderList{
			ID:        dbOrder.ID,
			ProductID: dbOrder.ProductID,
		})
	}
	return orders
}

func OrderStructToByte(order OrderDetail) ([]byte, error) {
	return json.Marshal(order)
}

func ByteToOrderStruct(orderByte []byte) (OrderDetail, error) {
	var order OrderDetail
	err := json.Unmarshal(orderByte, &order)
	if err != nil {
		return order, err
	}
	return order, nil
}
