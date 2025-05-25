package model

type OrderAction int

const (
	ActionCreate OrderAction = iota + 1
	ActionUpdate
	ActionDelete
)

var actionName = map[OrderAction]string{
	ActionCreate: "create",
	ActionUpdate: "update",
	ActionDelete: "delete",
}

func (oa OrderAction) String() string {
	return actionName[oa]
}

type Order struct {
	ID           string      `json:"id"`
	Description  string      `json:"description"`
	CustomerName string      `json:"CustomerName"`
	Total        float64     `json:"total"`
	ReceivedAt   string      `json:"received_at"`
	Action       OrderAction `json:"action"`
}
