package response

import (
	"context"
	"net/http"

	"esb-test/src/middleware/request"

	i18n_err "esb-test/library/i18n/errors"
	"esb-test/library/logger"
)

func JSONSuccessResponse(ctx context.Context, w http.ResponseWriter, data interface{}) {
	JSONResponse(ctx, w, createSuccessResponse(data, request.GetRequestID(ctx)), http.StatusOK)
}

func JSONUnauthorizedResponse(ctx context.Context, w http.ResponseWriter) {
	JSONResponse(ctx, w, createErrorResponse(i18n_err.ErrUnauthorized, request.GetRequestID(ctx), request.GetLanguage(ctx)),
		http.StatusUnauthorized)
}

func JSONInternalErrorResponse(ctx context.Context, w http.ResponseWriter) {
	JSONResponse(ctx, w, createErrorResponse(i18n_err.ErrInternalServer, request.GetRequestID(ctx), request.GetLanguage(ctx)),
		http.StatusInternalServerError)
}

func JSONBadRequestResponse(ctx context.Context, w http.ResponseWriter) {
	JSONResponse(ctx, w, createErrorResponse(i18n_err.ErrBadRequest, request.GetRequestID(ctx), request.GetLanguage(ctx)),
		http.StatusBadRequest)
}

func JSONUnprocessableEntity(ctx context.Context, w http.ResponseWriter, err i18n_err.I18nError) {
	logger.GetLogger(ctx).Error("JSONUnprocessableEntity err:", err)
	JSONResponse(ctx, w, createErrorResponse(err, request.GetRequestID(ctx), request.GetLanguage(ctx)),
		http.StatusUnprocessableEntity)
}
