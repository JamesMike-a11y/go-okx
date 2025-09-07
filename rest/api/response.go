package api

import "encoding/json"

type IResponse interface {
	GetCode() string
	GetMessage() string
	IsOk() bool
}

type Response struct {
	Code    string          `json:"code"`
	Message string          `json:"msg"`
	Data    json.RawMessage `json:"data"`
}

func (r Response) GetCode() string {
	return r.Code
}

func (r Response) GetMessage() string {
	data, _ := json.Marshal(r)
	return string(data)
}

func (r Response) IsOk() bool {
	return r.Code == "0"
}
