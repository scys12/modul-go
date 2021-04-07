package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scys12/modul-go/internal/middleware"
)

func SetupRoutes(appDelivery AppDelivery) *mux.Router {
	router := mux.NewRouter()

	usr := router.PathPrefix("/user").Subrouter()
	usr.Methods(http.MethodPost).Path("/register").HandlerFunc(appDelivery.UserDelivery.CreateNewUser)
	usr.Methods(http.MethodPost).Path("/login").HandlerFunc(appDelivery.UserDelivery.LoginUser)
	usr.Path("/me").Methods(http.MethodGet).Handler(middleware.TokenMiddleware(http.HandlerFunc(appDelivery.UserDelivery.GetUserDetail)))
	usr.Methods(http.MethodPost).Path("/buy/{id}").Handler(middleware.TokenMiddleware(http.HandlerFunc(appDelivery.ProductDelivery.BuyProduct)))
	usr.Methods(http.MethodGet).Path("/{id}/products").HandlerFunc(appDelivery.ProductDelivery.GetSellerProducts)

	router.Path("/products").Methods(http.MethodGet).HandlerFunc(appDelivery.ProductDelivery.GetAvailableProducts)
	prod := router.PathPrefix("/product").Subrouter()
	authProd := prod.NewRoute().Subrouter()
	authProd.Use(middleware.TokenMiddleware)
	authProd.Methods(http.MethodPut).Path("/{id}").HandlerFunc(appDelivery.ProductDelivery.UpdateProduct)
	authProd.Methods(http.MethodDelete).Path("/{id}").HandlerFunc(appDelivery.ProductDelivery.DeleteProduct)
	authProd.Methods(http.MethodPost).HandlerFunc(appDelivery.ProductDelivery.InsertProduct)
	prod.Methods(http.MethodGet).Path("/{id}").HandlerFunc(appDelivery.ProductDelivery.GetProduct)

	return router
}
