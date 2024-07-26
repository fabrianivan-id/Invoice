package invoice

import (
	"context"
	"database/sql"

	"esb-test/src/app"
	"esb-test/src/entity"
	appErr "esb-test/src/errors"

	"esb-test/src/v1/contract"

	"esb-test/library/atomic"
	"esb-test/library/logger"
	"esb-test/library/utils"
)

type InvoiceService struct {
	atomicSession atomic.AtomicSessionProvider
	invoiceRepo   InvoiceRepository
}

func InitInvoiceService(invoiceRepo InvoiceRepository, aSession atomic.AtomicSessionProvider) *InvoiceService {
	return &InvoiceService{
		atomicSession: aSession,
		invoiceRepo:   invoiceRepo,
	}
}

func (s *InvoiceService) GetInvoiceList(ctx context.Context, params contract.GetListParam) (contract.InvoiceListResponse, error) {
	ctx, span := app.Tracer().Start(ctx, "GetInvoiceListService")
	defer span.End()

	var resp contract.InvoiceListResponse
	invoices, err := s.invoiceRepo.GetInvoiceList(ctx, params)
	if err != nil {
		return resp, err
	}

	invoiceCounts, err := s.invoiceRepo.GetInvoiceCount(ctx, params)
	if err != nil {
		return resp, err
	}

	pagination := utils.GetPaginationData(params.Page, params.Limit, int(invoiceCounts))

	var invoiceResp []*contract.InvoiceResponseBody
	for _, invoice := range invoices {
		invoiceResp = append(invoiceResp, &contract.InvoiceResponseBody{
			ID:         invoice.ID,
			IssueDate:  invoice.IssueDate,
			Subject:    invoice.Subject,
			CustomerID: invoice.CustomerID,
			DueDate:    invoice.DueDate,
			Status:     invoice.Status,
			TotalItems: invoice.TotalItems,
			UpdatedAt:  invoice.UpdatedAt,
			DeletedAt:  invoice.DeletedAt,
		})
	}

	return contract.InvoiceListResponse{
		Data:       invoiceResp,
		Pagination: pagination,
	}, nil
}

func (s *InvoiceService) GetInvoice(ctx context.Context, id int64) (resp *contract.InvoiceResponseBody, err error) {
	ctx, span := app.Tracer().Start(ctx, "GetInvoiceService")
	defer span.End()

	invoice, err := s.invoiceRepo.GetInvoiceById(ctx, int(id))
	if err != nil {
		logger.GetLogger(ctx).Errorf("GetInvoiceService err: %v", err)
		if err == sql.ErrNoRows {
			err = appErr.ErrInvoiceIdNotFound
		}

		return
	}

	resp = &contract.InvoiceResponseBody{
		ID:         invoice.ID,
		IssueDate:  invoice.CreatedAt,
		Subject:    invoice.Subject,
		CustomerID: invoice.CustomerID,
		DueDate:    invoice.DueDate,
		Status:     invoice.Status,
		UpdatedAt:  invoice.UpdatedAt,
		DeletedAt:  invoice.DeletedAt,
	}

	return
}

func (s *InvoiceService) CreateInvoice(ctx context.Context, param *contract.InvoiceRequest) (resp *contract.InvoiceResponseBody, err error) {
	ctx, span := app.Tracer().Start(ctx, "CreateInvoiceService")
	defer span.End()

	invoiceData := entity.InvoiceData{
		Subject:    param.Subject,
		CustomerID: param.CustomerID,
		DueDate:    param.DueDate,
		Status:     param.Status,
	}

	err = atomic.Atomic(ctx, s.atomicSession, func(ctx context.Context) error {
		invoiceId, err := s.invoiceRepo.Create(ctx, &invoiceData)

		if err != nil {
			// Handle rollback if needed
			return err
		}

		resp = &contract.InvoiceResponseBody{
			ID:         invoiceId,
			Subject:    invoiceData.Subject,
			CustomerID: invoiceData.CustomerID,
			DueDate:    invoiceData.DueDate,
			Status:     invoiceData.Status,
		}
		return nil
	})

	return resp, nil
}

func (s *InvoiceService) GetInvoiceByID(ctx context.Context, id int64) (resp *contract.InvoiceResponseBody, err error) {
	ctx, span := app.Tracer().Start(ctx, "GetInvoiceByIDService")
	defer span.End()

	// Fetch the invoice from the repository
	invoice, err := s.invoiceRepo.GetInvoiceById(ctx, int(id))
	if err != nil {
		// Log and handle errors, such as not finding the invoice
		logger.GetLogger(ctx).Errorf("GetInvoiceByIDService err: %v", err)
		if err == sql.ErrNoRows {
			err = appErr.ErrInvoiceIdNotFound // Ensure this error is defined in your errors package
		}
		return nil, err
	}

	// Map the invoice data to the response structure
	resp = &contract.InvoiceResponseBody{
		ID:         invoice.ID,
		IssueDate:  invoice.CreatedAt,
		Subject:    invoice.Subject,
		CustomerID: invoice.CustomerID,
		DueDate:    invoice.DueDate,
		Status:     invoice.Status,
		UpdatedAt:  invoice.UpdatedAt,
		DeletedAt:  invoice.DeletedAt,
	}

	return resp, nil
}

// func (s *InvoiceService) UpdateInvoice(ctx context.Context, id int64, params contract.InvoiceRequest) (*contract.InvoiceResponseBody, error) {
// 	ctx, span := app.Tracer().Start(ctx, "UpdateInvoiceService")
// 	defer span.End()

// 	var resp contract.InvoiceResponseBody

// 	_, err := s.invoiceRepo.Update(ctx, &entity.Invoice{
// 		ModelID: entity.ModelID{
// 			ID: id,
// 		},
// 		InvoiceData: entity.InvoiceData{
// 			InvoiceID:  params.InvoiceID,
// 			InvoiceNo:  params.InvoiceNo,
// 			Subject:    params.Subject,
// 			CustomerID: params.CustomerID,
// 			DueDate:    params.DueDate,
// 			Status:     params.Status,
// 		},
// 	})

// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return &resp, appErr.ErrInvoiceIdNotFound
// 		}

// 		return &resp, err
// 	}

// 	resp = contract.InvoiceResponseBody{
// 		ID:         id,
// 		InvoiceID:  params.InvoiceID,
// 		InvoiceNo:  params.InvoiceNo,
// 		Subject:    params.Subject,
// 		CustomerID: params.CustomerID,
// 		DueDate:    params.DueDate,
// 		Status:     params.Status,
// 	}

// 	return &resp, err
// }

// func (s *InvoiceService) DeleteInvoice(ctx context.Context, id int) error {
// 	ctx, span := app.Tracer().Start(ctx, "DeleteInvoiceService")
// 	defer span.End()

// 	rowsAffected, err := s.invoiceRepo.Delete(ctx, id)
// 	if err != nil {
// 		logger.GetLogger(ctx).Error("DeleteInvoiceMethod err: ", err)
// 		return err
// 	}

// 	if rowsAffected == 0 {
// 		return appErr.ErrInvoiceIdNotFound
// 	}

// 	return nil
// }
