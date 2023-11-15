package account

import (
	"errors"
	"log"

	"github.com/game-core/gocrafter/config/key"
	"github.com/game-core/gocrafter/config/token"
	repository "github.com/game-core/gocrafter/domain/repository/user"
	request "github.com/game-core/gocrafter/api/presentation/request/account"
	response "github.com/game-core/gocrafter/api/presentation/response/account"
	accountEntity "github.com/game-core/gocrafter/domain/entity/user/account"
	accountRepository "github.com/game-core/gocrafter/domain/repository/user/account"
)

type AccountService interface {
	RegisterAccount(request *request.RegisterAccount) (*response.RegisterAccount, error)
	LoginAccount(request *request.LoginAccount) (*response.LoginAccount, error)
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

	uuid, err := key.GenerateUUID()
	if err != nil {
		return nil, err
	}

	password, hashedPassword, err := key.GeneratePassword()
	if err != nil {
		return nil, err
	}

	account := &accountEntity.Account{
		UUID:     uuid,
		Name:     req.Name,
		Password: hashedPassword,
	}

	ar, err := a.accountRepository.Create(account, tx)
	if err != nil {
		return nil, err
	}

	return &response.RegisterAccount{
		Status: 200,
		Item: response.Account{
			ID:       ar.ID,
			UUID:     ar.UUID,
			Name:     ar.Name,
			Password: password,
			Token:    "",
		},
	}, nil
}

// LoginAccount アカウントをログインする
func (a *accountService) LoginAccount(req *request.LoginAccount) (*response.LoginAccount, error) {
	ar, err := a.accountRepository.FindByUUID(req.UUID)
	if err != nil {
		return nil, err
	}

	if !key.CheckPassword(req.Password, ar.Password) {
		return nil,  errors.New("faild to key.CheckPassword")
	}

	token, err := token.GenerateAuthToken(ar.UUID, ar.Name)
	if err != nil {
		return nil,  errors.New("faild to token.GenerateAuthToken")
	}

	return &response.LoginAccount{
		Status: 200,
		Item: response.Account{
			ID:       ar.ID,
			UUID:     ar.UUID,
			Name:     ar.Name,
			Password: req.Password,
			Token:    token,
		},
	}, nil
}
