package handler

import (
	"context"

	"esb-test/src/v1/contract"
)

type InvoiceService interface {
	GetInvoiceList(ctx context.Context, params contract.GetListParam) (contract.InvoiceListResponse, error)
	GetInvoiceByID(ctx context.Context, id int64) (*contract.InvoiceResponseBody, error)
	CreateInvoice(ctx context.Context, data *contract.InvoiceRequest) (*contract.InvoiceResponseBody, error)
	// UpdateInvoice(ctx context.Context, id int64, data *contract.InvoiceRequest) (*contract.InvoiceResponseBody, error)
	// DeleteInvoice(ctx context.Context, id int64) error
}
