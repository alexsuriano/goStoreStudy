package models

import (
	"goStoreStudy/dataBase"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := dataBase.DbConnect()

	selectAllProducts, err := db.Query("select * from produtos order by id asc")

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
		p.Name = nome
		p.Description = descricao
		p.Price = preco
		p.Quantity = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func GetProductById(id string) Product {
	db := dataBase.DbConnect()

	selectedProduct, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	productEditing := Product{}

	for selectedProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := selectedProduct.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		productEditing.Id = id
		productEditing.Name = name
		productEditing.Description = description
		productEditing.Quantity = quantity
		productEditing.Price = price

		println(productEditing.Name)
	}
	defer db.Close()
	return productEditing
}

func CreateNewProduct(nome, descricao string, preco float64, quantidade int) {
	db := dataBase.DbConnect()

	createquery, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	createquery.Exec(nome, descricao, preco, quantidade)

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

func UpdateProduct(id int, quantity int, name string, description string, price float64) {
	db := dataBase.DbConnect()

	UpdateQuery, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	UpdateQuery.Exec(name, description, price, quantity, id)
	defer db.Close()
}
