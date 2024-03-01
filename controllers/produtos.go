package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/emersontenorio1/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	todosOsProdutos := models.BuscaTodosOsProdutos()

	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do preco: ", err)
		}

		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao do quantidade: ", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConv, quantidadeConv)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	IdDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(IdDoProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	IdDoProduto := r.URL.Query().Get("id")
	produto := models.UpdateProduto(IdDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversao do ID para Int:", err)
		}
		precoConv, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do preço para float:", err)
		}
		quantidadeConv, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao do quantidade para Int:", err)
		}

		models.AtualizaProduto(idConv, nome, descricao, precoConv, quantidadeConv)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
