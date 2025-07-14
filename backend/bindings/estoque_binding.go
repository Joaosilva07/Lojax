package bindings

import (
	"lojaX/backend/models"
	service "lojaX/backend/services"
)

type EstoqueBindings struct {
	service *service.EstoqueService
}

func NewEstoqueBindings(service *service.EstoqueService) *EstoqueBindings {
	return &EstoqueBindings{service: service}
}

func (b *EstoqueBindings) CreateEstoque(estoque models.Estoque) (int64, error) {
	return b.service.Create(&estoque)
}

func (b *EstoqueBindings) GetAllEstoques() ([]*models.Estoque, error) {
	return b.service.GetAll()
}

func (b *EstoqueBindings) UpdateEstoque(estoque models.Estoque) error {
	return b.service.Update(&estoque)
}
