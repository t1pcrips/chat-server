package chat

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/t1pcrips/chat-service/pkg/errs"
	"github.com/t1pcrips/platform-pkg/pkg/database"
	"time"
)

func (repo *ChatRepositoryImpl) CreateChat(ctx context.Context) (int64, error) {
	builderCreateChat := squirrel.Insert(tableChats).
		PlaceholderFormat(squirrel.Dollar).
		Columns(createdAtColumn).
		Values(time.Now()).
		Suffix(returningId)

	query, args, err := builderCreateChat.ToSql()
	if err != nil {
		return 0, errs.ErrCreateQuery
	}

	q := database.Query{
		Name:     "chat repository - createChat",
		QueryRow: query,
	}

	var chatId int64

	err = repo.db.DB().ScanOneContext(ctx, &chatId, q, args...)
	if err != nil {
		return 0, errs.ErrQueryRowScan
	}

	return chatId, nil
}
