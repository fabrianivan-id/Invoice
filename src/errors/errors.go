package errors

import (
	i18n_err "esb-test/library/i18n/errors"
)

var (
	ErrExpiredToken      = i18n_err.NewI18nError("err_expired_token")
	ErrInvoiceIdNotFound = i18n_err.NewI18nError("err_invoice_id_not_found")
	ErrInvoicesDuplicate = i18n_err.NewI18nError("err_invoices_data_duplicate")
)
