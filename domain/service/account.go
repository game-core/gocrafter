package service

import (
	request "github.com/game-core/gocrafter/api/presentation/request/account"
	response "github.com/game-core/gocrafter/api/presentation/response/account"
	entity "github.com/game-core/gocrafter/domain/entity/user/account"
	repository "github.com/game-core/gocrafter/domain/repository/user/account"
)

type AccountService interface {
	RegisterAccount(request *request.RegisterAccount) (*response.RegisterAccount, error)
}

type accountService struct {
	accountRepository repository.AccounRepository
}

func NewAccountService(
	accountRepository repository.AccounRepository,
) AccounService {
	return &accountService{
		accountRepository: accountRepository,
	}
}

// RegisterAccount アカウントを登録する
func (a *accountService) RegisterAccount(request *request.RegisterAccount) (*response.RegisterAccount, error) {
	return &response.RegisterAccount{

	}, nil
}
