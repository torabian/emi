package emigo

import "net/http"

type APIClient struct {
	BaseURL    string
	Headers    http.Header
	HTTPClient *http.Client
	Httpr      *http.Request
}
