package http

import (
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Minute * 2,
	}
}
