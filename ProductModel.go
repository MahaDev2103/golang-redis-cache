package dbmodel

type Product struct {
	Name            string `db:"name"`
	ProductType     string `db:"ptype"`
	ManufactureDate string `db:"manufacture_date"`
	Seller          string `db:"seller"`
	Availability    string `db:"expire_date"`
}
