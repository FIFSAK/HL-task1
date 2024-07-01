package models

import (
	"fmt"
)

type Request struct {
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func (req *Request) String() string {
	return "Request: {Method: " + req.Method + ", Url: " + req.Url + ", Headers: " + fmt.Sprintf("%v", req.Headers) + "}"
}

func NewRequest(method, url string, headers map[string]string) *Request {
	return &Request{
		Method:  method,
		Url:     url,
		Headers: headers,
	}
}

func (req *Request) GetHeaders() map[string]string {
	return req.Headers
}

func (req *Request) GetHeadersValue(key string) (string, bool) {
	val, ok := req.Headers[key]
	return val, ok
}

func (req *Request) SetHeadersValue(key, value string) {
	req.Headers[key] = value
}
