//go:generate mockgen -source=./loginReward_service.go -destination=./loginReward_service_mock.gen.go -package=loginReward
package loginReward

import (
	request "github.com/game-core/gocrafter/api/presentation/request/loginReward"
	response "github.com/game-core/gocrafter/api/presentation/response/loginReward"
	configRepository "github.com/game-core/gocrafter/domain/repository/config"
	masterLoginRewardRepository "github.com/game-core/gocrafter/domain/repository/master/loginReward"
	userLoginRewardRepository "github.com/game-core/gocrafter/domain/repository/user/loginReward"
)

type LoginRewardService interface {
	ReceiveLoginReward(req request.ReceiveLoginReward) (*response.ReceiveLoginReward, error)
}

type loginRewardService struct {
	transactionRepository       configRepository.TransactionRepository
	loginRewardStatusRepository userLoginRewardRepository.LoginRewardStatusRepository
	loginRewardModelRepository  masterLoginRewardRepository.LoginRewardModelRepository
}

func NewLoginRewardService(
	transactionRepository configRepository.TransactionRepository,
	loginRewardStatusRepository userLoginRewardRepository.LoginRewardStatusRepository,
	loginRewardModelRepository masterLoginRewardRepository.LoginRewardModelRepository,
) LoginRewardService {
	return &loginRewardService{
		transactionRepository:       transactionRepository,
		loginRewardStatusRepository: loginRewardStatusRepository,
		loginRewardModelRepository:  loginRewardModelRepository,
	}
}

// ReceiveLoginReward 受け取る
func (s *loginRewardService) ReceiveLoginReward(req request.ReceiveLoginReward) (*response.ReceiveLoginReward, error) {
	return &response.ReceiveLoginReward{}, nil
}
