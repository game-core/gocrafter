//go:generate mockgen -source=./loginReward_service.go -destination=./loginReward_service_mock.gen.go -package=loginReward
package loginReward

import (
	"log"
	"time"

	request "github.com/game-core/gocrafter/api/presentation/request/loginReward"
	response "github.com/game-core/gocrafter/api/presentation/response/loginReward"
	userLoginRewardEntity "github.com/game-core/gocrafter/domain/entity/user/loginReward"
	configRepository "github.com/game-core/gocrafter/domain/repository/config"
	eventRepository "github.com/game-core/gocrafter/domain/repository/master/event"
	masterLoginRewardRepository "github.com/game-core/gocrafter/domain/repository/master/loginReward"
	userLoginRewardRepository "github.com/game-core/gocrafter/domain/repository/user/loginReward"
)

type LoginRewardService interface {
	ReceiveLoginReward(req request.ReceiveLoginReward, now time.Time) (*response.ReceiveLoginReward, error)
}

type loginRewardService struct {
	transactionRepository       configRepository.TransactionRepository
	loginRewardStatusRepository userLoginRewardRepository.LoginRewardStatusRepository
	loginRewardModelRepository  masterLoginRewardRepository.LoginRewardModelRepository
	eventRepository             eventRepository.EventRepository
}

func NewLoginRewardService(
	transactionRepository configRepository.TransactionRepository,
	loginRewardStatusRepository userLoginRewardRepository.LoginRewardStatusRepository,
	loginRewardModelRepository masterLoginRewardRepository.LoginRewardModelRepository,
	eventRepository eventRepository.EventRepository,
) LoginRewardService {
	return &loginRewardService{
		transactionRepository:       transactionRepository,
		loginRewardStatusRepository: loginRewardStatusRepository,
		loginRewardModelRepository:  loginRewardModelRepository,
		eventRepository:             eventRepository,
	}
}

// ReceiveLoginReward 受け取る
func (s *loginRewardService) ReceiveLoginReward(req request.ReceiveLoginReward, now time.Time) (*response.ReceiveLoginReward, error) {
	// transaction
	tx, err := s.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			if err := s.transactionRepository.Rollback(tx); err != nil {
				log.Panicln(err)
			}
		} else {
			if err := s.transactionRepository.Commit(tx); err != nil {
				log.Panicln(err)
			}
		}
	}()

	// ログイン報酬モデルを取得
	lrmr, err := s.loginRewardModelRepository.FindByID(req.LoginRewardModelID)
	if err != nil {
		return nil, err
	}

	// イベントを取得
	er, err := s.eventRepository.FindByID(lrmr.EventID)
	if err != nil {
		return nil, err
	}

	// ログイン報酬ステータスを取得
	lrsr, err := s.loginRewardStatusRepository.FindOrNilByLoginRewardModelID(lrmr.ID, req.ShardKey)
	if err != nil {
		return nil, err
	}

	// 初回ログインの場合は追加
	if lrsr == nil {
		entity := &userLoginRewardEntity.LoginRewardStatus{
			ShardKey:           req.ShardKey,
			AccountID:          req.AccountID,
			LoginRewardModelID: lrmr.ID,
			LastReceivedAt:     now,
		}
		_, err := s.loginRewardStatusRepository.Create(entity, req.ShardKey, tx)
		if err != nil {
			return nil, err
		}
	}

	return &response.ReceiveLoginReward{
		Status: 200,
		Item: response.LoginRewardStatus{
			LoginRewardModel: response.LoginRewardModel{
				ID:   lrmr.ID,
				Name: lrmr.Name,
				Event: response.Event{
					ID:      er.ID,
					Name:    er.Name,
					Repeat:  er.Repeat,
					StartAt: er.StartAt,
					EndAt:   er.EndAt,
				},
			},
			LastReceivedAt: lrsr.LastReceivedAt,
		},
	}, nil
}
