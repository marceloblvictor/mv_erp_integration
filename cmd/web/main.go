package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"

	"github.com/marceloblvictor/mv_erp_integration/internal/common"
	"github.com/marceloblvictor/mv_erp_integration/internal/controller"
	"github.com/marceloblvictor/mv_erp_integration/internal/service"
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

	azCreds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		logger.Error("Failed to create Azure credentials", "error", err)
		os.Exit(1)
	}

	clientOptions := &azcosmos.ClientOptions{
		EnableContentResponseOnWrite: true,
	}

	epCosmos := os.Getenv("COSMOS_ENDPOINT")
	if epCosmos == "" {
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			fmt.Println(pair[0])
		}

		logger.Error("You must provide CosmosDB endpoint")
		os.Exit(1)
	}

	cosmosClient, err := azcosmos.NewClient(epCosmos, azCreds, clientOptions)
	if err != nil {
		logger.Error("Failed to connect to CosmosDB", "error", err)
		os.Exit(1)
	}

	dbClient, err := cosmosClient.NewDatabase("integration-db")
	if err != nil {
		logger.Error("Failed to connect to CosmosDB database", "error", err)
		os.Exit(1)
	}

	ordersDbContainer, err := dbClient.NewContainer("orders")
	if err != nil {
		logger.Error("Failed to connect to CosmosDB container", "error", err)
		os.Exit(1)
	}

	orderSvc := &service.OrderService{Container: ordersDbContainer}
	orderCtrl := &controller.OrderController{Service: orderSvc, Dependencies: deps}
	mux := http.NewServeMux()

	RegisterRoutes(mux, orderCtrl)

	logger.Info("Starting server on host " + cfg.host + " port " + cfg.port)

	err = http.ListenAndServe(":"+cfg.port, mux)

	logger.Error(err.Error())
	os.Exit(1)
}
