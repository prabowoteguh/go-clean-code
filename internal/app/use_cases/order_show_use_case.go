package use_cases

import "go-assignment-bootcamp/internal/app/models"

type OrderShowUseCase struct {
	orderShowRepository OrderShowRepositoryInterface
}

func NewOrderShowUseCase(orderShowRepository OrderShowRepositoryInterface) *OrderShowUseCase {
	return &OrderShowUseCase{orderShowRepository}
}

func (useCase *OrderShowUseCase) Execute(id int) (*models.Order, error) {
	return useCase.orderShowRepository.Make(id)
}
