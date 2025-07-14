package service

import (
	"lojaX/backend/models"
	repository "lojaX/backend/repositories"
)

type EstoqueService struct {
	repo repository.EstoqueRepository
}

func NewEstoqueService(repo repository.EstoqueRepository) *EstoqueService {
	return &EstoqueService{repo: repo}
}

func (s *EstoqueService) Create(estoque *models.Estoque) (int64, error) {
	return s.repo.Create(estoque)
}

func (s *EstoqueService) GetAll() ([]*models.Estoque, error) {
	return s.repo.FindAll()
}

func (s *EstoqueService) Update(estoque *models.Estoque) error {
	return s.repo.Update(estoque)
}
