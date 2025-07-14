package bindings

import (
	"lojaX/backend/models"
	service "lojaX/backend/services"
)

type CategoriaBindings struct {
	service *service.CategoriaService
}

func NewCategoriaBindings(service *service.CategoriaService) *CategoriaBindings {
	return &CategoriaBindings{service: service}
}

func (b *CategoriaBindings) CreateCategoria(categoria models.Categoria) (int64, error) {
	return b.service.Create(&categoria)
}

func (b *CategoriaBindings) ListarCategorias() ([]models.Categoria, error) {
	categorias, err := b.service.Listar()
	if err != nil {
		return nil, err
	}
	return categorias, nil
}

func (b *CategoriaBindings) UpdateCategoria(categoria models.Categoria) error {
	return b.service.Update(&categoria)
}

func (b *CategoriaBindings) DeleteCategoria(id int64) error {
	return b.service.Delete(id)
}
