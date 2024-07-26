package contract

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"esb-test/library/logger"
	"esb-test/library/utils"
	"esb-test/src/app"

	"github.com/go-chi/chi/v5"
)

type InvoiceResponseBody struct {
	ID         int64                `json:"id,omitempty"`
	IssueDate  time.Time            `json:"created_at,omitempty"`
	Subject    string               `json:"subject"`
	CustomerID int                  `json:"customer_id"`
	DueDate    time.Time            `json:"due_date"`
	Status     string               `json:"status"`
	TotalItems int                  `json:"total_items"`
	Items      []InvoiceItemRequest `json:"items" validate:"required"`
	UpdatedAt  time.Time            `json:"updated_at,omitempty"`
	DeletedAt  *time.Time           `json:"deleted_at,omitempty"`
}

type InvoiceListResponse struct {
	Data       []*InvoiceResponseBody `json:"data"`
	Pagination *utils.Pagination      `json:"pagination"`
}

type InvoiceRequest struct {
	IssueDate  string               `json:"issue_date" validate:"required"`
	Subject    string               `json:"subject" validate:"required"`
	CustomerID int                  `json:"customer_id" validate:"required"`
	DueDate    time.Time            `json:"due_date" validate:"required"`
	Status     string               `json:"status" validate:"required"`
	Items      []InvoiceItemRequest `json:"items" validate:"required"`
}

type InvoiceItemRequest struct {
	ProductID int     `json:"product_id" validate:"required"`
	ItemName  string  `json:"item_name" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	UnitPrice float64 `json:"unit_price" validate:"required"`
	Total     float64 `json:"total" validate:"required"`
}

type InvoiceResponse struct {
	ID            int64                 `json:"id"`
	IssueDate     string                `json:"issue_date"`
	Subject       string                `json:"subject"`
	CustomerID    int                   `json:"customer_id"`
	DueDate       time.Time             `json:"due_date"`
	Status        string                `json:"status"`
	Items         []InvoiceItemResponse `json:"items,omitempty"`
	SubtotalPrice int                   `json:"subtotal_price"`
	TotalPrice    int                   `json:"total_price"`
	CreatedAt     time.Time             `json:"created_at,omitempty"`
	UpdatedAt     time.Time             `json:"updated_at,omitempty"`
	DeletedAt     *time.Time            `json:"deleted_at,omitempty"`
}

type InvoiceItemResponse struct {
	ProductID int     `json:"product_id"`
	ItemName  string  `json:"item_name"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	Total     float64 `json:"total"`
}

func ValidateInvoiceRequestBody(r *http.Request) (request InvoiceRequest, err error) {
	_, span := app.Tracer().Start(r.Context(), "ValidateInvoiceRequestBody")
	defer span.End()

	var payload InvoiceRequest
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		logger.GetLogger(r.Context()).Error("read request body err: ", err)
		return payload, err
	}

	if err := json.Unmarshal(bodyByte, &payload); err != nil {
		logger.GetLogger(r.Context()).Error("unmarshal request body err: ", err)
		return payload, err
	}

	if err := app.RequestValidator().Struct(payload); err != nil {
		logger.GetLogger(r.Context()).Error("validate request body err: ", err)
		return payload, err
	}

	return payload, nil
}

func ValidateIDRequest(r *http.Request) (int, error) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, err
	}

	return id, nil
}
