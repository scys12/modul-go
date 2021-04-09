package main

import (
	"net/http"

	"github.com/scys12/modul-go/internal/config"
	"github.com/scys12/modul-go/internal/database"
	prodDelivery "github.com/scys12/modul-go/internal/delivery/product"
	usrDelivery "github.com/scys12/modul-go/internal/delivery/user"
	prodRepo "github.com/scys12/modul-go/internal/persistence/product"
	usrRepo "github.com/scys12/modul-go/internal/persistence/user"
	productService "github.com/scys12/modul-go/internal/service/product"
	userService "github.com/scys12/modul-go/internal/service/user"
	"github.com/sirupsen/logrus"
)

type AppDelivery struct {
	UserDelivery    usrDelivery.IUserDelivery
	ProductDelivery prodDelivery.IProductDelivery
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
	userSvc := userService.NewInstance(userRepo)
	userDelivery := usrDelivery.NewInstance(userSvc)

	productRepo := prodRepo.NewInstance(db)
	productSvc := productService.NewInstance(productRepo)
	productDelivery := prodDelivery.NewInstance(productSvc)

	app := AppDelivery{userDelivery, productDelivery}

	routes := SetupRoutes(app)
	logrus.Fatal(http.ListenAndServe(":8000", routes))
}
