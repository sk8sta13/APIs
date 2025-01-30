package database

import (
	"database/sql"
	"time"
)

type Order struct {
	db             *sql.DB
	Id             int32
	Payment_method string
	Total          float32
	Date_purchase  time.Time
	Status         string
	Created_at     time.Time
	Updated_at     time.Time
}

func NewOrder(db *sql.DB) *Order {
	return &Order{db: db}
}

func (o *Order) FindAll() ([]Order, error) {
	rows, err := o.db.Query("select id, payment_method, total, date_purchase, status, created_at, updated_at from orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []Order{}
	for rows.Next() {
		var id int32
		var total float32
		var payment_method, status string
		var date_purchase, created_at, updated_at sql.RawBytes

		if err := rows.Scan(&id, &payment_method, &total, &date_purchase, &status, &created_at, &updated_at); err != nil {
			return nil, err
		}

		orders = append(orders, Order{
			Id:             id,
			Payment_method: payment_method,
			Total:          total,
			Date_purchase:  parseTime(date_purchase),
			Status:         status,
			Created_at:     parseTime(created_at),
			Updated_at:     parseTime(updated_at),
		})
	}

	return orders, nil
}

func parseTime(data sql.RawBytes) time.Time {
	parsed, _ := time.Parse("2006-01-02 15:04:05", string(data))
	return parsed
}
