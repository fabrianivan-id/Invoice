package invoice

import (
	"context"
	"fmt"

	"esb-test/library/atomic"
	atomicSqlx "esb-test/library/atomic/sqlx"
	"esb-test/library/logger"
	sqlxUtils "esb-test/library/sqlx"

	"github.com/jmoiron/sqlx"
)

const (
	AllFields         = `id, subject, customer_id, status, created_at, updated_at, deleted_at`
	GetListFields     = `id, subject, customer_id, due_date, status, created_at, updated_at`
	AllItemFields     = `id, invoice_id, product_id, item_name, quantity, unit_price, created_at, updated_at, deleted_at`
	GetListItemFields = `id, invoice_id, product_id, item_name, quantity, unit_price, created_at, updated_at`

	GetById            = iota
	GetInvoiceItemById = iota
	GetInvoiceItemByInvoiceId
	GetList
	GetCountList
	Delete
	BaseQuery
	BaseQueryCount
	CountItemsByInvoiceId

	InsertInvoiceItem = iota + 200
	Insert            = iota + 200
	Update
)

var (
	masterQueries = []string{
		GetById:                   fmt.Sprintf("SELECT %s FROM invoices WHERE id = $1", AllFields),
		GetInvoiceItemById:        fmt.Sprintf("SELECT %s FROM invoice_items WHERE id = $1", AllItemFields),
		GetInvoiceItemByInvoiceId: fmt.Sprintf("SELECT %s FROM invoice_items WHERE invoice_id = $1", AllItemFields),
		GetList:                   fmt.Sprintf("SELECT %s FROM invoices WHERE deleted_at IS NULL LIMIT $1 OFFSET $2", GetListFields),
		GetCountList:              "SELECT COUNT(*) FROM invoices WHERE deleted_at IS NULL",
		Delete:                    "DELETE FROM invoices WHERE id = $1",
		BaseQuery:                 fmt.Sprintf("SELECT %s, (SELECT COUNT(*) FROM invoice_items ii WHERE ii.invoice_id = i.id) total_item FROM invoices i WHERE deleted_at IS NULL", GetListFields),
		BaseQueryCount:            "SELECT COUNT(*) FROM invoices WHERE deleted_at IS NULL",
		CountItemsByInvoiceId:     "SELECT COUNT(*) FROM invoice_items WHERE invoice_id = $1",
	}

	masterNamedQueries = []string{
		Insert:            `INSERT INTO invoices (subject, customer_id, due_date, status, created_at, updated_at) VALUES (:subject, :customer_id, :due_date, :status, now(), now()) RETURNING id`,
		InsertInvoiceItem: `INSERT INTO invoice_items (invoice_id, product_id, item_name, quantity, unit_price, created_at, updated_at) VALUES (:invoice_id, :product_id, :item_name, :quantity, :unit_price, now(), now()) RETURNING id`,
		Update:            `UPDATE invoices SET subject = :subject, customer_id = :customer_id, due_date = :due_date, status = :status, updated_at = now() WHERE id = :id`,
	}
)

type InvoiceRepository struct {
	db               *sqlx.DB
	masterStmts      []*sqlx.Stmt
	masterNamedStmts []*sqlx.NamedStmt
}

func InitInvoiceRepository(ctx context.Context, db *sqlx.DB) (*InvoiceRepository, error) {
	stmts, err := sqlxUtils.PrepareQueries(db, masterQueries)
	if err != nil {
		logger.GetLogger(ctx).Error("PrepareQueries err:", err)
		return nil, err
	}

	namedStmts, err := sqlxUtils.PrepareNamedQueries(db, masterNamedQueries)
	if err != nil {
		logger.GetLogger(ctx).Error("PrepareNamedQueries err:", err)
		return nil, err
	}

	return &InvoiceRepository{
		db:               db,
		masterStmts:      stmts,
		masterNamedStmts: namedStmts,
	}, nil
}

func (r *InvoiceRepository) getStatement(ctx context.Context, queryId int) (*sqlx.Stmt, error) {
	var err error
	var statement *sqlx.Stmt
	if atomicSessionCtx, ok := ctx.(atomic.AtomicSessionContext); ok {
		if atomicSession, ok := atomicSessionCtx.AtomicSession.(atomicSqlx.SqlxAtomicSession); ok {
			statement, err = atomicSession.Tx().PreparexContext(ctx, masterQueries[queryId])
		} else {
			err = atomic.InvalidAtomicSessionProvider
		}
	} else {
		statement = r.masterStmts[queryId]
	}
	return statement, err
}

func (r *InvoiceRepository) getNamedStatement(ctx context.Context, queryId int) (*sqlx.NamedStmt, error) {
	var err error
	var namedStmt *sqlx.NamedStmt
	if atomicSessionCtx, ok := ctx.(atomic.AtomicSessionContext); ok {
		if atomicSession, ok := atomicSessionCtx.AtomicSession.(atomicSqlx.SqlxAtomicSession); ok {
			namedStmt, err = atomicSession.Tx().PrepareNamedContext(ctx, masterNamedQueries[queryId])
		} else {
			err = atomic.InvalidAtomicSessionProvider
		}
	} else {
		namedStmt = r.masterNamedStmts[queryId]
	}
	return namedStmt, err
}
