package model

type OrderAction int

const (
	Create OrderAction = iota + 1
	Update
	Delete
)

var actionName = map[OrderAction]string{
	Create: "create",
	Update: "update",
	Delete: "delete",
}

func (oa OrderAction) String() string {
	return actionName[oa]
}

type Order struct {
	ID           string      `json:"id"`
	Description  string      `json:"description"`
	CustomerName string      `json:"customer_name"`
	Total        float64     `json:"total"`
	ReceivedAt   string      `json:"received_at"`
	Action       OrderAction `json:"action"`
}
