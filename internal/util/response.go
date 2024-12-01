package util

import (
	"net/http"
	"time"
)

type NormalResponse struct {
	HttpCode  int         `json:"httpCode"`
	Result    interface{} `json:"result"`
	Timestamp time.Time   `json:"timestamp"`
}

type PaginationResponse struct {
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	TotalItems int         `json:"totalItems"`
	TotalPages int         `json:"totalPages"`
	Result     interface{} `json:"result"`
}

func OK(data interface{}) *NormalResponse {
	return &NormalResponse{
		HttpCode:  http.StatusOK,
		Result:    data,
		Timestamp: time.Now(),
	}
}

func Created(data interface{}) *NormalResponse {
	return &NormalResponse{
		HttpCode:  http.StatusCreated,
		Result:    data,
		Timestamp: time.Now(),
	}
}

func Pagination(page int, limit int, totalPage int, totalItems int, data interface{}) *NormalResponse {
	return &NormalResponse{
		HttpCode: 200,
		Result: &PaginationResponse{
			Page:       page,
			Limit:      limit,
			TotalItems: totalItems,
			TotalPages: totalPage,
			Result:     data,
		},
		Timestamp: time.Now(),
	}
}
