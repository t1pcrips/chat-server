package member_model

type CreateMemberRequest struct {
	ChatId    int64
	Usernames []string
}

type Username struct {
	Username string
}
