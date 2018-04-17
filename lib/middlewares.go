package lib

import (
	"net/http"
)

type AppHandler struct {
	AppContext
	Handler func(AppContext, http.ResponseWriter, *http.Request) (int, error)
}

func (h AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := h.Handler(h.AppContext, w, r)
	if err != nil {
		switch status {
		case http.StatusBadRequest:
			w.WriteHeader(status)
		case http.StatusNotFound:
			w.WriteHeader(status)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}