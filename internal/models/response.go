package models

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Response struct {
	ID      int                 `json:"id"`
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers"`
	Length  int                 `json:"length"`
}

func createKeyValuePairs(m map[string][]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

func (resp *Response) String() string {
	jsonData, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		return fmt.Sprintf("Response: {Id: %d, Status: %d, Headers: %s, Length: %d}", resp.ID, resp.Status, createKeyValuePairs(resp.Headers), resp.Length)
	}
	return string(jsonData)
}

func NewResponse(id int, status int, headers map[string][]string, length int) *Response {
	return &Response{
		ID:      id,
		Status:  status,
		Headers: headers,
		Length:  length,
	}
}

func (resp *Response) GetHeaders() map[string][]string {
	return resp.Headers
}

func (resp *Response) GetHeadersValue(key string) ([]string, bool) {
	val, ok := resp.Headers[key]
	return val, ok
}
