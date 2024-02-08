package api

import (
	"go-assignment-bootcamp/internal/app/models"
	"github.com/gin-gonic/gin"
)

type OrderIndexUseCaseInterface interface {
	Execute(interface{ OrderIndexRequestInterface }) []*models.Order
}

type OrderIndexRequestInterface interface {
	Validate(*gin.Context) error
	GetStart() int
	GetEnd() int
	GetSortColumn() string
	GetSortType() string
}

type OrderIndexPresenterInterface interface {
	Present(*gin.Context, []*models.Order)
}
