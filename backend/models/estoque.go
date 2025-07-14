package models

type Estoque struct {
	ID_Estoque            int `json:"id_estoque"`
	ID_Produto            int `json:"id_produto"`
	Quantidade_em_Estoque int `json:"quantidade_em_estoque"`
}
