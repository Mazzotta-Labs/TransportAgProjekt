package mysql

import (
	"TransportAgProjekt/model/entity"
)

func (r *MySqlRepository) FindAllProduct() []entity.Product {
	var Products []entity.Product
	result, err := db.Query("select p.id, p.description, p.name, c.id from `product` p left join category c on c.id = p.category_id")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var Product entity.Product

		err := result.Scan(&Product.ProductId, &Product.Description, &Product.Name, &Product.CategoryId)
		if err != nil {
			panic(err.Error())
		}
		Products = append(Products, Product)
	}
	return Products
}

func (r *MySqlRepository) AddProduct(product entity.Product) {
	description := product.Description
	name := product.Name
	categoryId := product.CategoryId

	stmt, err := db.Prepare("insert into `product` (`description`,`category_id`,`name`) values (?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(description, categoryId, name)
	if err != nil {
		panic(err.Error())
	}
}

func (r *MySqlRepository) UpdateProduct(product entity.Product) {
	stmt, err := db.Prepare("update `product` set description = ?, category_id = ?, name = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(product.Description, product.CategoryId, product.Name, product.ProductId)
	if err != nil {
		panic(err.Error())
	}

}

func (r *MySqlRepository) DeleteProduct(productId string) {
	stmt, err := db.Prepare("delete from `orderToProduct` where product_id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(productId)
	if err != nil {
		panic(err.Error())
	}

	stmt, err = db.Prepare("delete from `product` where id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(productId)
	if err != nil {
		panic(err.Error())
	}
}
