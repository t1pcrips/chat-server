package methods

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func DeleteInPostgres(ctx context.Context, tx pgx.Tx, table, key string, value int64) error {
	builderDeleteChatMember := squirrel.Delete(table).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{key: value})

	query, args, err := builderDeleteChatMember.ToSql()
	if err != nil {
		return fmt.Errorf("failed to create - delete query %s: %w", table, err)
	}

	result, err := tx.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to exec delete query: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("no rows to delete in %s", table)
	}

	return nil
}
