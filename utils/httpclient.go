package utils

import (
	"crypto/tls"
	"net/http"
	"time"
)

var (
	tr = &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	HttpClient = &http.Client{Transport: tr}
)
