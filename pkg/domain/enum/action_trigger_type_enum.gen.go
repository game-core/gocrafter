// Package enum アクショントリガータイプ
package enum

type ActionTriggerType int32

const (
	ActionTriggerType_Continuation    ActionTriggerType = 0
	ActionTriggerType_Discontinuation ActionTriggerType = 1
)
