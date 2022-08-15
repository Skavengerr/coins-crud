package main

import (
	"log"
	"net/http"

	"Skavengerr/coins-crud/internal/repository/psql"
	"Skavengerr/coins-crud/internal/service"
	"Skavengerr/coins-crud/internal/transport/rest"
	"Skavengerr/coins-crud/pkg/config"
	"Skavengerr/coins-crud/pkg/server"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.InitViper("..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db := server.ConnectDB(&cfg)

	defer db.Close()

	coinsRepo := psql.NewCoins(db)
	coinsService := service.NewCoins(coinsRepo)
	handler := rest.NewHandler(coinsService)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
