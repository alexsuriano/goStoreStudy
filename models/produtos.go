package models

import (
	"goStoreStudy/dataBase"
)

type Product struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetAllProducts() []Product {
	db := dataBase.DbConnect()

	selectAllProducts, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	produtos := []Product{}

	for selectAllProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectAllProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)

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

func CreateNewProduct(nome, descricao string, preco float64, quantidade int) {
	db := dataBase.DbConnect()

	query, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	query.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := dataBase.DbConnect()
	deleteProduct, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}
