package presenters

import (
	"go-assignment-bootcamp/internal/app/http/resources"
	"go-assignment-bootcamp/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderIndexPresenter struct {
}

func NewOrderListPresenter() *OrderIndexPresenter {
	return &OrderIndexPresenter{}
}

func (presenter *OrderIndexPresenter) Present(context *gin.Context, orders []*models.Order) {
	var response []*resources.OrderIndexResource

	for _, order := range orders {
		resource := resources.NewOrderIndexResource(order)
		response = append(response, resource)
	}

	context.JSON(
		http.StatusOK,
		gin.H{"data": response},
	)
}
