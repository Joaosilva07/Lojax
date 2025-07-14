package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(filepath string) error {
	var err error
	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		return err
	}
	return createTables()
}

func createTables() error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS categorias (
            ID_Categoria INTEGER PRIMARY KEY AUTOINCREMENT,
            Nome_Categoria VARCHAR(255)
        );`,
		`CREATE TABLE IF NOT EXISTS produtos (
            ID_Produto INTEGER PRIMARY KEY AUTOINCREMENT,
            Nome_Produto VARCHAR(255),
            Descricao TEXT,
            Preco DECIMAL(10,2),
            Classificacao VARCHAR(50)
        );`,
		`CREATE TABLE IF NOT EXISTS produtos_categorias (
            ID_Produto INTEGER,
            ID_Categoria INTEGER,
            PRIMARY KEY (ID_Produto, ID_Categoria),
            FOREIGN KEY (ID_Produto) REFERENCES produtos(ID_Produto),
            FOREIGN KEY (ID_Categoria) REFERENCES categorias(ID_Categoria)
        );`,
		`CREATE TABLE IF NOT EXISTS estoque (
            ID_Estoque INTEGER PRIMARY KEY AUTOINCREMENT,
            ID_Produto INTEGER,
            Quantidade_em_Estoque INT,
            FOREIGN KEY (ID_Produto) REFERENCES produtos(ID_Produto)
        );`,
		`CREATE TABLE IF NOT EXISTS avaliacoes_produtos (
            ID_Avaliacao INTEGER PRIMARY KEY AUTOINCREMENT,
            ID_Produto INTEGER,
            Comentario TEXT,
            Classificacao INT,
            Data_Avaliacao DATE,
            FOREIGN KEY (ID_Produto) REFERENCES produtos(ID_Produto)
        );`,
		`CREATE TABLE IF NOT EXISTS clientes (
            ID_Cliente INTEGER PRIMARY KEY AUTOINCREMENT,
            Nome_Cliente VARCHAR(255),
            Endereco VARCHAR(255),
            Email VARCHAR(255),
            Num_Telefone VARCHAR(15)
        );`,
		`CREATE TABLE IF NOT EXISTS estados (
            ID_Estado INTEGER PRIMARY KEY AUTOINCREMENT,
            Nome_Estado VARCHAR(50),
            Sigla_Estado VARCHAR(2)
        );`,
		`CREATE TABLE IF NOT EXISTS cidades (
            ID_Cidade INTEGER PRIMARY KEY AUTOINCREMENT,
            Nome_Cidade VARCHAR(50),
            ID_Estado INTEGER,
            FOREIGN KEY (ID_Estado) REFERENCES estados(ID_Estado)
        );`,
		`CREATE TABLE IF NOT EXISTS enderecos_entrega (
            ID_Endereco_Entrega INTEGER PRIMARY KEY AUTOINCREMENT,
            ID_Cliente INTEGER,
            Rua VARCHAR(255),
            ID_Cidade INTEGER,
            CEP VARCHAR(10),
            FOREIGN KEY (ID_Cliente) REFERENCES clientes(ID_Cliente),
            FOREIGN KEY (ID_Cidade) REFERENCES cidades(ID_Cidade)
        );`,
		`CREATE TABLE IF NOT EXISTS funcionarios (
            ID_Funcionario INTEGER PRIMARY KEY AUTOINCREMENT,
            Nome_Funcionario VARCHAR(255),
            Cargo VARCHAR(255)
        );`,
		`CREATE TABLE IF NOT EXISTS pedidos (
            ID_Pedido INTEGER PRIMARY KEY AUTOINCREMENT,
            Data_Pedido DATE,
            ID_Funcionario INTEGER,
            ID_Cliente INTEGER,
            Status_Pedido VARCHAR(20),
            FOREIGN KEY (ID_Funcionario) REFERENCES funcionarios(ID_Funcionario),
            FOREIGN KEY (ID_Cliente) REFERENCES clientes(ID_Cliente)
        );`,
		`CREATE TABLE IF NOT EXISTS itens_pedido (
            ID_Item INTEGER PRIMARY KEY AUTOINCREMENT,
            ID_Pedido INTEGER,
            ID_Produto INTEGER,
            Quantidade INT,
            FOREIGN KEY (ID_Pedido) REFERENCES pedidos(ID_Pedido),
            FOREIGN KEY (ID_Produto) REFERENCES produtos(ID_Produto)
        );`,
		`CREATE TABLE IF NOT EXISTS registros_pagamento (
            ID_Pagamento INTEGER PRIMARY KEY AUTOINCREMENT,
            ID_Pedido INTEGER,
            Metodo_Pagamento VARCHAR(45),
            Preco_Pagamento DECIMAL(10,2),
            Data_Pagamento DATE,
            FOREIGN KEY (ID_Pedido) REFERENCES pedidos(ID_Pedido)
        );`,
		`CREATE TABLE IF NOT EXISTS vendas (
            ID_Venda INTEGER PRIMARY KEY AUTOINCREMENT,
            Data_Venda DATE,
            ID_Produto INTEGER,
            Quantidade_Vendida INT,
            FOREIGN KEY (ID_Produto) REFERENCES produtos(ID_Produto)
        );`,
		`CREATE TABLE IF NOT EXISTS outras_lojas (
            ID_Loja INTEGER PRIMARY KEY AUTOINCREMENT,
            Nome_Loja VARCHAR(255),
            Categoria VARCHAR(255),
            Nome_Produto VARCHAR(255),
            Preco DECIMAL(10,2),
            Data_Inclusao DATE
        );`,
	}
	for _, stmt := range stmts {
		_, err := DB.Exec(stmt)
		if err != nil {
			log.Println("Erro ao criar tabela:", err)
			return err
		}
	}
	return nil
}
