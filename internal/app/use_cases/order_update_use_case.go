package use_cases

import "go-assignment-bootcamp/internal/app/models"

type OrderUpdateUseCase struct {
	showUseCase           OrderShowUseCaseInterface
	orderUpdateRepository OrderUpdateRepositoryInterface
}

func NewOrderUpdateUseCase(
	showUseCase OrderShowUseCaseInterface,
	orderUpdateRepository OrderUpdateRepositoryInterface,
) *OrderUpdateUseCase {
	return &OrderUpdateUseCase{
		showUseCase:           showUseCase,
		orderUpdateRepository: orderUpdateRepository,
	}
}

func (useCase *OrderUpdateUseCase) Execute(
	request interface{ OrderUpdateRequestInterface },
	id int,
) (*models.Order, error) {
	response, err := useCase.showUseCase.Execute(id)
	if err != nil {
		return response, err
	}

	repositoryError := useCase.orderUpdateRepository.Make(request, response)

	return response, repositoryError
}
