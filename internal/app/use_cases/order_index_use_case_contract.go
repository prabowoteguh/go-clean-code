package use_cases

import (
	"go-assignment-bootcamp/internal/app/models"
	"github.com/gin-gonic/gin"
)

type OrderIndexRepositoryInterface interface {
	Make(interface{ PaginationDTOInterface }) []*models.Order
}

type PaginationDTOInterface interface {
	GetSortColumn() string
	GetSortType() string
	GetLimitValue() int
	GetOffsetValue() int
}

type OrderIndexRequestInterface interface {
	Validate(*gin.Context) error
	GetStart() int
	GetEnd() int
	GetSortColumn() string
	GetSortType() string
}
