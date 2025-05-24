package service

import "fmt"

type OrderService struct {
}

func (svc OrderService) GetList() ([]string, error) {
	return []string{"Order1", "Order2", "Order3"}, nil
}

func (svc OrderService) GetById(id int) (string, error) {
	if id <= 0 {
		return "", nil
	}
	return fmt.Sprintf("Order with ID %d", id), nil
}
