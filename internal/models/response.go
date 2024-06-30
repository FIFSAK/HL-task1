package models

import "sync"

type Response struct {
	sync.Mutex
	id      int
	status  string
	headers map[string]string
	length  int
}

func NewResponse(id int, status string, headers map[string]string, length int) *Response {
	return &Response{
		id:      id,
		status:  status,
		headers: headers,
		length:  length,
	}
}

func (resp *Response) GetHeaders() map[string]string {
	resp.Lock()
	defer resp.Unlock()
	return resp.headers
}

func (resp *Response) GetHeadersValue(key string) (string, bool) {
	resp.Lock()
	defer resp.Unlock()
	val, ok := resp.headers[key]
	return val, ok
}
