//go:generate mockgen -source=./account_service.go -destination=./account_service_mock.gen.go -package=account
package account

import (
	"errors"
	"log"

	request "github.com/game-core/gocrafter/api/presentation/request/account"
	response "github.com/game-core/gocrafter/api/presentation/response/account"
	"github.com/game-core/gocrafter/config/key"
	"github.com/game-core/gocrafter/config/token"
	accountEntity "github.com/game-core/gocrafter/domain/entity/user/account"
	userRepository "github.com/game-core/gocrafter/domain/repository/user"
	accountRepository "github.com/game-core/gocrafter/domain/repository/user/account"
	shardService "github.com/game-core/gocrafter/domain/service/api/shard"
)

type AccountService interface {
	RegisterAccount(req *request.RegisterAccount) (*response.RegisterAccount, error)
	LoginAccount(req *request.LoginAccount) (*response.LoginAccount, error)
	CheckAccount(req *request.CheckAccount) (*response.CheckAccount, error)
}

type accountService struct {
	shardService          shardService.ShardService
	transactionRepository userRepository.TransactionRepository
	accountRepository     accountRepository.AccountRepository
}

func NewAccountService(
	shardService shardService.ShardService,
	transactionRepository userRepository.TransactionRepository,
	accountRepository accountRepository.AccountRepository,
) AccountService {
	return &accountService{
		shardService:          shardService,
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
	}
}

// RegisterAccount アカウントを登録する
func (s *accountService) RegisterAccount(req *request.RegisterAccount) (*response.RegisterAccount, error) {
	// シャードキーを取得
	ss, err := s.shardService.GetShard()
	if err != nil {
		return nil, err
	}

	// transaction
	tx, err := s.transactionRepository.Begin(ss.NextShardKey)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := s.transactionRepository.CommitOrRollback(tx, err); err != nil {
			log.Panicln(err)
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
		ShardKey: ss.NextShardKey,
		Name:     req.Name,
		Password: hashedPassword,
	}

	ar, err := s.accountRepository.Create(account, ss.NextShardKey, tx)
	if err != nil {
		return nil, err
	}

	return response.ToRegisterAccount(200, *response.ToAccount(ar.ID, ar.ShardKey, ar.UUID, ar.Name, password, "")), nil
}

// LoginAccount アカウントをログインする
func (s *accountService) LoginAccount(req *request.LoginAccount) (*response.LoginAccount, error) {
	ar, err := s.accountRepository.FindByUUID(req.UUID, req.ShardKey)
	if err != nil {
		return nil, err
	}

	if !key.CheckPassword(req.Password, ar.Password) {
		return nil, errors.New("faild to key.CheckPassword")
	}

	token, err := token.GenerateAuthTokenByUUID(ar.UUID, ar.Name)
	if err != nil {
		return nil, errors.New("faild to token.GenerateAuthToken")
	}

	return response.ToLoginAccount(200, *response.ToAccount(ar.ID, ar.ShardKey, ar.UUID, ar.Name, req.Password, token)), nil
}

// CheckAccount アカウントを確認する
func (s *accountService) CheckAccount(req *request.CheckAccount) (*response.CheckAccount, error) {
	ar, err := s.accountRepository.FindByUUID(req.UUID, req.ShardKey)
	if err != nil {
		return nil, err
	}

	return response.ToCheckAccount(200, *response.ToAccount(ar.ID, ar.ShardKey, ar.UUID, ar.Name, "", "")), nil
}
