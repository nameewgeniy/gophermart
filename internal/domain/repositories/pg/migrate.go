package pg

import (
	"context"
	"embed"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func (p Pg) migrationUp(ctx context.Context) error {

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(p.mDialect); err != nil {
		return err
	}

	if err := goose.UpContext(ctx, p.db, p.mDir); err != nil {
		return err
	}

	return nil
}

func (p Pg) migrationDown() error {

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(p.mDialect); err != nil {
		return err
	}

	if err := goose.Down(p.db, p.mDir); err != nil {
		return err
	}

	return nil
}
