package contracts

type HttpRequest struct {
	Body interface{} `json:"body,omitempty"`
}

type HttpResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data,omitempty"`
}
