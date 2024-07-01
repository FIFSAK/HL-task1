package models

import (
	"fmt"
	"sync"
)

type Request struct {
	sync.Mutex
	Method  string            `json:"method"`
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func (req *Request) String() string {
	req.Lock()
	defer req.Unlock()
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
	req.Lock()
	defer req.Unlock()
	return req.Headers
}

func (req *Request) GetHeadersValue(key string) (string, bool) {
	req.Lock()
	defer req.Unlock()
	val, ok := req.Headers[key]
	return val, ok
}

func (req *Request) SetHeadersValue(key, value string) {
	req.Lock()
	defer req.Unlock()
	req.Headers[key] = value
}
