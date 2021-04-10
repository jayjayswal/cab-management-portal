package handlers

import "net/http"


func (h *Handler)  HealthCheck(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write([]byte(`{"status":"ok"}`))
}
