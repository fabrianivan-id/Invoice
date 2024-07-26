package v1

import (
	"context"

	"esb-test/src/app"
	ftpRepo "esb-test/src/repository/ftp"
	invoiceRepo "esb-test/src/repository/invoice"
	invoiceSvc "esb-test/src/v1/service/invoice"

	atomicSQLX "esb-test/library/atomic/sqlx"
	"esb-test/library/logger"
)

type repositories struct {
	AtomicSessionProvider *atomicSQLX.SqlxAtomicSessionProvider
	Invoice               *invoiceRepo.InvoiceRepository
	FTP                   *ftpRepo.FtpRepository
}

type services struct {
	InvoiceSvc *invoiceSvc.InvoiceService
}

type Dependency struct {
	Repositories *repositories
	Services     *services
}

func initRepositories(ctx context.Context) *repositories {
	var r repositories
	var err error

	r.AtomicSessionProvider = atomicSQLX.NewSqlxAtomicSessionProvider(app.DB(), app.Tracer())
	r.Invoice, err = invoiceRepo.InitInvoiceRepository(ctx, app.DB())
	if err != nil {
		logger.GetLogger(ctx).Fatal("init invoice repo err: ", err)
	}

	r.FTP, err = ftpRepo.InitFTPRepository(ctx, app.FTP())
	if err != nil {
		logger.GetLogger(ctx).Fatal("init ftp repo err: ", err)
	}

	return &r
}

func initServices(ctx context.Context, r *repositories) *services {
	return &services{
		InvoiceSvc: invoiceSvc.InitInvoiceService(r.Invoice, r.AtomicSessionProvider),
	}
}

func Dependencies(ctx context.Context) *Dependency {
	repositories := initRepositories(ctx)
	services := initServices(ctx, repositories)

	return &Dependency{
		Repositories: repositories,
		Services:     services,
	}
}
