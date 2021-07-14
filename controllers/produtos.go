package controllers

import (
	"goStoreStudy/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	allProducts := models.GetAllProducts()
	templates.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CreateNewProduct(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	IdProduct := r.URL.Query().Get("id")
	models.DeleteProduct(IdProduct)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.GetProductById(idProduct)
	println(product.Name)
	templates.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		decription := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		idParsed, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID para int: ", err)
		}

		priceParsed, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do Preço para float: ", err)
		}

		quantityParsed, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Erro na conversão da Quantidade para int: ", err)
		}

		models.UpdateProduct(idParsed, quantityParsed, name, decription, priceParsed)
	}

	http.Redirect(w, r, "/", 301)
}
