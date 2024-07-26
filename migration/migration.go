package migration

import (
	"context"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	logger "esb-test/library/logger"
	ms "esb-test/library/mysql"
	"esb-test/src/app"
)

const (
	migrateLogIdentifier = "esb-test"
)

type MigrationService interface {
	Up(context.Context) error
	Rollback(context.Context) error
	Version(context.Context) (int, bool, error)
}

type migrationService struct {
	driver  database.Driver
	migrate *migrate.Migrate
}

func New(ctx context.Context, cfg app.MySQL) (MigrationService, error) {
	msCfg := ms.MySQLConfig{
		ConnectionUrl: cfg.ConnURI,
	}

	sqlxDB, err := ms.InitSQLX(ctx, msCfg)
	if err != nil {
		logger.GetLogger(ctx).Errorf("error connecting to sqlxDB url:%s, err: %v", msCfg.ConnectionUrl, err)
		return nil, err
	}

	databaseInstance, err := mysql.WithInstance(sqlxDB.DB, &mysql.Config{})
	if err != nil {
		logger.GetLogger(ctx).Errorf("go-migrate mysql drv init failed: %v", err)
		return nil, err
	}

	migrate, err := migrate.NewWithDatabaseInstance("file://migration/sql",
		migrateLogIdentifier, databaseInstance)
	if err != nil {
		logger.GetLogger(ctx).Errorf("migrate init failed %v", err)
		return nil, err
	}

	return migrationService{
		driver:  databaseInstance,
		migrate: migrate,
	}, nil
}

func (s migrationService) Up(ctx context.Context) error {
	currVersion, _, err := s.Version(ctx)
	if err != nil {
		logger.GetLogger(ctx).Error("Failed get current version err: ", err)
		return err
	}

	logger.GetLogger(ctx).Infof("Running migration from version: %d", currVersion)
	if err := s.migrate.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.GetLogger(ctx).Info("No Changes")
			return nil
		}
		logger.GetLogger(ctx).Error("Failed run migrate err: ", err)
		return err
	}

	currVersion, _, _ = s.Version(ctx)
	logger.GetLogger(ctx).Info("Migration success, current version:", currVersion)
	return nil
}

func (s migrationService) Rollback(ctx context.Context) error {
	currVersion, _, err := s.Version(ctx)
	if err != nil {
		logger.GetLogger(ctx).Error("Failed get current version err: ", err)
		return err
	}

	logger.GetLogger(ctx).Infof("Rollingback 1 step from version: %d", currVersion)

	if err := s.migrate.Steps(-1); err != nil {
		logger.GetLogger(ctx).Errorf("Failed to rollback, err:%v", err)
		return err
	}

	currVersion, _, _ = s.Version(ctx)
	logger.GetLogger(ctx).Infof("Rollback success, current version:%d", currVersion)
	return nil
}

func (s migrationService) Version(ctx context.Context) (int, bool, error) {
	currVersion, dirty, err := s.driver.Version()
	if err != nil {
		logger.GetLogger(ctx).Errorf("Failed to get version:%w", err)
		return 0, false, err
	}
	return currVersion, dirty, nil
}
