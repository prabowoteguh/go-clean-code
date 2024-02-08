package use_cases

import "go-assignment-bootcamp/internal/app/models"

type OrderStoreUseCase struct {
	orderStoreRepository OrderStoreRepositoryInterface
}

func NewOrderStoreUseCase(orderStoreRepository OrderStoreRepositoryInterface) *OrderStoreUseCase {
	return &OrderStoreUseCase{orderStoreRepository}
}

func (useCase *OrderStoreUseCase) Execute(request interface{ OrderStoreRequestInterface }) (*models.Order, error) {
	return useCase.orderStoreRepository.Make(request)
}
