package main

import (
	"net/http"

	"github.com/scys12/modul-go/internal/config"
	"github.com/scys12/modul-go/internal/database"
	prodDelivery "github.com/scys12/modul-go/internal/delivery/product"
	usrDelivery "github.com/scys12/modul-go/internal/delivery/user"
	prodRepo "github.com/scys12/modul-go/internal/persistence/product"
	usrRepo "github.com/scys12/modul-go/internal/persistence/user"
	"github.com/sirupsen/logrus"
)

type AppDelivery struct {
	UserDelivery    usrDelivery.UserDelivery
	ProductDelivery prodDelivery.ProductDelivery
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

	userRepo := usrRepo.NewInstance(db)
	userDelivery := usrDelivery.NewInstance(userRepo)

	productRepo := prodRepo.NewInstance(db)
	productDelivery := prodDelivery.NewInstance(productRepo)

	appDelivery := AppDelivery{
		UserDelivery:    userDelivery,
		ProductDelivery: productDelivery,
	}
	routes := SetupRoutes(appDelivery)
	logrus.Fatal(http.ListenAndServe(":8000", routes))
}
