package invoice

import (
	"context"

	"esb-test/src/app"
	"esb-test/src/entity"
	"esb-test/src/v1/contract"

	"esb-test/library/logger"
)

const (
	capitalLetters    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	smallLetters      = "abcdefghijklmnopqrstuvwxyz"
	digits            = "0123456789"
	passwordMaxLength = 12
)

func (m *InvoiceRepository) GetInvoiceList(ctx context.Context, params contract.GetListParam) ([]*contract.InvoiceResponseBody, error) {
	ctx, span := app.Tracer().Start(ctx, "GetInvoiceListRepository")
	defer span.End()

	var invoices []*contract.InvoiceResponseBody

	params.Sort = "desc"

	query := masterQueries[BaseQuery]
	query, params = BuildFilter(query, params)

	rows, err := m.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		logger.GetLogger(ctx).Error("GetInvoiceList err: ", err)
		return invoices, err
	}

	for rows.Next() {
		var invoice contract.InvoiceResponseBody
		err = rows.StructScan(&invoice)
		if err != nil {
			logger.GetLogger(ctx).Error("GetInvoiceList err: ", err)
			return invoices, err
		}
		invoices = append(invoices, &invoice)
	}

	return invoices, nil
}

func (m *InvoiceRepository) GetInvoiceById(ctx context.Context, id int) (*entity.Invoice, error) {
	ctx, span := app.Tracer().Start(ctx, "GetInvoiceByIdRepository")
	defer span.End()

	stmt, err := m.getStatement(ctx, GetById)
	if err != nil {
		logger.GetLogger(ctx).Error("getStatement err: ", err)
		return nil, err
	}

	var invoice entity.Invoice
	err = stmt.GetContext(ctx, &invoice, id)
	if err != nil {
		logger.GetLogger(ctx).Error("getInvoiceById err: ", err)
		return nil, err
	}

	// Get all items for the invoice
	itemsStmt, err := m.getStatement(ctx, GetInvoiceItemByInvoiceId)
	if err != nil {
		logger.GetLogger(ctx).Error("getStatement err: ", err)
		return nil, err
	}

	var items []entity.InvoiceItem
	err = itemsStmt.SelectContext(ctx, &items, invoice.ID)
	if err != nil {
		logger.GetLogger(ctx).Error("getInvoiceItems err: ", err)
		return nil, err
	}

	invoice.Items = items

	return &invoice, nil
}

func (m *InvoiceRepository) GetInvoiceCount(ctx context.Context, params contract.GetListParam) (int64, error) {
	ctx, span := app.Tracer().Start(ctx, "GetInvoiceCountRepository")
	defer span.End()

	var count int64
	query := masterQueries[BaseQueryCount]
	query, params = BuildFilter(query, params)

	rows, err := m.db.NamedQueryContext(ctx, query, params)
	if err != nil {
		logger.GetLogger(ctx).Error("GetInvoiceCount err: ", err)
		return count, err
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			logger.GetLogger(ctx).Error("GetInvoiceCount err: ", err)
			return count, err
		}
	}

	return count, nil
}

func (m *InvoiceRepository) Create(ctx context.Context, data *entity.InvoiceData) (int64, error) {
	_, span := app.Tracer().Start(ctx, "CreateInvoiceRepository")
	defer span.End()

	var invoice entity.Invoice
	namedStmt, err := m.getNamedStatement(ctx, Insert)
	if err != nil {
		logger.GetLogger(ctx).Error("getNamedStatement err: ", err)
		return 0, err
	}

	if err = namedStmt.GetContext(ctx, &invoice, data); err != nil {
		logger.GetLogger(ctx).Error("CreateInvoiceRepository err: ", err)
		return 0, err
	}

	return invoice.ID, nil
}

func (m *InvoiceRepository) Delete(ctx context.Context, id int) (int64, error) {
	ctx, span := app.Tracer().Start(ctx, "DeleteInvoiceRepository")
	defer span.End()

	var rowsAffected int64
	res, err := m.masterStmts[Delete].ExecContext(ctx, id)
	if err != nil {
		logger.GetLogger(ctx).Error("DeleteInvoiceMethod err: ", err)
		return rowsAffected, err
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		logger.GetLogger(ctx).Error("Get rows affected err: ", err)
		return rowsAffected, err
	}

	return rowsAffected, nil
}

func (m *InvoiceRepository) Update(ctx context.Context, data *entity.Invoice) (int64, error) {
	_, span := app.Tracer().Start(ctx, "UpdateInvoiceRepository")
	defer span.End()

	namedStmt, err := m.getNamedStatement(ctx, Update)
	if err != nil {
		logger.GetLogger(ctx).Error("getNamedStatement err: ", err)
		return 0, err
	}

	_, err = namedStmt.ExecContext(ctx, &data)
	if err != nil {
		logger.GetLogger(ctx).Error("UpdateInvoice err: ", err)
		return 0, err
	}

	return data.ID, nil
}

func BuildFilter(query string, param contract.GetListParam) (string, contract.GetListParam) {
	if param.Keyword != "" {
		query += " AND subject ILIKE :keyword"
		param.Keyword = "%" + param.Keyword + "%"
	}

	if param.Sort == "desc" {
		query += " ORDER BY id DESC"
	}

	if param.Sort == "asc" {
		query += " ORDER BY id ASC"
	}

	query += " LIMIT :limit OFFSET :offset"
	return query, param
}
