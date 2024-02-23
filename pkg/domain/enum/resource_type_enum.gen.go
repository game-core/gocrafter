// Package enum リソースタイプ
package enum

type ResourceType int32

const (
	ResourceType_Normal  ResourceType = 0
	ResourceType_Card    ResourceType = 1
	ResourceType_Ticket  ResourceType = 2
	ResourceType_Coin    ResourceType = 3
	ResourceType_Crystal ResourceType = 4
)
