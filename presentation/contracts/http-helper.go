package contracts

func ServerError() *HttpResponse {
	return &HttpResponse{
		Data:       "Internal Server Error",
		StatusCode: 500,
	}
}

func BadRequest(error error) *HttpResponse {
	return &HttpResponse{
		Data:       error.Error(),
		StatusCode: 400,
	}
}

func Ok(data interface{}) *HttpResponse {
	return &HttpResponse{
		Data:       data,
		StatusCode: 200,
	}
}

func Created() *HttpResponse {
	return &HttpResponse{
		Data:       nil,
		StatusCode: 201,
	}
}
