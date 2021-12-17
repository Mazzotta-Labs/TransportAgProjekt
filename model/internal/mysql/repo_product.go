package mysql

import (
	"TransportAgProjekt/model/entity"
)

func (r *MySqlRepository) FindAllProduct() []entity.Product {
	var Products []entity.Product
	result, err := db.Query("select * from Product")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var Product entity.Product

		err := result.Scan(&Product.CustomerId, &Product.DriverId, &Product.ProductDate, &Product.ProductId, &Product.ProductsId)
		if err != nil {
			panic(err.Error())
		}
		Products = append(Products, Product)
	}
	return Products
}

func (r *MySqlRepository) AddProduct(Product entity.Product) {
	customerId := Product.CustomerId
	driverId := Product.DriverId
	ProductDate := Product.ProductDate
	ProductId := Product.ProductId
	productId := Product.ProductsId

	stmt, err := db.Prepare("insert into model values (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(customerId, driverId, ProductDate, ProductId, productId)
	if err != nil {
		panic(err.Error())
	}
}

func (r *MySqlRepository) UpdateProduct(Product entity.Product) {
	// customerId := Product.CustomerId
	// driverId := Product.DriverId
	// ProductDate := Product.ProductDate
	// ProductId := Product.ProductId
	// productId := Product.ProductsId

}

func (r *MySqlRepository) DeleteProduct(Product entity.Product) {
	// ProductId, _ := strconv.Atoi(Product)
}
