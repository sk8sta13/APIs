package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sk8sta13/APIs/Internal/database"
)

type OrderHandler struct {
	OrderDB database.Order
}

func NewProductHandler(order database.Order) *OrderHandler {
	return &OrderHandler{
		OrderDB: order,
	}
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.OrderDB.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}
