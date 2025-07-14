package main

import (
	"context"
	"fmt"
	db "lojaX/backend"
	"lojaX/backend/bindings"
	repository "lojaX/backend/repositories"
	svc "lojaX/backend/services"
)

// App struct
type App struct {
	ctx               context.Context
	produtoBindings   *bindings.ProdutoBindings
	categoriaBindings *bindings.CategoriaBindings
	estoqueBindings   *bindings.EstoqueBindings
}

// NewApp creates a new App application struct
func NewApp() *App {
	// Initialize database
	err := db.InitDB("loja.db")
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize database: %v", err))
	}

	// Initialize repositories
	produtoRepo := repository.NewProdutoRepository(db.DB)
	categoriaRepo := repository.NewCategoriaRepository(db.DB)
	estoqueRepo := repository.NewEstoqueRepository(db.DB)

	// Initialize services
	produtoService := svc.NewProdutoService(*produtoRepo)
	categoriaService := svc.NewCategoriaService(*categoriaRepo)
	estoqueService := svc.NewEstoqueService(*estoqueRepo)
	// Initialize bindings
	produtoBindings := bindings.NewProdutoBindings(produtoService)
	categoriaBindings := bindings.NewCategoriaBindings(categoriaService)
	estoqueBindings := bindings.NewEstoqueBindings(estoqueService)

	return &App{
		produtoBindings:   produtoBindings,
		categoriaBindings: categoriaBindings,
		estoqueBindings:   estoqueBindings,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
