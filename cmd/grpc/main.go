package main

import (
	"database/sql"
	"fmt"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sk8sta13/APIs/Internal/database"
	"github.com/sk8sta13/APIs/Internal/pb"
	"github.com/sk8sta13/APIs/Internal/service"
	"github.com/sk8sta13/APIs/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	orderSv := service.NewCategoryService(*orderDb)

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, orderSv)

	reflection.Register(grpcServer) // Necessário para execusão de testes

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
