package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"dev.azure.com/MarceloVictor1/_git/mv_erp_integration/internal/common"
	"dev.azure.com/MarceloVictor1/_git/mv_erp_integration/internal/controller"
	"dev.azure.com/MarceloVictor1/_git/mv_erp_integration/internal/service"
)

type cliConfig struct {
	port string
	host string
}

func main() {

	var cfg cliConfig

	flag.StringVar(&cfg.host, "h", "localhost", "host to run the server on")
	flag.StringVar(&cfg.port, "p", "8082", "port to run the server on")

	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	})
	logger := slog.New(loggerHandler)

	flag.Parse()

	logger.Info("Starting server...")
	logger.Info("Registering routes...")

	fileServer := http.FileServer(http.Dir("./static"))

	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.StripPrefix(("/static/"), fileServer))
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the mv_erp_integration!"))
	})

	deps := &common.Dependencies{
		Logger: logger,
	}

	orderSvc := &service.OrderService{}
	orderCtrl := &controller.OrderController{Service: orderSvc, Dependencies: deps}

	mux.HandleFunc("GET /orders/{$}", orderCtrl.GetList)
	mux.HandleFunc("GET /orders/{id}", orderCtrl.GetById)
	mux.HandleFunc("POST /orders/create", orderCtrl.PostCreate)
	mux.HandleFunc("PUT /orders/update/{id}", orderCtrl.PutUpdate)
	mux.HandleFunc("DELETE /orders/delete/{id}", orderCtrl.Delete)

	logger.Info("Starting server on host " + cfg.host + " port " + cfg.port)

	err := http.ListenAndServe(":"+cfg.port, mux)

	logger.Error(err.Error())
	os.Exit(1)
}
