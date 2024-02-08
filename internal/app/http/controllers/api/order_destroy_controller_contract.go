package api

import (
	"go-assignment-bootcamp/internal/app/models"
	"github.com/gin-gonic/gin"
)

type OrderDestroyUseCaseInterface interface {
	Execute(id int) (*models.Order, error)
}

type OrderDestroyPresenterInterface interface {
	Present(*gin.Context, *models.Order)
}
