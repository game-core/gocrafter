// Package commonChat ルーム
package commonChat

import (
	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type CommonChats []*CommonChat

type CommonChat struct {
	ChatId        string
	UserId        string
	ChatSpaceType enum.ChatSpaceType
	RoomId        string
	Content       string
}

func NewCommonChat() *CommonChat {
	return &CommonChat{}
}

func NewCommonChats() CommonChats {
	return CommonChats{}
}

func SetCommonChat(chatId string, userId string, chatSpaceType enum.ChatSpaceType, roomId string, content string) *CommonChat {
	return &CommonChat{
		ChatId:        chatId,
		UserId:        userId,
		ChatSpaceType: chatSpaceType,
		RoomId:        roomId,
		Content:       content,
	}
}
