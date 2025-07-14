package factories

import (
	"database/sql"
	repository "lojaX/backend/repositories"
	service "lojaX/backend/services"
)

func NewEstoqueService(db *sql.DB) *service.EstoqueService {
	repo := repository.NewEstoqueRepository(db)
	return service.NewEstoqueService(*repo)
}
