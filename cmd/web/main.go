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

	deps := &common.Dependencies{
		Logger: logger,
	}
	orderSvc := &service.OrderService{}
	orderCtrl := &controller.OrderController{Service: orderSvc, Dependencies: deps}
	mux := http.NewServeMux()

	RegisterRoutes(mux, orderCtrl)

	logger.Info("Starting server on host " + cfg.host + " port " + cfg.port)

	err := http.ListenAndServe(":"+cfg.port, mux)

	logger.Error(err.Error())
	os.Exit(1)
}
