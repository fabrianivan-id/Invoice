package main

import (
	"context"
	"os"

	"esb-test/library/logger"
	"esb-test/migration"
	"esb-test/src/app"
)

func main() {
	ctx := context.Background()

	app.Init(ctx)
	logger.Init(ctx)

	args := os.Args
	if len(args) < 2 {
		logger.GetLogger(ctx).Fatal("Missing args. args: [up | rollback]")
	}

	migrationSvc, err := migration.New(ctx, app.Config().MySQL)
	if err != nil {
		logger.GetLogger(ctx).Fatal("Failed to initiate migration", err)
	}

	switch args[1] {
	case "up":
		migrationSvc.Up(ctx)
	case "rollback":
		migrationSvc.Rollback(ctx)
	default:
		logger.GetLogger(ctx).Fatal("Invalid migration command")
	}
}
