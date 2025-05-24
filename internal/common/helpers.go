package common

import (
	"net/http"
)

func (deps *Dependencies) ServerError(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		deps.Logger.Error("Server error", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
