export namespace models {
	
	export class Categoria {
	    id: number;
	    nome_categoria: string;
	
	    static createFrom(source: any = {}) {
	        return new Categoria(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.nome_categoria = source["nome_categoria"];
	    }
	}
	export class Estoque {
	    id_estoque: number;
	    id_produto: number;
	    quantidade_em_estoque: number;
	
	    static createFrom(source: any = {}) {
	        return new Estoque(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id_estoque = source["id_estoque"];
	        this.id_produto = source["id_produto"];
	        this.quantidade_em_estoque = source["quantidade_em_estoque"];
	    }
	}
	export class Produto {
	    id_produto: number;
	    nome_produto: string;
	    descricao: string;
	    preco: number;
	    classificacao: string;
	    categorias: number[];
	
	    static createFrom(source: any = {}) {
	        return new Produto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id_produto = source["id_produto"];
	        this.nome_produto = source["nome_produto"];
	        this.descricao = source["descricao"];
	        this.preco = source["preco"];
	        this.classificacao = source["classificacao"];
	        this.categorias = source["categorias"];
	    }
	}

}

