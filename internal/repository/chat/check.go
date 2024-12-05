package chat

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/t1pcrips/chat-service/pkg/errs"
	"github.com/t1pcrips/platform-pkg/pkg/database"
)

func (repo *ChatRepositoryImpl) CheckChatExists(ctx context.Context, chatId int64) (bool, error) {
	buidlerCheckChat := squirrel.Select("1").
		From(tableChats).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumn: chatId})

	query, args, err := buidlerCheckChat.ToSql()
	if err != nil {
		return false, errs.ErrCreateQuery
	}

	q := database.Query{
		Name:     "check chat repository - check chat",
		QueryRow: query,
	}

	var exists string

	err = repo.db.DB().ScanOneContext(ctx, &exists, q, args...)
	if err != nil {
		return false, errs.ErrQueryRowScan
	}

	return exists == "1", nil
}
