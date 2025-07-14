package factories

import (
	"database/sql"
	repository "lojaX/backend/repositories"
	service "lojaX/backend/services"
)

func NewCategoriaService(db *sql.DB) *service.CategoriaService {
	repo := repository.NewCategoriaRepository(db)
	return service.NewCategoriaService(*repo)
}
