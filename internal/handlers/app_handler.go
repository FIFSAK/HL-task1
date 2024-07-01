package handlers

import (
	"HL-task1/internal/models"
	"HL-task1/pkg"
	"encoding/json"
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func NewRequestHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "data reading error: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Request created successfully\n%s", req.String())
	resp, err := pkg.Ping(req)
	if err != nil {
		http.Error(w, "ping error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Response received successfully\n%s", resp.String())
}

//func GetResponse(writer http.ResponseWriter, request *http.Request) {
//
//}
