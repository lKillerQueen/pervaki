package titlerepo

import (
	"context"
	"errors"
	"pervaki/database/titlerepo/query"
	"pervaki/model"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Repository struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
}

func NewRepository(logger *zap.SugaredLogger, db *sqlx.DB) Repository {
	return Repository{
		logger: logger,
		db:     db,
	}
}

func (r Repository) Upsert(ctx context.Context, title model.Title) error {
	var data = MapServiceToDb(title)

	if len(data.Code) == 0 {
		return errors.New("data empty")
	}

	_, err := r.db.ExecContext(ctx, query.UpsertTitleSql, data.Code, data.NameRu)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) UpsertThroughBuilder(ctx context.Context, title model.Title) error {
	var (
		psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

		data = MapServiceToDb(title)
	)

	if len(data.Code) == 0 {
		return errors.New("data empty")
	}
	builder := psql.Insert("title").
		Columns("code", "name_ru").
		Values(title.Code, title.NameRu)
	builder = builder.Suffix("on conflict (code) do update set name_ru = excluded.name_ru")
	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) InCache(ctx context.Context, title model.Title) error {
	var (
		psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

		data = MapServiceToDb(title)
	)

	if len(data.Code) == 0 {
		return errors.New("data empty")
	}

	builder := psql.Insert("title").
		Columns("code", "name_ru").
		Values(title.Code, title.NameRu)
	builder = builder.Suffix("on conflict (code) do update set name_ru = excluded.name_ru")
	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) RemoveRow(ctx context.Context, title model.Title) error {
	var (
		psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

		data = MapServiceToDb(title)
	)

	if len(data.Code) == 0 {
		return errors.New("data empty")
	}

	builder := psql.Insert("title").
		Columns("code", "name_ru").
		Values(title.Code, title.NameRu)
	builder = builder.Suffix("on conflict (code) do update set name_ru = excluded.name_ru")
	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) ShowAll(ctx context.Context) (string, error) {

	rows := []model.Title{}
	strRow := ""
	err := r.db.Select(&rows, "SELECT name_ru, code FROM title")
	if err != nil {
		return "", err
	}

	for i := 0; i < len(rows); i++ {
		strRow += "/n" + rows[i].Code + " " + rows[i].NameRu
	}
	return strRow, nil
}

func (r Repository) ClearCache(ctx context.Context, title model.Title) error {
	var (
		psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

		data = MapServiceToDb(title)
	)

	if len(data.Code) == 0 {
		return errors.New("data empty")
	}

	builder := psql.Insert("title").
		Columns("code", "name_ru").
		Values(title.Code, title.NameRu)
	builder = builder.Suffix("on conflict (code) do update set name_ru = excluded.name_ru")
	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return err
	}

	return nil
}
