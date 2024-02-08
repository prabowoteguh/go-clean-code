package use_cases

import (
	"go-assignment-bootcamp/internal/app/models"
	"github.com/gin-gonic/gin"
)

type OrderShowUseCaseInterface interface {
	Execute(id int) (*models.Order, error)
}

type OrderUpdateRepositoryInterface interface {
	Make(interface{ OrderUpdateRequestInterface }, *models.Order) error
}

type OrderUpdateRequestInterface interface {
	Validate(*gin.Context) error
	GetGuestCount() int
	GetTransportCount() int
	GetUserName() string
	GetEmail() string
	GetPhone() string
	GetNote() string
	GetAdminNote() string
}
