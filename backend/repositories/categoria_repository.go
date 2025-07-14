package repository

import (
	"database/sql"
	"log"
	"lojaX/backend/models"
)

type CategoriaRepository struct {
	db *sql.DB
}

func NewCategoriaRepository(db *sql.DB) *CategoriaRepository {
	return &CategoriaRepository{db: db}
}

func (r *CategoriaRepository) Create(categoria *models.Categoria) (int64, error) {
	stmt, err := r.db.Prepare("INSERT INTO categorias (Nome_Categoria) VALUES (?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(categoria.NomeCategoria)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *CategoriaRepository) Listar() ([]models.Categoria, error) {
	rows, err := r.db.Query("SELECT ID_Categoria, Nome_Categoria FROM categorias")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorias []models.Categoria
	for rows.Next() {
		var c models.Categoria
		if err := rows.Scan(&c.ID, &c.NomeCategoria); err != nil {
			log.Println("Erro ao ler categoria:", err)
			continue
		}
		categorias = append(categorias, c)
	}
	return categorias, nil
}

func (r *CategoriaRepository) Update(categoria *models.Categoria) error {
	stmt, err := r.db.Prepare("UPDATE categorias SET Nome_Categoria = ? WHERE ID_Categoria = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(categoria.NomeCategoria, categoria.ID)
	return err
}

func (r *CategoriaRepository) Delete(id int64) error {
	stmt, err := r.db.Prepare("DELETE FROM categorias WHERE ID_Categoria = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
