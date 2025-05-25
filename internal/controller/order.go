package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/marceloblvictor/mv_erp_integration/internal/common"
	"github.com/marceloblvictor/mv_erp_integration/internal/service"
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

	ctrl.Logger.Info("Creating new order")

	res, err := ctrl.Service.Create()
	if err != nil {
		ctrl.ServerError(w, r, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message":"%s"}`, res)
}

func (ctrl *OrderController) PutUpdate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("updated!"))
}

func (ctrl *OrderController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("deleted!"))
}
