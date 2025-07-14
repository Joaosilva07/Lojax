package bindings

import (
	"lojaX/backend/models"
	service "lojaX/backend/services"
)

type ProdutoBindings struct {
	service *service.ProdutoService
}

func NewProdutoBindings(service *service.ProdutoService) *ProdutoBindings {
	return &ProdutoBindings{service: service}
}

func (b *ProdutoBindings) CreateProduto(produto models.Produto) (int64, error) {
	return b.service.Create(&produto)
}

func (b *ProdutoBindings) GetProduto(id int64) (*models.Produto, error) {
	return b.service.GetByID(id)
}

func (b *ProdutoBindings) UpdateProduto(produto models.Produto) error {
	return b.service.Update(&produto)
}

func (b *ProdutoBindings) DeleteProduto(id int64) error {
	return b.service.Delete(id)
}

func (b *ProdutoBindings) ListAllProdutos() ([]models.Produto, error) {
	produtos, err := b.service.ListAll()
	if err != nil {
		return nil, err
	}
	return produtos, nil
}
