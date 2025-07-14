package factories

import (
	"database/sql"
	repository "lojaX/backend/repositories"
	service "lojaX/backend/services"
)

func NewProdutoService(db *sql.DB) *service.ProdutoService {
	repo := repository.NewProdutoRepository(db)
	return service.NewProdutoService(*repo)
}
