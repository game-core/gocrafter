//go:generate mockgen -source=./account_service.go -destination=./account_service_mock.gen.go -package=account
package account

import (
	"errors"
	"log"

	"github.com/game-core/gocrafter/config/key"
	"github.com/game-core/gocrafter/config/token"

	request "github.com/game-core/gocrafter/auth/presentation/request/account"
	response "github.com/game-core/gocrafter/auth/presentation/response/account"
	accountEntity "github.com/game-core/gocrafter/domain/entity/auth/account"
	accountRepository "github.com/game-core/gocrafter/domain/repository/auth/account"
	transactionRepository "github.com/game-core/gocrafter/domain/repository/auth/transaction"
)

type AccountService interface {
	RegisterAccount(req *request.RegisterAccount) (*response.RegisterAccount, error)
	LoginAccount(req *request.LoginAccount) (*response.LoginAccount, error)
	CheckAccount(req *request.CheckAccount) (*response.CheckAccount, error)
}

type accountService struct {
	transactionRepository transactionRepository.TransactionRepository
	accountRepository     accountRepository.AccountRepository
}

func NewAccountService(
	transactionRepository transactionRepository.TransactionRepository,
	accountRepository accountRepository.AccountRepository,
) AccountService {
	return &accountService{
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
	}
}

// RegisterAccount アカウントを登録する
func (s *accountService) RegisterAccount(req *request.RegisterAccount) (*response.RegisterAccount, error) {
	// transaction
	tx, err := s.transactionRepository.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := s.transactionRepository.CommitOrRollback(tx, err); err != nil {
			log.Panicln(err)
		}
	}()

	a, err := s.accountRepository.FindOrNilByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if a != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := key.GenerateHashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	account := &accountEntity.Account{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	a, err = s.accountRepository.Create(account, tx)
	if err != nil {
		return nil, err
	}

	return response.ToRegisterAccount(200, *response.ToAccount(a.ID, a.Name, a.Email, req.Password, "")), nil
}

// LoginAccount アカウントをログインする
func (s *accountService) LoginAccount(req *request.LoginAccount) (*response.LoginAccount, error) {
	a, err := s.accountRepository.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if !key.CheckPassword(req.Password, a.Password) {
		return nil, errors.New("faild to key.CheckPassword")
	}

	token, err := token.GenerateAuthTokenByEmail(a.Email, a.Name)
	if err != nil {
		return nil, errors.New("faild to token.GenerateAuthToken")
	}

	return response.ToLoginAccount(200, *response.ToAccount(a.ID, a.Name, a.Email, req.Password, token)), nil
}

// CheckAccount アカウントを確認する
func (s *accountService) CheckAccount(req *request.CheckAccount) (*response.CheckAccount, error) {
	a, err := s.accountRepository.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	return response.ToCheckAccount(200, *response.ToAccount(a.ID, a.Name, a.Email, "", "")), nil
}
