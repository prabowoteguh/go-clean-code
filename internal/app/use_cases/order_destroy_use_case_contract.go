package use_cases

import "go-assignment-bootcamp/internal/app/models"

type OrderShowDestroyUseCaseInterface interface {
	Execute(id int) (*models.Order, error)
}

type OrderDestroyRepositoryInterface interface {
	Make(*models.Order) error
}
