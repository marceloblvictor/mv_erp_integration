package main

import (
	"net/http"

	"dev.azure.com/MarceloVictor1/_git/mv_erp_integration/internal/controller"
)

func RegisterRoutes(mux *http.ServeMux, orderCtrl *controller.OrderController) {

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("GET /static/", http.StripPrefix(("/static"), fileServer))
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the mv_erp_integration!"))
	})

	mux.HandleFunc("GET /orders/{$}", orderCtrl.GetList)
	mux.HandleFunc("GET /orders/{id}", orderCtrl.GetById)
	mux.HandleFunc("POST /orders/create", orderCtrl.PostCreate)
	mux.HandleFunc("PUT /orders/update/{id}", orderCtrl.PutUpdate)
	mux.HandleFunc("DELETE /orders/delete/{id}", orderCtrl.Delete)
}
