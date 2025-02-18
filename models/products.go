package models

import "alura-store/db"

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Amount            int
}

func GetAllProducts() []Product {
	db := db.ConnectDB()

	selectAllProducts, err := db.Query("select * from products order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, amount int) {
	db := db.ConnectDB()

	insertData, err := db.Prepare("insert into products(name, description, price, amount) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, amount)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDB()

	deleteProduct, err := db.Prepare("delete from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectDB()

	productDB, err := db.Query("select * from products where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for productDB.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = productDB.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}
		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Amount = amount
	}
	defer db.Close()
	return productToUpdate
}

func UpdateProduct(id int, name, description string, price float64, amount int) {
	db := db.ConnectDB()

	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, amount=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, amount, id)
	defer db.Close()
}
