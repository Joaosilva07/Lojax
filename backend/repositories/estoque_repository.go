package repository

import (
	"database/sql"
	"lojaX/backend/models"
)

type EstoqueRepository struct {
	db *sql.DB
}

func NewEstoqueRepository(db *sql.DB) *EstoqueRepository {
	return &EstoqueRepository{db: db}
}

func (r *EstoqueRepository) Create(estoque *models.Estoque) (int64, error) {
	var existingID int
	err := r.db.QueryRow("SELECT ID_Estoque FROM estoque WHERE ID_Produto = ?", estoque.ID_Produto).Scan(&existingID)
	if err == nil {

		_, err = r.db.Exec("UPDATE estoque SET Quantidade_em_Estoque = ? WHERE ID_Produto = ?", estoque.Quantidade_em_Estoque, estoque.ID_Produto)
		return int64(existingID), err
	} else if err != sql.ErrNoRows {
		return 0, err
	}

	stmt, err := r.db.Prepare("INSERT INTO estoque (ID_Produto, Quantidade_em_Estoque) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(estoque.ID_Produto, estoque.Quantidade_em_Estoque)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *EstoqueRepository) FindAll() ([]*models.Estoque, error) {
	rows, err := r.db.Query("SELECT ID_Estoque, ID_Produto, Quantidade_em_Estoque FROM estoque")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var estoques []*models.Estoque
	for rows.Next() {
		var estoque models.Estoque
		if err := rows.Scan(&estoque.ID_Estoque, &estoque.ID_Produto, &estoque.Quantidade_em_Estoque); err != nil {
			return nil, err
		}
		estoques = append(estoques, &estoque)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return estoques, nil
}

func (r *EstoqueRepository) Update(estoque *models.Estoque) error {
	stmt, err := r.db.Prepare("UPDATE estoque SET ID_Produto = ?, Quantidade_em_Estoque = ? WHERE ID_Estoque = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(estoque.ID_Produto, estoque.Quantidade_em_Estoque, estoque.ID_Estoque)
	return err
}
