package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sk8sta13/APIs/Internal/database"
	"github.com/sk8sta13/APIs/Internal/infra/web/handlers"
	"github.com/sk8sta13/APIs/config"
)

func main() {
	cfgDb, err := config.LoadConfigs(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(cfgDb.Drive, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfgDb.User, cfgDb.Pass, cfgDb.Host, cfgDb.Port, cfgDb.Name))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderDb := database.NewOrder(db)
	orderHandler := handlers.NewProductHandler(*orderDb)

	route := chi.NewRouter()
	route.Use(middleware.Logger)
	route.Get("/orders", orderHandler.GetOrders)

	http.ListenAndServe(":8000", route)
}
