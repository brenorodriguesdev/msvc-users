package contracts

type Controller interface {
	Handle(httpRequest HttpRequest) (*HttpResponse, error)
}
