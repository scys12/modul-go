package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scys12/modul-go/internal/middleware"
)

func SetupRoutes(appDelivery AppDelivery) *mux.Router {
	router := mux.NewRouter()

	router.Path("/products").Methods(http.MethodGet).HandlerFunc(appDelivery.ProductDelivery.GetAvailableProducts)

	cust := router.PathPrefix("/customer").Subrouter()
	cust.Methods(http.MethodPost).Path("/register").HandlerFunc(appDelivery.CustomerDelivery.CreateNewCustomer)
	cust.Methods(http.MethodPost).Path("/login").HandlerFunc(appDelivery.CustomerDelivery.LoginCustomer)
	cust.Handle("/me", middleware.TokenMiddleware(http.HandlerFunc(appDelivery.CustomerDelivery.GetUserDetail))).Methods(http.MethodGet)
	buy := cust.PathPrefix("/buy").Subrouter()
	buy.Use(middleware.TokenMiddleware)
	buy.Methods(http.MethodPost).Path("/{id}").HandlerFunc(appDelivery.ProductDelivery.BuyProduct)

	sell := router.PathPrefix("/seller").Subrouter()
	sell.Methods(http.MethodPost).Path("/register").HandlerFunc(appDelivery.SellerDelivery.CreateNewSeller)
	sell.Methods(http.MethodPost).Path("/login").HandlerFunc(appDelivery.SellerDelivery.LoginSeller)
	sell.Handle("/me", middleware.TokenMiddleware(http.HandlerFunc(appDelivery.CustomerDelivery.GetUserDetail))).Methods(http.MethodGet)
	sellProduct := router.PathPrefix("/{id}/products")
	sellProduct.Methods(http.MethodGet).HandlerFunc(appDelivery.ProductDelivery.GetSellerProducts)

	prod := router.PathPrefix("/product").Subrouter()
	prod.Methods(http.MethodGet).Path("").HandlerFunc(appDelivery.ProductDelivery.GetProduct)
	prod.Use(middleware.TokenMiddleware)
	prod.Use(middleware.IsRoleSeller)
	prod.Methods(http.MethodPost).Path("").HandlerFunc(appDelivery.ProductDelivery.InsertProduct)
	prod.Methods(http.MethodPut).Path("").HandlerFunc(appDelivery.ProductDelivery.UpdateProduct)
	prod.Methods(http.MethodDelete).Path("").HandlerFunc(appDelivery.ProductDelivery.DeleteProduct)

	return router
}
