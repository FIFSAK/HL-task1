package models

import "sync"

type Request struct {
	sync.Mutex
	method  string
	url     string
	headers map[string]string
}

func NewRequest(method, url string, headers map[string]string) *Request {
	return &Request{
		method:  method,
		url:     url,
		headers: headers,
	}
}

func (req *Request) GetHeaders() map[string]string {
	req.Lock()
	defer req.Unlock()
	return req.headers
}

func (req *Request) GetHeadersValue(key string) (string, bool) {
	req.Lock()
	defer req.Unlock()
	val, ok := req.headers[key]
	return val, ok
}
