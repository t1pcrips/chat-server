package chat

import (
	"chat-server/internal/repository"
	"chat-server/pkg/methods"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"time"
)

const (
	tableChats      = "chats"
	tableMessages   = "messages"
	tableMembers    = "chat_members"
	idColumn        = "id"
	chatIdColumn    = "chat_id"
	usernameColumn  = "username"
	idChatColumn    = "id_chat"
	fromColumn      = "from_username"
	textColumn      = "text"
	timestampColumn = "timestamp"
	createdAtColumn = "created_at"
	returningId     = "RETURNING id"
)

type ChatRepository struct {
	pool   *pgxpool.Pool
	logger *zerolog.Logger
}

func NewChatRepository(pool *pgxpool.Pool, logger *zerolog.Logger) *ChatRepository {
	return &ChatRepository{
		pool:   pool,
		logger: logger,
	}
}

func (repo *ChatRepository) CreateChat(ctx context.Context, usernames []string) (int64, error) {
	tx, err := repo.pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to begin acquires a connection from the Pool and starts a transaction: %w", err)
	}

	builderCreateChats := squirrel.Insert(tableChats).
		PlaceholderFormat(squirrel.Dollar).
		Columns(createdAtColumn).
		Values(time.Now()).
		Suffix(returningId)

	query, args, err := builderCreateChats.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed create chat into a SQL string and bound args: %w", err)
	}

	var chatId int64

	err = tx.QueryRow(ctx, query, args...).Scan(&chatId)
	if err != nil {
		return 0, fmt.Errorf("failed to acquires a connection and executes a query: %w", err)
	}

	for _, username := range usernames {
		builderCreateChatMember := squirrel.Insert(tableMembers).
			PlaceholderFormat(squirrel.Dollar).
			Columns(chatIdColumn, usernameColumn).
			Values(chatId, username)

		query, args, err := builderCreateChatMember.ToSql()
		if err != nil {
			return 0, fmt.Errorf("failed create query chat member: %w", err)
		}

		_, err = tx.Exec(ctx, query, args...)
		if err != nil {
			return 0, fmt.Errorf("failed to create exec: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to commit create transaction: %w", err)
	}

	return chatId, nil
}

func (repo *ChatRepository) DeleteChat(ctx context.Context, chatId int64) error {
	tx, err := repo.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin acquires a connection from the Pool and starts a transaction: %w", err)
	}

	err = methods.DeleteInPostgres(ctx, tx, tableMembers, chatIdColumn, chatId)
	if err != nil {
		return err
	}

	err = methods.DeleteInPostgres(ctx, tx, tableMessages, chatIdColumn, chatId)
	if err != nil {
		if err.Error() != "no rows to delete in messages" {
			return err
		}
	}

	err = methods.DeleteInPostgres(ctx, tx, tableChats, idColumn, chatId)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit delete transaction: %w", err)
	}

	return nil
}

func (repo *ChatRepository) SendMessage(ctx context.Context, req *repository.SendMessageRequest) error {
	builderSearchMember := squirrel.Select(chatIdColumn).
		From(tableMembers).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{usernameColumn: req.From})

	query, args, err := builderSearchMember.ToSql()
	if err != nil {
		return fmt.Errorf("failed to create - select query chat member: %w", err)
	}

	var chatId int64

	err = repo.pool.QueryRow(ctx, query, args...).Scan(&chatId)
	if err != nil {
		return fmt.Errorf("failed to acquires a connection and executes a query: %w", err)
	}

	builderSendMessage := squirrel.Insert(tableMessages).
		PlaceholderFormat(squirrel.Dollar).
		Columns(chatIdColumn, fromColumn, textColumn, timestampColumn).
		Values(chatId, req.From, req.Text, req.TimeSend)

	query, args, err = builderSendMessage.ToSql()
	if err != nil {
		return fmt.Errorf("failed to create - insert query messages: %w", err)
	}

	result, err := repo.pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to create exec: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("no rows to add in Messages")
	}

	return nil
}
