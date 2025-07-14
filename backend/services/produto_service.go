package service

import (
	"lojaX/backend/models"
	repository "lojaX/backend/repositories"
)

type ProdutoService struct {
	repo repository.ProdutoRepository
}

func NewProdutoService(repo repository.ProdutoRepository) *ProdutoService {
	return &ProdutoService{repo: repo}
}

func (s *ProdutoService) Create(produto *models.Produto) (int64, error) {

	return s.repo.Create(produto)
}

func (s *ProdutoService) GetByID(id int64) (*models.Produto, error) {
	return s.repo.GetByID(id)
}

func (s *ProdutoService) Update(produto *models.Produto) error {
	return s.repo.Update(produto)
}

func (s *ProdutoService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *ProdutoService) ListAll() ([]models.Produto, error) {
	return s.repo.ListAll()
}

func (s *ProdutoService) AddCategoria(produtoID int64, categoriaID int64) error {
	return s.repo.AddCategoriaToProduto(produtoID, categoriaID)
}

func (s *ProdutoService) RemoveCategoria(produtoID int64, categoriaID int64) error {
	return s.repo.RemoveCategoriaProduto(produtoID, categoriaID)
}

func (s *ProdutoService) ListCategorias(produtoID int64) ([]models.Categoria, error) {
	categoriaIDs, err := s.repo.ListCategoriasByProduto(produtoID)
	if err != nil {
		return nil, err
	}
	categorias := make([]models.Categoria, len(categoriaIDs))
	for i, id := range categoriaIDs {
		categorias[i] = models.Categoria{ID: int(id)}
	}
	return categorias, nil
}
