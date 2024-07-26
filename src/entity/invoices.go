package entity

import "time"

type Invoice struct {
	ModelID
	ModelLogTime
	InvoiceData
}

type InvoiceData struct {
	Subject    string        `db:"subject"`     // E.g., "Spring Marketing Campaign"
	CustomerID int           `db:"customer_id"` // Assuming a foreign key to a Customer table
	DueDate    time.Time     `db:"due_date"`
	Status     string        `db:"status"` // E.g., "Paid", "Unpaid"
	Items      []InvoiceItem `db:"-"`      // Nested relationship with items
}

type InvoiceItem struct {
	ModelID
	ModelLogTime
	InvoiceItemData
}

type InvoiceItemData struct {
	InvoiceID int     `db:"invoice_id"`
	ProductID int     `db:"product_id"` // Foreign key to the Product table
	ItemName  string  `db:"item_name"`  // Save current product name into invoice item
	Quantity  int     `db:"quantity"`
	UnitPrice float64 `db:"unit_price"` // Save current product price into invoice item
}
