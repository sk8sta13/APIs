package service

import (
	"context"

	"github.com/sk8sta13/APIs/Internal/database"
	"github.com/sk8sta13/APIs/Internal/pb"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	db database.Order
}

func NewCategoryService(db database.Order) *OrderService {
	return &OrderService{
		db: db,
	}
}

func (o *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.Orders, error) {
	rows, err := o.db.FindAll()
	if err != nil {
		return nil, err
	}

	var orders []*pb.Order
	for _, row := range rows {
		orders = append(orders, &pb.Order{
			Id:            row.Id,
			PaymentMethod: row.Payment_method,
			Total:         row.Total,
			DatePurchase:  row.Date_purchase.String(),
			Status:        row.Status,
			CreatedAt:     row.Created_at.String(),
			UpdatedAt:     row.Updated_at.String(),
		})
	}

	return &pb.Orders{List: orders}, nil
}
