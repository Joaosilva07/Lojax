package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	db "lojaX/backend"
	"lojaX/backend/bindings"
	repository "lojaX/backend/repositories"
	service "lojaX/backend/services"
)

//go:embed all:frontend/dist
var assets embed.FS

var (
	produtoBindings   *bindings.ProdutoBindings
	categoriaBindings *bindings.CategoriaBindings
	estoqueBindings   *bindings.EstoqueBindings
)

func InitBindings() {
	repoProduto := repository.NewProdutoRepository(db.DB)
	serviceProduto := service.NewProdutoService(*repoProduto)
	produtoBindings = bindings.NewProdutoBindings(serviceProduto)

	repoCategoria := repository.NewCategoriaRepository(db.DB)
	serviceCategoria := service.NewCategoriaService(*repoCategoria)
	categoriaBindings = bindings.NewCategoriaBindings(serviceCategoria)

	repoEstoque := repository.NewEstoqueRepository(db.DB)
	serviceEstoque := service.NewEstoqueService(*repoEstoque)
	estoqueBindings = bindings.NewEstoqueBindings(serviceEstoque)
}
func main() {
	InitializeDatabase()
	InitBindings()

	app := NewApp()

	err := wails.Run(&options.App{
		Title:  "lojaX",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			produtoBindings,
			categoriaBindings,
			app.estoqueBindings,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

}
