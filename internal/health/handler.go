package health

import "net/http"

type HealthHandler struct{}

func NewHealthHandler() HealthHandler {
	return HealthHandler{}
}

func (h HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
