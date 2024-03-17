//go:generate mockgen -source=./action_service.go -destination=./action_service_mock.gen.go -package=action
package action

import (
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterAction"
	"github.com/game-core/gocrafter/pkg/domain/model/action/userAction"
)

type ActionService interface {
}

type actionService struct {
	masterActionRepository masterAction.MasterActionRepository
	userActionRepository   userAction.UserActionRepository
}

func NewActionService(
	masterActionRepository masterAction.MasterActionRepository,
	userActionRepository userAction.UserActionRepository,
) ActionService {
	return &actionService{
		masterActionRepository: masterActionRepository,
		userActionRepository:   userActionRepository,
	}
}
