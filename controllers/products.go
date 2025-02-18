package controllers

import (
	"alura-store/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		convertedPriceToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price: ", err)
		}

		convertedAmountToInt, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error converting amount: ", err)
		}

		models.CreateNewProduct(name, description, convertedPriceToFloat, convertedAmountToInt)
	}
	http.Redirect(w, r, "/", 301)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		convertedPriceToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting price: ", err)
		}

		convertedAmountToInt, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Error converting amount: ", err)
		}

		convertedIdToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error converting id: ", err)
		}

		models.UpdateProduct(convertedIdToInt, name, description, convertedPriceToFloat, convertedAmountToInt)
	}
	http.Redirect(w, r, "/", 301)
}
