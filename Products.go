package dao

import (
	"database/sql"
	"fmt"
	"golang-redis/dbmodel"
	"log"
)

func GetProductsFromDB(db *sql.DB) map[uint8]dbmodel.Product {
	// get products from db
	// Unmarshall the resultset as <k,V> ex: <ProductID, ProductDetails>
	products := make(map[uint8]dbmodel.Product)
	var id uint8
	var name string
	var ptype string
	var manufactureDate string
	var seller string
	var expireDate string

	sqlStatement := "select product.id, product.name,product.type,product.manufacture_date,product.seller,product.expire_date from product"
	rows, err := db.Query(sqlStatement)
	fmt.Printf("get product query %v \n ", sqlStatement)
	if err != nil {
		log.Printf("Error while Querying  Records  Error : %s", err.Error())
	}
	defer rows.Close()
	//var metrics []model.Metric

	for rows.Next() {

		err = rows.Scan(&id, &name, &ptype, &manufactureDate, &seller, &expireDate)
		products[id] = dbmodel.Product{Name: name, ProductType: ptype, ManufactureDate: manufactureDate, Seller: seller, Availability: expireDate}

		if err != nil {
			panic("Error Reading Tasks from the db Row Scan ; Error : " + err.Error())
		}

	}
	return products
}
