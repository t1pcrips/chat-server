package chat_members

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/t1pcrips/chat-service/pkg/errs"
	"github.com/t1pcrips/platform-pkg/pkg/database"
)

func (repo *MembersRepositoryImpl) CheckMemberExists(ctx context.Context, userName string) (bool, error) {
	builderCheckUser := squirrel.Select("1").
		From(tableMembers).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{usernameColumn: userName})

	query, args, err := builderCheckUser.ToSql()
	if err != nil {
		return false, errs.ErrCreateQuery
	}

	q := database.Query{
		Name:     "check members repository - check user",
		QueryRow: query,
	}

	var exists string

	err = repo.db.DB().ScanOneContext(ctx, &exists, q, args...)
	if err != nil {
		return false, errs.ErrQueryRowScan
	}

	return exists == "1", nil
}
