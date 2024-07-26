package entity

type Product struct {
	ModelID
	ModelLogTime
	ProductData
}

type ProductData struct {
	Name string `db:"name"`
	Type string `db:"type"`
}
