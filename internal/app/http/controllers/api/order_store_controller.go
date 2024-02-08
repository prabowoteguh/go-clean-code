package api

import (
	"go-assignment-bootcamp/internal/app/http/requests"
	"go-assignment-bootcamp/internal/app/presenters"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderStoreController struct {
	useCase   OrderStoreUseCaseInterface
	presenter OrderStorePresenterInterface
}

func NewOrderStoreController(
	useCase OrderStoreUseCaseInterface,
	presenter OrderStorePresenterInterface,
) *OrderStoreController {
	return &OrderStoreController{
		useCase:   useCase,
		presenter: presenter,
	}
}

func (controller *OrderStoreController) Store(context *gin.Context) {
	var request requests.OrderStoreRequest

	err := request.Validate(context)
	if err != nil {
		presenters.PresentErrorPresenter(context, http.StatusBadRequest, err)
		return
	}

	response, useCaseError := controller.useCase.Execute(&request)
	if useCaseError != nil {
		presenters.PresentErrorPresenter(context, http.StatusBadRequest, useCaseError)
		return
	}

	controller.presenter.Present(context, response)
}
