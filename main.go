package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func dbConnect() *sql.DB {
	connectionString := "user=postgres dbname=ale_loja password=postpass host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return db

}

type Produto struct {
	id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := dbConnect()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()

	selectAllProducts, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectAllProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAllProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	templates.ExecuteTemplate(w, "Index", produtos)

	defer db.Close()
}
