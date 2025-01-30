package graph

import "github.com/sk8sta13/APIs/Internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	OrderDb *database.Order
}
