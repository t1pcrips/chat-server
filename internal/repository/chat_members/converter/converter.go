package converter

import (
	"chat-server/internal/model"
	"chat-server/internal/repository/chat_members/member_model"
)

func ToMemberFromMemberService(member *model.CreateMemberRequest) *member_model.CreateMemberRequest {
	return &member_model.CreateMemberRequest{
		ChatId:    member.ChatId,
		Usernames: ToStringFromUser(member.Usernames),
	}
}

func ToStringFromUser(users []model.User) []string {
	usernames := make([]string, len(users))
	for i, username := range users {
		usernames[i] = username.Username
	}

	return usernames
}
