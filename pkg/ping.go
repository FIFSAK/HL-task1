package pkg

import (
	"HL-task1/internal/models"
	"fmt"
	"io"
	"log"
	"net/http"
)

var idCounter int

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
