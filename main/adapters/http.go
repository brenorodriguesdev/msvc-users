package adapters

import (
	"msvc-users/presentation/contracts"

	"github.com/gin-gonic/gin"
)

type Adapter struct {
	controller contracts.Controller
}

func NewAdapter(controller contracts.Controller) *Adapter {
	return &Adapter{controller: controller}
}

func (a *Adapter) Handle(c *gin.Context) {
	var httpRequest contracts.HttpRequest
	httpResponse, err := a.controller.Handle(httpRequest)
	if err != nil {
		c.JSON(httpResponse.StatusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpResponse.StatusCode, httpResponse.Data)
}
