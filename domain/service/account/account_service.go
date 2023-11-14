package account

import (
	"log"

	"github.com/game-core/gocrafter/config/key"
	repository "github.com/game-core/gocrafter/domain/repository/user"
	request "github.com/game-core/gocrafter/api/presentation/request/account"
	response "github.com/game-core/gocrafter/api/presentation/response/account"
	accountEntity "github.com/game-core/gocrafter/domain/entity/user/account"
	accountRepository "github.com/game-core/gocrafter/domain/repository/user/account"
)

type AccountService interface {
	RegisterAccount(request *request.RegisterAccount) (*response.RegisterAccount, error)
}

type accountService struct {
	transactionRepository repository.TransactionRepository
	accountRepository     accountRepository.AccountRepository
}

func NewAccountService(
	transactionRepository repository.TransactionRepository,
	accountRepository     accountRepository.AccountRepository,
) AccountService {
	return &accountService{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
	}
}

// RegisterAccount アカウントを登録する
func (a *accountService) RegisterAccount(req *request.RegisterAccount) (*response.RegisterAccount, error) {
	// transaction
	tx, err := a.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			if err := a.transactionRepository.Rollback(tx); err != nil {
				log.Panicln(err)
			}
		} else {
			if err := a.transactionRepository.Commit(tx); err != nil {
				log.Panicln(err)
			}
		}
	}()

	uuid, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	pass, err := key.GenerateKey()
	if err != nil {
		return nil, err
	}

	account := &accountEntity.Account{
		UUID:     uuid,
		Name:     req.Name,
		Password: pass,
	}

	ar, err := a.accountRepository.Create(account, tx)
	if err != nil {
		return nil, err
	}

	return &response.RegisterAccount{
		Status: 200,
		ID:       ar.ID,
		UUID:     ar.UUID,
		Name:     ar.Name,
		Password: ar.Password,
	}, nil
}
