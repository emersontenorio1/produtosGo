package models

import (
	"log"

	"github.com/emersontenorio1/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	insereDadosNoBanco, err := db.Prepare("insert into produtos(Nome, Descricao, Preco, Quantidade) values(?, ?, ?, ?)")
	if err != nil {
		log.Panic(err.Error())

	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()
	deletarOProDuto, err := db.Prepare("delete from produtos where id = ?")
	if err != nil {
		log.Panic(err.Error())

	}
	deletarOProDuto.Exec(id)
	defer db.Close()
}

func UpdateProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()
	produtoDOBanco, err := db.Query("select * from produtos where id = ?", id)
	if err != nil {
		log.Panic(err.Error())

	}
	produtoUpdate := Produto{}

	for produtoDOBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDOBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			log.Panic(err.Error())

		}
		produtoUpdate.Id = id
		produtoUpdate.Nome = nome
		produtoUpdate.Descricao = descricao
		produtoUpdate.Quantidade = quantidade
		produtoUpdate.Preco = preco

	}
	defer db.Close()
	return produtoUpdate
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	AtualizaProduto, err := db.Prepare("update produtos set Nome = ?, Descricao = ?, Preco = ?, Quantidade = ? where Id=?")
	if err != nil {
		log.Panic(err.Error())

	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
