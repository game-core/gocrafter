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
	Check(ctx context.Context, now time.Time, req *ActionCheckRequest) (*ActionCheckResponse, error)
	Run(ctx context.Context, tx *gorm.DB, now time.Time, req *ActionRunRequest) (*ActionRunResponse, error)
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
func (s *actionService) Check(ctx context.Context, now time.Time, req *ActionCheckRequest) (*ActionCheckResponse, error) {
	masterActionModel, err := s.masterActionRepository.FindByActionStepType(ctx, req.ActionStepType)
	if err != nil {
		return nil, errors.NewMethodError("s.masterActionRepository.FindByActionStepType", err)
	}

	masterActionStepModel, err := s.masterActionStepRepository.FindByActionStepType(ctx, req.ActionStepType)
	if err != nil {
		return nil, errors.NewMethodError("s.masterActionStepRepository.FindByActionStepType", err)
	}

	if err := s.checkTriggerAction(ctx, now, req.UserId, masterActionModel); err != nil {
		return nil, errors.NewMethodError("s.checkTriggerAction", err)
	}

	return SetActionCheckResponse(true, masterActionModel, masterActionStepModel), nil
}

// Run アクションを実行する
func (s *actionService) Run(ctx context.Context, tx *gorm.DB, now time.Time, req *ActionRunRequest) (*ActionRunResponse, error) {
	masterActionModel, err := s.masterActionRepository.FindByActionStepType(ctx, req.ActionStepType)
	if err != nil {
		return nil, errors.NewMethodError("s.masterActionRepository.FindByActionStepType", err)
	}

	masterActionStepModel, err := s.masterActionStepRepository.FindByActionStepType(ctx, req.ActionStepType)
	if err != nil {
		return nil, errors.NewMethodError("s.masterActionStepRepository.FindByActionStepType", err)
	}

	if err := s.checkTriggerAction(ctx, now, req.UserId, masterActionModel); err != nil {
		return nil, errors.NewMethodError("s.checkTriggerAction", err)
	}

	userActionModel, err := s.run(ctx, tx, now, req.UserId, masterActionModel)
	if err != nil {
		return nil, errors.NewMethodError("s.run", err)
	}

	return SetActionRunResponse(userActionModel, masterActionModel, masterActionStepModel), nil
}

// checkTriggerAction トリガーになるアクションを確認する
func (s *actionService) checkTriggerAction(ctx context.Context, now time.Time, userId string, masterActionModel *masterAction.MasterAction) error {
	if masterActionModel.TriggerActionId == nil {
		return nil
	}

	triggerMasterActionModel, err := s.masterActionRepository.Find(ctx, *masterActionModel.TriggerActionId)
	if err != nil {
		return errors.NewMethodError("s.masterActionRepository.Find", err)
	}

	triggerUserActionModel, err := s.userActionRepository.Find(ctx, userId, triggerMasterActionModel.Id)
	if err != nil {
		return errors.NewMethodError("s.userActionRepository.Find", err)
	}

	if triggerMasterActionModel.Expiration != nil && triggerUserActionModel.CheckExpiration(triggerMasterActionModel.Expiration, now) {
		return errors.NewError("expiration date has expired")
	}

	return nil
}

// run 実行する
func (s *actionService) run(ctx context.Context, tx *gorm.DB, now time.Time, userId string, masterActionModel *masterAction.MasterAction) (*userAction.UserAction, error) {
	userActionModel, err := s.update(ctx, tx, userAction.SetUserAction(userId, masterActionModel.Name, masterActionModel.Id, now))
	if err != nil {
		return nil, errors.NewMethodError("s.update", err)
	}

	// 実行されるアクションがある場合は更新する
	masterActionRunModels, err := s.masterActionRunRepository.FindList(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.masterActionRunRepository.FindList", err)
	}

	for _, model := range masterActionRunModels {
		if _, err := s.update(ctx, tx, userAction.SetUserAction(userId, model.Name, model.ActionId, now)); err != nil {
			return nil, errors.NewMethodError("s.userActionRepository.Create", err)
		}
	}

	return userActionModel, nil
}

// update 更新する
func (s *actionService) update(ctx context.Context, tx *gorm.DB, model *userAction.UserAction) (*userAction.UserAction, error) {
	userActionModel, err := s.userActionRepository.FindOrNil(ctx, model.UserId, model.MasterActionId)
	if err != nil {
		return nil, errors.NewMethodError("s.userActionRepository.FindOrNil", err)
	}

	if userActionModel != nil {
		result, err := s.userActionRepository.Update(ctx, tx, model)
		if err != nil {
			return nil, errors.NewMethodError("s.userActionRepository.Update", err)
		}

		return result, nil
	}

	result, err := s.userActionRepository.Create(ctx, tx, model)
	if err != nil {
		return nil, errors.NewMethodError("s.userActionRepository.Create", err)
	}

	return result, nil
}
