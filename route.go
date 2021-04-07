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

	cust.Path("/me").Methods(http.MethodGet).Handler(
		middleware.TokenMiddleware(
			middleware.IsRoleCustomer(
				http.HandlerFunc(appDelivery.CustomerDelivery.GetUserDetail))))

	cust.Methods(http.MethodPost).Path("/buy/{id}").Handler(
		middleware.TokenMiddleware(
			middleware.IsRoleCustomer(
				http.HandlerFunc(appDelivery.ProductDelivery.BuyProduct))))

	sell := router.PathPrefix("/seller").Subrouter()
	sell.Methods(http.MethodPost).Path("/register").HandlerFunc(appDelivery.SellerDelivery.CreateNewSeller)
	sell.Methods(http.MethodPost).Path("/login").HandlerFunc(appDelivery.SellerDelivery.LoginSeller)

	sell.Path("/me").Methods(http.MethodGet).Handler(
		middleware.TokenMiddleware(
			middleware.IsRoleSeller(
				http.HandlerFunc(appDelivery.SellerDelivery.GetUserDetail))))

	sell.Methods(http.MethodGet).Path("/{id}/products").HandlerFunc(appDelivery.ProductDelivery.GetSellerProducts)

	prod := router.PathPrefix("/product").Subrouter()

	authProd := prod.NewRoute().Subrouter()
	authProd.Use(middleware.TokenMiddleware)
	authProd.Use(middleware.IsRoleSeller)

	authProd.Methods(http.MethodPut).Path("/{id}").HandlerFunc(appDelivery.ProductDelivery.UpdateProduct)
	authProd.Methods(http.MethodDelete).Path("/{id}").HandlerFunc(appDelivery.ProductDelivery.DeleteProduct)
	authProd.Methods(http.MethodPost).HandlerFunc(appDelivery.ProductDelivery.InsertProduct)

	prod.Methods(http.MethodGet).Path("/{id}").HandlerFunc(appDelivery.ProductDelivery.GetProduct)

	return router
}
