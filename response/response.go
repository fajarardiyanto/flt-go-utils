package response

import (
	"encoding/json"
	"net/http"
)

const (
	contentTypeHeader = "Content-Type"
	contentTypeJSON   = "application/json"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Errors  string      `json:"errors"`
}

// ResponsePagination is used to encode all response json's with pagination
type ResponsePagination struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data,omitempty"`
	Page      int         `json:"page"`
	Limit     int         `json:"limit"`
	TotalPage int         `json:"total_page"`
	Errors    string      `json:"errors"`
}

func WriteResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set(contentTypeHeader, contentTypeJSON)
	if data == nil {
		w.WriteHeader(statusCode)
		return
	}
	if _, err := json.Marshal(data); err != nil {
		data = Response{
			Success: false,
			Data:    []interface{}{},
			Errors:  err.Error(),
		}

		statusCode = 500
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(data)
		return
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Data:    data,
		Errors:  "",
	})
}

func WriteResponseWithPagination(w http.ResponseWriter, statusCode int, data interface{}, page, limit, totalPage int) {
	w.Header().Set(contentTypeHeader, contentTypeJSON)
	if data == nil {
		w.WriteHeader(statusCode)
		return
	}
	if _, err := json.Marshal(data); err != nil {
		data = Response{
			Success: false,
			Data:    []interface{}{},
			Errors:  err.Error(),
		}

		statusCode = 500
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(data)
		return
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ResponsePagination{
		Success:   true,
		Data:      data,
		Page:      page,
		Limit:     limit,
		TotalPage: totalPage,
		Errors:    "",
	})
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set(contentTypeHeader, contentTypeJSON)
	if msg == "" {
		w.WriteHeader(statusCode)
		return
	}
	data := Response{
		Success: false,
		Data:    []interface{}{},
		Errors:  msg,
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
	return
}
