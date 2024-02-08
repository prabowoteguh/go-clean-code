package providers

import "go-assignment-bootcamp/internal/app/helpers"

type ContainerProvider struct {
	dependenciesContainer *helpers.DependenciesContainer
}

func NewContainerProvider() *ContainerProvider {
	return &ContainerProvider{
		helpers.NewDependenciesContainer(),
	}
}

func (containerProvider *ContainerProvider) GetContainer() *helpers.DependenciesContainer {
	return containerProvider.dependenciesContainer
}
