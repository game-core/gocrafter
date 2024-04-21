// Package chat チャット送信リクエスト
package chat

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type ChatSendRequests []*ChatSendRequest

type ChatSendRequest struct {
	UserId        string
	ChatSpaceType enum.ChatSpaceType
	RoomId        string
	Content       string
}

func NewChatSendRequest() *ChatSendRequest {
	return &ChatSendRequest{}
}

func NewChatSendRequests() ChatSendRequests {
	return ChatSendRequests{}
}

func SetChatSendRequest(userId string, chatSpaceType enum.ChatSpaceType, roomId string, content string) *ChatSendRequest {
	return &ChatSendRequest{
		UserId:        userId,
		ChatSpaceType: chatSpaceType,
		RoomId:        roomId,
		Content:       content,
	}
}
