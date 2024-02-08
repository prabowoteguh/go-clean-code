package use_cases

import "go-assignment-bootcamp/internal/app/models"

type OrderShowRepositoryInterface interface {
	Make(id int) (*models.Order, error)
}
