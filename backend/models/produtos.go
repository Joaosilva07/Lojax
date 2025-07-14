package models

type Produto struct {
	IDProduto     int     `json:"id_produto"`
	NomeProduto   string  `json:"nome_produto"`
	Descricao     string  `json:"descricao"`
	Preco         float64 `json:"preco"`
	Classificacao string  `json:"classificacao"`
	Categorias    []int64 `json:"categorias"`
}
