// Package chat チャット送信レスポンス
package chat

import (
	"github.com/game-core/gocrafter/pkg/domain/model/chat/commonChat"
)

type ChatSendResponses []*ChatSendResponse

type ChatSendResponse struct {
	CommonChat *commonChat.CommonChat
}

func NewChatSendResponse() *ChatSendResponse {
	return &ChatSendResponse{}
}

func NewChatSendResponses() ChatSendResponses {
	return ChatSendResponses{}
}

func SetChatSendResponse(commonChat *commonChat.CommonChat) *ChatSendResponse {
	return &ChatSendResponse{
		CommonChat: commonChat,
	}
}
