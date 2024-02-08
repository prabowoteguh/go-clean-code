package bindings

import (
	"go-assignment-bootcamp/internal/app/helpers"
	"go-assignment-bootcamp/internal/app/presenters"
)

type PresenterProvider struct {
	OrderIndexPresenter *presenters.OrderIndexPresenter
	OrderPresenter      *presenters.OrderPresenter
}

func NewPresenterProvider(container *helpers.DependenciesContainer) *PresenterProvider {
	return &PresenterProvider{
		OrderIndexPresenter: presenters.NewOrderListPresenter(),
		OrderPresenter:      presenters.NewOrderPresenter(),
	}
}
