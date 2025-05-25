package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"dev.azure.com/MarceloVictor1/_git/mv_erp_integration/internal/common"
	"dev.azure.com/MarceloVictor1/_git/mv_erp_integration/internal/service"
)

type OrderController struct {
	Service *service.OrderService
	*common.Dependencies
}

func (ctrl *OrderController) GetList(w http.ResponseWriter, r *http.Request) {
	ctrl.Logger.Info("Fetching order list")

	res, _ := ctrl.Service.GetList()

	w.Write([]byte(strings.Join(res, "\n")))
}

func (ctrl *OrderController) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		ctrl.ClientError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprintf(w, `{"id":"%d"}`, id)
}

func (ctrl *OrderController) PostCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("created!"))
}

func (ctrl *OrderController) PutUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updated!"))
}

func (ctrl *OrderController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleted!"))
}
