package member_model

type CreateMemberRequest struct {
	ChatId    int64    `db:"chat_id"`
	Usernames []string `db:"usernames"`
}

type Username struct {
	Username string `db:"username"`
}
