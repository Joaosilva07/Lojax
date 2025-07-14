package repository

import (
	"database/sql"
	"lojaX/backend/models"
)

type ProdutoRepository struct {
	DB *sql.DB
}

func NewProdutoRepository(db *sql.DB) *ProdutoRepository {
	return &ProdutoRepository{DB: db}
}

func (r *ProdutoRepository) Create(produto *models.Produto) (int64, error) {
	stmt, err := r.DB.Prepare(`INSERT INTO produtos (Nome_Produto, Descricao, Preco, Classificacao) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(produto.NomeProduto, produto.Descricao, produto.Preco, produto.Classificacao)
	if err != nil {
		return 0, err
	}
	produtoID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	for _, catID := range produto.Categorias {
		_ = r.AddCategoriaToProduto(produtoID, catID)
	}

	return produtoID, nil
}

func (r *ProdutoRepository) GetByID(id int64) (*models.Produto, error) {
	row := r.DB.QueryRow(`SELECT ID_Produto, Nome_Produto, Descricao, Preco, Classificacao FROM produtos WHERE ID_Produto = ?`, id)

	produto := &models.Produto{}
	err := row.Scan(&produto.IDProduto, &produto.NomeProduto, &produto.Descricao, &produto.Preco, &produto.Classificacao)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	catRows, err := r.DB.Query("SELECT ID_Categoria FROM produtos_categorias WHERE ID_Produto = ?", produto.IDProduto)
	if err == nil {
		var categorias []int64
		for catRows.Next() {
			var catID int64
			if err := catRows.Scan(&catID); err == nil {
				categorias = append(categorias, catID)
			}
		}
		catRows.Close()
		produto.Categorias = categorias
	}

	return produto, nil
}

func (r *ProdutoRepository) Update(produto *models.Produto) error {
	stmt, err := r.DB.Prepare(`UPDATE produtos SET Nome_Produto = ?, Descricao = ?, Preco = ?, Classificacao = ? WHERE ID_Produto = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(produto.NomeProduto, produto.Descricao, produto.Preco, produto.Classificacao, produto.IDProduto)
	if err != nil {
		return err
	}

	_, _ = r.DB.Exec("DELETE FROM produtos_categorias WHERE ID_Produto = ?", produto.IDProduto)
	for _, catID := range produto.Categorias {
		_ = r.AddCategoriaToProduto(int64(produto.IDProduto), catID)
	}

	return nil
}

func (r *ProdutoRepository) Delete(id int64) error {
	stmt, err := r.DB.Prepare(`DELETE FROM produtos WHERE ID_Produto = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

func (r *ProdutoRepository) ListAll() ([]models.Produto, error) {
	rows, err := r.DB.Query(`SELECT ID_Produto, Nome_Produto, Descricao, Preco, Classificacao FROM produtos`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var produtos []models.Produto
	for rows.Next() {
		var produto models.Produto
		err := rows.Scan(&produto.IDProduto, &produto.NomeProduto, &produto.Descricao, &produto.Preco, &produto.Classificacao)
		if err != nil {
			return nil, err
		}

		// Buscar categorias associadas
		catRows, err := r.DB.Query("SELECT ID_Categoria FROM produtos_categorias WHERE ID_Produto = ?", produto.IDProduto)
		if err == nil {
			var categorias []int64
			for catRows.Next() {
				var catID int64
				if err := catRows.Scan(&catID); err == nil {
					categorias = append(categorias, catID)
				}
			}
			catRows.Close()
			produto.Categorias = categorias
		}

		produtos = append(produtos, produto)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if produtos == nil {
		produtos = []models.Produto{}
	}

	return produtos, nil
}

func (r ProdutoRepository) AddCategoriaToProduto(produtoID, categoriaID int64) error {
	_, err := r.DB.Exec(
		"INSERT OR IGNORE INTO produtos_categorias (ID_Produto, ID_Categoria) VALUES (?, ?)",
		produtoID, categoriaID,
	)
	return err
}

func (r ProdutoRepository) RemoveCategoriaProduto(produtoID, categoriaID int64) error {
	_, err := r.DB.Exec(
		"DELETE FROM produtos_categorias WHERE ID_Produto = ? AND ID_Categoria = ?",
		produtoID, categoriaID,
	)
	return err
}

func (r ProdutoRepository) ListCategoriasByProduto(produtoID int64) ([]int64, error) {
	rows, err := r.DB.Query(
		"SELECT ID_Categoria FROM produtos_categorias WHERE ID_Produto = ?",
		produtoID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}
