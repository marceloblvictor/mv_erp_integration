package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/google/uuid"

	"github.com/marceloblvictor/mv_erp_integration/internal/model"
)

type OrderService struct {
	Container *azcosmos.ContainerClient
}

func (svc OrderService) GetList() ([]string, error) {

	// query := svc.Container.NewQueryItemsPager("SELECT * FROM c", azcosmos.NewPartitionKey(), &azcosmos.QueryOptions{})
	// query.NextPage()

	return []string{"Order1", "Order2", "Order3"}, nil
}

func (svc OrderService) GetById(id int) (string, error) {
	if id <= 0 {
		return "", nil
	}
	return fmt.Sprintf("Order with ID %d", id), nil
}

func (svc OrderService) Create() (string, error) {
	order := model.Order{
		ID:           uuid.New().String(),
		Description:  "test",
		CustomerName: "Yamba Surfboard",
		ReceivedAt:   time.Now().UTC().Format(time.DateTime),
		Total:        850.00,
		Action:       model.ActionCreate,
	}

	partitionKey := azcosmos.NewPartitionKeyString(order.CustomerName)

	context := context.TODO()

	bytes, err := json.Marshal(order)
	if err != nil {
		return "", err
	}

	response, err := svc.Container.UpsertItem(context, partitionKey, bytes, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Order created: %s", response.RawResponse.Status), nil
}
