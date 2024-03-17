//go:generate mockgen -source=./action_service.go -destination=./action_service_mock.gen.go -package=action
package action

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterAction"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionRun"
	"github.com/game-core/gocrafter/pkg/domain/model/action/masterActionStep"
	"github.com/game-core/gocrafter/pkg/domain/model/action/userAction"
)

type ActionService interface {
	GetMaster(ctx context.Context) (*ActionGetMasterResponse, error)
	Check(ctx context.Context, now time.Time, req *ActionCheckRequest) error
	Run(ctx context.Context, tx *gorm.DB, now time.Time, req *ActionRunRequest) error
}

type actionService struct {
	masterActionRepository     masterAction.MasterActionRepository
	masterActionRunRepository  masterActionRun.MasterActionRunRepository
	masterActionStepRepository masterActionStep.MasterActionStepRepository
	userActionRepository       userAction.UserActionRepository
}

func NewActionService(
	masterActionRepository masterAction.MasterActionRepository,
	masterActionRunRepository masterActionRun.MasterActionRunRepository,
	masterActionStepRepository masterActionStep.MasterActionStepRepository,
	userActionRepository userAction.UserActionRepository,
) ActionService {
	return &actionService{
		masterActionRepository:     masterActionRepository,
		masterActionRunRepository:  masterActionRunRepository,
		masterActionStepRepository: masterActionStepRepository,
		userActionRepository:       userActionRepository,
	}
}

// GetMaster アクション一覧を取得する
func (s *actionService) GetMaster(ctx context.Context) (*ActionGetMasterResponse, error) {
	masterActionModels, err := s.masterActionRepository.FindList(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.masterActionRepository.FindList", err)
	}

	masterActionStepModels, err := s.masterActionStepRepository.FindList(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.masterActionStepRepository.FindList", err)
	}

	return SetActionGetMasterResponse(masterActionModels, masterActionStepModels), nil
}

// Check アクションが実行可能か確認する
func (s *actionService) Check(ctx context.Context, now time.Time, req *ActionCheckRequest) error {
	masterActionModel, err := s.masterActionRepository.FindByActionStepType(ctx, req.ActionStepType)
	if err != nil {
		return errors.NewMethodError("s.masterActionRepository.FindByActionStepType", err)
	}

	if _, err := s.checkTriggerAction(ctx, now, req.UserId, masterActionModel); err != nil {
		return errors.NewMethodError("s.checkTriggerAction", err)
	}

	return nil
}

// Run アクションを実行する
func (s *actionService) Run(ctx context.Context, tx *gorm.DB, now time.Time, req *ActionRunRequest) error {
	masterActionModel, err := s.masterActionRepository.FindByActionStepType(ctx, req.ActionStepType)
	if err != nil {
		return errors.NewMethodError("s.masterActionRepository.FindByActionStepType", err)
	}

	triggerUserActionModel, err := s.checkTriggerAction(ctx, now, req.UserId, masterActionModel)
	if err != nil {
		return errors.NewMethodError("s.checkTriggerAction", err)
	}

	if err := s.run(ctx, tx, now, req.UserId, masterActionModel, triggerUserActionModel); err != nil {
		return errors.NewMethodError("s.run", err)
	}

	return nil
}

// checkTriggerAction トリガーになるアクションを確認する
func (s *actionService) checkTriggerAction(ctx context.Context, now time.Time, userId string, masterActionModel *masterAction.MasterAction) (*userAction.UserAction, error) {
	if masterActionModel.TriggerActionId == nil {
		return nil, nil
	}

	triggerMasterActionModel, err := s.masterActionRepository.Find(ctx, *masterActionModel.TriggerActionId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterActionRepository.Find", err)
	}

	triggerUserActionModel, err := s.userActionRepository.Find(ctx, userId, triggerMasterActionModel.Id)
	if err != nil {
		return nil, errors.NewMethodError("s.userActionRepository.Find", err)
	}

	if triggerMasterActionModel.Expiration != nil && triggerUserActionModel.CheckExpiration(triggerMasterActionModel.Expiration, now) {
		return nil, errors.NewError("expiration date has expired")
	}

	return triggerUserActionModel, nil
}

// run 実行する
func (s *actionService) run(ctx context.Context, tx *gorm.DB, now time.Time, userId string, masterActionModel *masterAction.MasterAction, triggerUserActionModel *userAction.UserAction) error {
	if err := s.update(ctx, tx, userAction.SetUserAction(userId, masterActionModel.Name, masterActionModel.Id, now), triggerUserActionModel); err != nil {
		return errors.NewMethodError("s.update", err)
	}

	// 実行されるアクションがある場合は更新する
	masterActionRunModels, err := s.masterActionRunRepository.FindList(ctx)
	if err != nil {
		return errors.NewMethodError("s.masterActionRunRepository.FindList", err)
	}

	for _, model := range masterActionRunModels {
		if err := s.update(ctx, tx, userAction.SetUserAction(userId, model.Name, model.ActionId, now), nil); err != nil {
			return errors.NewMethodError("s.userActionRepository.Create", err)
		}
	}

	return nil
}

// update 更新する
func (s *actionService) update(ctx context.Context, tx *gorm.DB, model, trigger *userAction.UserAction) error {
	if trigger != nil {
		if err := s.userActionRepository.Delete(ctx, tx, trigger); err != nil {
			return errors.NewMethodError("s.userActionRepository.Delete", err)
		}
	}

	userActionModel, err := s.userActionRepository.FindOrNil(ctx, model.UserId, model.MasterActionId)
	if err != nil {
		return errors.NewMethodError("s.userActionRepository.FindOrNil", err)
	}

	if userActionModel != nil {
		if _, err := s.userActionRepository.Update(ctx, tx, model); err != nil {
			return errors.NewMethodError("s.userActionRepository.Update", err)
		}

		return nil
	}

	if _, err := s.userActionRepository.Create(ctx, tx, model); err != nil {
		return errors.NewMethodError("s.userActionRepository.Create", err)
	}

	return nil
}
