package entity

type Customer struct {
	ModelID
	ModelLogTime
	CustomerData
}

type CustomerData struct {
	CustomerName string `db:"customer_name"` // Name of the customer
	Email        string `db:"email"`         // Email address of the customer
	Address      string `db:"address"`       // Address of the customer
	Phone        string `db:"phone"`         // Phone number of the customer
}
