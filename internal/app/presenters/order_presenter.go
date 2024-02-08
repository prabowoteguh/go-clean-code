package presenters

import (
	"go-assignment-bootcamp/internal/app/http/resources"
	"go-assignment-bootcamp/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderPresenter struct {
}

func NewOrderPresenter() *OrderPresenter {
	return &OrderPresenter{}
}

func (presenter *OrderPresenter) Present(context *gin.Context, order *models.Order) {
	orderResource := resources.NewOrderResource(order)
	response := gin.H{
		"data": orderResource,
	}
	context.JSON(http.StatusOK, response)
}
