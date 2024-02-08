package api

import (
	"go-assignment-bootcamp/internal/app/presenters"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderShowController struct {
	useCase   OrderShowUseCaseInterface
	presenter OrderShowPresenterInterface
}

func NewOrderShowController(
	useCase OrderShowUseCaseInterface,
	presenter OrderShowPresenterInterface,
) *OrderShowController {
	return &OrderShowController{
		useCase:   useCase,
		presenter: presenter,
	}
}

func (controller *OrderShowController) Show(context *gin.Context) {
	paramId := context.Param("id")

	id, err := strconv.Atoi(paramId)
	if err != nil {
		message := "The ID parameter is invalid."
		presenters.PresentErrorPresenter(context, http.StatusBadRequest, message)
		return
	}

	response, useCaseError := controller.useCase.Execute(id)
	if useCaseError != nil {
		presenters.PresentErrorPresenter(context, http.StatusBadRequest, useCaseError)
		return
	}

	controller.presenter.Present(context, response)
}
