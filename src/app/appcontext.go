package app

import (
	"context"
	"path/filepath"
	"runtime"

	"esb-test/library/i18n"
	"esb-test/library/logger"
	mysql "esb-test/library/mysql"
	"esb-test/library/storage/ftp"
	"esb-test/library/tracer"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type appContext struct {
	db               *sqlx.DB
	requestValidator *validator.Validate
	ftp              *ftp.FTPClient
	cfg              *Configuration
	tracer           trace.Tracer
}

var appCtx appContext
var appTransFile = func() string {
	_, f, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filepath.Dir(f))

	// Return the project root directory path.
	return filepath.Join(basepath, "translation")
}()

func Init(ctx context.Context) error {
	logger.Init(ctx)

	cfg, err := InitConfig(ctx)
	if err != nil {
		return err
	}

	if err := i18n.Init(ctx, cfg.Translation.FilePath, appTransFile, cfg.Translation.DefaultLanguage); err != nil {
		panic(err)
	}

	ftp := ftp.InitFTP(ftp.FTPConfig{
		Host:            cfg.Ftp.Host,
		Username:        cfg.Ftp.Username,
		Password:        cfg.Ftp.Password,
		BaseDomainImage: cfg.Ftp.BaseDomain,
		RootDirectory:   cfg.Ftp.Root,
	})

	db, err := mysql.InitSQLX(ctx, mysql.MySQLConfig{
		ConnectionUrl:      cfg.MySQL.ConnURI,
		MaxPoolSize:        cfg.MySQL.MaxPoolSize,
		MaxIdleConnections: cfg.MySQL.MaxIdleConnections,
		ConnMaxIdleTime:    cfg.MySQL.MaxIdleTime,
		ConnMaxLifeTime:    cfg.MySQL.MaxLifeTime,
	})
	if err != nil {
		return err
	}

	_, err = tracer.Init(ctx, cfg.ServiceName, "0.0.1")
	if err != nil {
		return err
	}

	appCtx = appContext{
		db:               db,
		requestValidator: validator.New(),
		cfg:              cfg,
		ftp:              &ftp,
		tracer:           otel.Tracer(cfg.ServiceName),
	}

	return nil
}

func RequestValidator() *validator.Validate {
	return appCtx.requestValidator
}

func DB() *sqlx.DB {
	return appCtx.db
}

func Config() Configuration {
	return *appCtx.cfg
}

func Tracer() trace.Tracer {
	return appCtx.tracer
}

func FTP() ftp.FTPClient {
	return *appCtx.ftp
}
