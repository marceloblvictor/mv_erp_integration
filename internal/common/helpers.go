package common

import (
	"net/http"
	"runtime/debug"
)

func (deps *Dependencies) ServerError(w http.ResponseWriter, r *http.Request, err error) {
	deps.Logger.Error("Internal Server error", "error", err, "method", r.Method, "url", r.URL.Path, "stack", string(debug.Stack()))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (deps *Dependencies) ClientError(w http.ResponseWriter, status int, msg string) {
	http.Error(w, msg, status)
}
