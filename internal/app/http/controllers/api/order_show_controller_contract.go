package api

import (
	"go-assignment-bootcamp/internal/app/models"
	"github.com/gin-gonic/gin"
)

type OrderShowUseCaseInterface interface {
	Execute(id int) (*models.Order, error)
}

type OrderShowPresenterInterface interface {
	Present(*gin.Context, *models.Order)
}
