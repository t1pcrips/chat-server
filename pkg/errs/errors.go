package errs

import "errors"

var (
	ErrNotFound          = errors.New("not found")
	ErrInvalidInput      = errors.New("invalid input")
	ErrDatabaseError     = errors.New("database error")
	ErrCommitTransaction = errors.New("failed to commit transaction")
	ErrNoRows            = errors.New("no rows affected")
	ErrBeginTransaction  = errors.New("failed to begin transaction")
	ErrQueryRowScan      = errors.New("failed to scan query result")
	ErrCreateQuery       = errors.New("failed to create query")
	ErrExec              = errors.New("failed to execute query")
	ErrCreateMember      = errors.New("failed to create member")
	ErrCreateChat        = errors.New("failed to create chat")
	ErrDeleteChat        = errors.New("failed to delete chat")
	ErrSendMessage       = errors.New("failed to send message")
	ErrConnectToDb       = errors.New("failed to connect to database")
	ErrChatNotFound      = errors.New("chat not found")
	ErrMemberNotFound    = errors.New("member not found")
)
