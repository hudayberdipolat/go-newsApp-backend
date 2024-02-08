package http

import (
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout: time.Minute * 1,
			MaxIdleConns:        30,
			MaxIdleConnsPerHost: 100,
			MaxConnsPerHost:     100,
			IdleConnTimeout:     30 * time.Second,
		},
		Timeout: time.Minute * 2,
	}
}
