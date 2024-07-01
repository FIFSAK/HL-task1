package handlers

import (
	"HL-task1/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var (
	idCounter int
	db        sync.Map
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
	fmt.Printf("Request created successfully\n%s\n", req.String())
	response, err := Ping(req)
	if err != nil {
		http.Error(w, "ping error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Response received successfully\n%s\nj", response.String())

	db.Store(response.ID, response)

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func generateID() int {
	idCounter++
	return idCounter
}

func Ping(r models.Request) (models.Response, error) {
	serverResp, err := http.Get(r.Url)
	if err != nil {
		log.Fatal(err)
		return models.Response{}, err
	}
	defer serverResp.Body.Close()

	length := serverResp.ContentLength
	if length == -1 {
		body, err := io.ReadAll(serverResp.Body)
		if err != nil {
			log.Fatal(err)
			return models.Response{}, err
		}
		length = int64(len(body))
	}

	response := models.NewResponse(generateID(), serverResp.StatusCode, serverResp.Header, int(length))
	fmt.Println(response.String())
	return *response, nil
}

func GetStoredResponseHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}

	value, ok := db.Load(id)
	if !ok {
		http.Error(w, "response not found", http.StatusNotFound)
		return
	}

	response, ok := value.(models.Response)
	if !ok {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "error encoding response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
