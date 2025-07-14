package service

import (
	"lojaX/backend/models"

	repository "lojaX/backend/repositories"
)

type CategoriaService struct {
	repo repository.CategoriaRepository
}

func NewCategoriaService(repo repository.CategoriaRepository) *CategoriaService {
	return &CategoriaService{repo: repo}
}

func (s *CategoriaService) Create(categoria *models.Categoria) (int64, error) {
	return s.repo.Create(categoria)
}
func (s *CategoriaService) Listar() ([]models.Categoria, error) {
	return s.repo.Listar()
}

func (s *CategoriaService) Update(categoria *models.Categoria) error {
	return s.repo.Update(categoria)
}

func (s *CategoriaService) Delete(id int64) error {
	return s.repo.Delete(id)
}
