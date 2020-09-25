package handlers

import (
	"encoding/json"
	"net/http"
)

type ResponseHTTP struct {
	StatusCode int
	Response   ResponseData
}

type ResponseData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Response is the new type for define all of the response from service
type Response interface{}

var (
	ErrRespServiceMaintance = ResponseHTTP{
		StatusCode: http.StatusServiceUnavailable,
		Response:   ResponseData{Status: 601, Message: "service temporary unavailable"}}
	ErrRespUnauthorize = ResponseHTTP{
		StatusCode: http.StatusUnauthorized,
		Response:   ResponseData{Status: 400, Message: "authorization required"}}
	ErrRespAuthInvalid = ResponseHTTP{
		StatusCode: http.StatusUnauthorized,
		Response:   ResponseData{Status: 401, Message: "authorization invalid"}}
	ErrRespBadRequest = ResponseHTTP{
		StatusCode: http.StatusBadRequest,
		Response:   ResponseData{Status: 500, Message: "bad request"}}
	ErrRespInternalServer = ResponseHTTP{
		StatusCode: http.StatusServiceUnavailable,
		Response:   ResponseData{Status: 600, Message: "internal server error"}}
)

func writeResponse(res http.ResponseWriter, resp Response, code int, err error) {
	res.Header().Set("Content-Type", "application/json")

	if err != nil {
		errJSON := NewError("901", "404", err.Error())
		respErr, _ := json.Marshal(errJSON)
		res.WriteHeader(code)
		res.Write(respErr)
		return
	}

	r, _ := json.Marshal(resp)

	res.WriteHeader(code)
	res.Write(r)
	return
}

type Error struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	Title  string `json:"title"`
}

func NewError(id string, status string, title string) *Error {
	return &Error{
		Id:     id,
		Status: status,
		Title:  title,
	}
}
