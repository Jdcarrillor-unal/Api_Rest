package models

import (
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	writer  http.ResponseWriter
}

func CreateDeafaultMessage(w http.ResponseWriter) Response {
	return Response{Status: http.StatusOK, writer: w}
}
func MessageNotFound(w http.ResponseWriter) Response {
	return Response{Status: http.StatusNotFound, writer: w, Message: "paila"}
}
