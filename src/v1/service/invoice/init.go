package invoice

import (
	"context"

	"esb-test/src/entity"
	"esb-test/src/v1/contract"
)

type InvoiceRepository interface {
	GetInvoiceList(ctx context.Context, params contract.GetListParam) ([]*contract.InvoiceResponseBody, error)
	GetInvoiceCount(ctx context.Context, params contract.GetListParam) (int64, error)
	GetInvoiceById(ctx context.Context, id int) (*entity.Invoice, error)
	Create(ctx context.Context, data *entity.InvoiceData) (int64, error)
	// Delete(ctx context.Context, id int) (int64, error)
	// Update(ctx context.Context, data *entity.Invoice) (int64, error)
}
