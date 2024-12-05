package converter

import (
	"github.com/t1pcrips/chat-service/internal/model"
	"github.com/t1pcrips/chat-service/internal/repository/chat_members/member_model"
)

func ToMemberFromMemberService(member *model.CreateMemberRequest) *member_model.CreateMemberRequest {
	return &member_model.CreateMemberRequest{
		ChatId:    member.ChatId,
		Usernames: member.Usernames,
	}
}
