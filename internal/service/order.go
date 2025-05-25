package service

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
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
