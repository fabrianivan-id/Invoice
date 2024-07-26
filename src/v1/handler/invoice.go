package handler

import (
	"net/http"
	// "strconv"

	"esb-test/src/app"
	appErr "esb-test/src/errors"
	"esb-test/src/middleware/response"
	"esb-test/src/v1/contract"

	"esb-test/library/logger"
)

// GetListInvoiceHandler godoc
//
//	@Summary		Get list invoice
//	@Description	get invoice pagination by page, limit, keyword
//	@Tags			v1-invoice
//	@Accept			json
//	@Produce		json
//	@Param			page	query		string	false	"page"
//	@Param			limit	query		string	false	"limit"
//	@Param			keyword	query		string	false	"keyword"
//	@Success		200		{object}	response.Response{data=contract.InvoiceListResponse}
//	@Failure		500		{object}	response.Response
//	@Router			/v1/invoice [get]
func GetListInvoiceHandler(svc InvoiceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := app.Tracer().Start(r.Context(), "GetListInvoiceHandler")
		defer span.End()

		getListParam, err := contract.ValidateAndBuildGetListRequest(r)
		if err != nil {
			logger.GetLogger(r.Context()).Error("ValidateQuery err:", err)
			response.JSONBadRequestResponse(ctx, w)
			return
		}

		res, err := svc.GetInvoiceList(ctx, *getListParam)
		if err != nil {
			logger.GetLogger(r.Context()).Error("GetInvoiceList err:", err)
			response.JSONInternalErrorResponse(ctx, w)
			return
		}

		response.JSONSuccess(ctx, w, http.StatusOK, res)
	}
}

// GetInvoiceHandler godoc
//
//	@Summary		Get invoice
//	@Description	get invoice by id, including items
//	@Tags			v1-invoice
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int	true	"Invoice ID"
//	@Success		200				{object}	response.Response{data=contract.InvoiceResponseBody}
//	@Failure		400				{object}	response.Response
//	@Failure		404				{object}	response.Response
//	@Failure		500				{object}	response.Response
//	@Router			/v1/invoice/{id} 	[get]
func GetInvoiceHandler(svc InvoiceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := app.Tracer().Start(r.Context(), "GetInvoiceHandler")
		defer span.End()

		id, err := contract.ValidateIDRequest(r)
		if err != nil {
			logger.GetLogger(r.Context()).Error("ValidateInvoiceIDRequest err:", err)
			response.JSONBadRequestResponse(ctx, w)
			return
		}

		resp, err := svc.GetInvoiceByID(ctx, int64(id))
		if err != nil {
			logger.GetLogger(r.Context()).Error("GetInvoiceHandler err:", err)
			if err == appErr.ErrInvoiceIdNotFound {
				response.JSONError(ctx, w, http.StatusNotFound, err)
				return
			}
			response.JSONInternalErrorResponse(ctx, w)
			return
		}
		response.JSONSuccess(ctx, w, http.StatusOK, resp)
	}
}

// CreateInvoiceHandler godoc
//
//	@Summary		Create invoice
//	@Description	create invoice
//	@Tags			v1-invoice
//	@Accept			json
//	@Produce		json
//	@Param			invoice		body		contract.InvoiceRequest	true	"invoice_information"
//	@Success		200			{object}	response.Response{data=contract.InvoiceResponseBody}
//	@Failure		500			{object}	response.Response
//	@Router			/v1/invoice 	[post]
func CreateInvoiceHandler(svc InvoiceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := app.Tracer().Start(r.Context(), "CreateInvoiceHandler")
		defer span.End()

		body, err := contract.ValidateInvoiceRequestBody(r)
		if err != nil {
			logger.GetLogger(r.Context()).Error("ValidateInvoiceRequestBody err:", err)
			response.JSONBadRequestResponse(ctx, w)
			return
		}

		res, err := svc.CreateInvoice(ctx, &body)
		if err != nil {
			logger.GetLogger(r.Context()).Error("CreateInvoiceHandler err:", err)
			response.JSONInternalErrorResponse(ctx, w)
			return
		}

		response.JSONSuccess(ctx, w, http.StatusOK, res)
	}
}

// DeleteInvoiceHandler godoc
//
//	@Summary		Delete invoice
//	@Description	delete invoice by id
//	@Tags			v1-invoice
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int	true	"Invoice ID"
//	@Success		200				{object}	response.Response
//	@Failure		403				{object}	response.Response
//	@Failure		500				{object}	response.Response
//	@Router			/v1/invoice/{id} 	[delete]
// func DeleteInvoiceHandler(svc InvoiceService) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		ctx, span := app.Tracer().Start(r.Context(), "DeleteInvoiceHandler")
// 		defer span.End()

// 		invoiceId := chi.URLParam(r, "id")
// 		id, _ := strconv.Atoi(invoiceId)
// 		err := svc.DeleteInvoice(ctx, int64(id))
// 		if err != nil {
// 			switch err {
// 			case appErr.ErrInvoiceIdNotFound:
// 				response.JSONError(ctx, w, http.StatusForbidden, err)
// 			default:
// 				response.JSONInternalErrorResponse(ctx, w)
// 			}
// 			return
// 		}
// 		response.JSONSuccessResponse(ctx, w, nil)
// 	}
// }

// UpdateInvoiceHandler godoc
//
//	@Summary		Update invoice
//	@Description	update invoice
//	@Tags			v1-invoice
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int						true	"Invoice ID"
//	@Param			invoice			body		contract.InvoiceRequest	true	"invoice_information"
//	@Success		200				{object}	response.Response
//	@Failure		404				{object}	response.Response
//	@Failure		500				{object}	response.Response
//	@Router			/v1/invoice/{id} 	[put]
// func UpdateInvoiceHandler(svc InvoiceService) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		ctx, span := app.Tracer().Start(r.Context(), "UpdateInvoiceHandler")
// 		defer span.End()

// 		id, err := contract.ValidateIDRequest(r)
// 		if err != nil {
// 			logger.GetLogger(r.Context()).Error("ValidateIDRequest err:", err)
// 			response.JSONBadRequestResponse(ctx, w)
// 			return
// 		}

// 		params, err := contract.ValidateInvoiceRequestBody(r)
// 		if err != nil {
// 			logger.GetLogger(r.Context()).Error("ValidateInvoiceRequestBody err:", err)
// 			response.JSONBadRequestResponse(ctx, w)
// 			return
// 		}

// 		res, err := svc.UpdateInvoice(ctx, int64(id), params)
// 		if err != nil {
// 			logger.GetLogger(r.Context()).Error("UpdateInvoiceHandler err:", err)
// 			if err == appErr.ErrInvoiceIdNotFound {
// 				response.JSONError(ctx, w, http.StatusNotFound, err)
// 				return
// 			}

// 			response.JSONInternalErrorResponse(ctx, w)
// 			return
// 		}

// 		response.JSONSuccess(ctx, w, http.StatusOK, map[string]interface{}{"id": res.ID})
// 	}
// }
