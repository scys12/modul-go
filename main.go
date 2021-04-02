package main

import (
	"net/http"

	"github.com/scys12/modul-go/internal/config"
	"github.com/scys12/modul-go/internal/database"
	custDelivery "github.com/scys12/modul-go/internal/delivery/customer"
	prodDelivery "github.com/scys12/modul-go/internal/delivery/product"
	selDelivery "github.com/scys12/modul-go/internal/delivery/seller"
	custRepo "github.com/scys12/modul-go/internal/persistence/customer"
	prodRepo "github.com/scys12/modul-go/internal/persistence/product"
	selRepo "github.com/scys12/modul-go/internal/persistence/seller"
	"github.com/sirupsen/logrus"
)

type AppDelivery struct {
	SellerDelivery   selDelivery.SellerDelivery
	CustomerDelivery custDelivery.CustomerDelivery
	ProductDelivery  prodDelivery.ProductDelivery
}

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		logrus.Error("Configuration file have a problem")
	}
	db, err := database.InitDatabase(cfg)
	if err != nil {
		logrus.Error(err)
	}
	testPing := db.Ping()
	if testPing != nil {
		logrus.Error(testPing)
	}
	defer db.Close()
	customerRepo := custRepo.NewInstance(db)
	customerDelivery := custDelivery.NewInstance(customerRepo)

	sellerRepo := selRepo.NewInstance(db)
	sellerDelivery := selDelivery.NewInstance(sellerRepo)

	productRepo := prodRepo.NewInstance(db)
	productDelivery := prodDelivery.NewInstance(productRepo)

	appDelivery := AppDelivery{
		SellerDelivery:   sellerDelivery,
		CustomerDelivery: customerDelivery,
		ProductDelivery:  productDelivery,
	}
	routes := SetupRoutes(appDelivery)
	logrus.Fatal(http.ListenAndServe(":8000", routes))
}
