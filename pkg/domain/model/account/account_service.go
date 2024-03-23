//go:generate mockgen -source=./account_service.go -destination=./account_service_mock.gen.go -package=account
package account

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/keys"
	"github.com/game-core/gocrafter/internal/times"
	"github.com/game-core/gocrafter/internal/tokens"
	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccount"
	"github.com/game-core/gocrafter/pkg/domain/model/shard"
)

type AccountService interface {
	FindByUserId(ctx context.Context, userId string) (*userAccount.UserAccount, error)
	Create(ctx context.Context, tx *gorm.DB, req *AccountCreateRequest) (*AccountCreateResponse, error)
	Login(ctx context.Context, tx *gorm.DB, req *AccountLoginRequest) (*AccountLoginResponse, error)
	Check(ctx context.Context, req *AccountCheckRequest) (*AccountCheckResponse, error)
	GenerateUserID(ctx context.Context) (string, error)
}

type accountService struct {
	shardService               shard.ShardService
	userAccountMysqlRepository userAccount.UserAccountMysqlRepository
}

func NewAccountService(
	shardService shard.ShardService,
	userAccountMysqlRepository userAccount.UserAccountMysqlRepository,
) AccountService {
	return &accountService{
		shardService:               shardService,
		userAccountMysqlRepository: userAccountMysqlRepository,
	}
}

// FindByUserId ユーザーIDから取得する
func (s *accountService) FindByUserId(ctx context.Context, userId string) (*userAccount.UserAccount, error) {
	userAccountModel, err := s.userAccountMysqlRepository.Find(ctx, userId)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountMysqlRepository.Find", err)
	}

	return userAccountModel, err
}

// Create アカウントを作成する
func (s *accountService) Create(ctx context.Context, tx *gorm.DB, req *AccountCreateRequest) (*AccountCreateResponse, error) {
	password, err := keys.GeneratePassword()
	if err != nil {
		return nil, errors.NewMethodError("keys.GeneratePassword", err)
	}

	hashPassword, err := keys.GenerateHashPassword(password)
	if err != nil {
		return nil, errors.NewMethodError("keys.GenerateHashPassword", err)
	}

	userAccountModel, err := s.userAccountMysqlRepository.Create(ctx, tx, userAccount.SetUserAccount(req.UserId, req.Name, hashPassword, times.Now(), times.Now()))
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountMysqlRepository.Create", err)
	}

	userAccountModel.Password = password

	return SetAccountCreateResponse(userAccountModel), nil
}

// Login ログインする
func (s *accountService) Login(ctx context.Context, tx *gorm.DB, req *AccountLoginRequest) (*AccountLoginResponse, error) {
	userAccountModel, err := s.userAccountMysqlRepository.Find(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountMysqlRepository.Find", err)
	}

	if !keys.CheckPassword(req.Password, userAccountModel.Password) {
		return nil, errors.NewError("invalid password")
	}

	userAccountModel.LoginAt = times.Now()
	result, err := s.userAccountMysqlRepository.Update(ctx, tx, userAccountModel)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountMysqlRepository.Update", err)
	}

	token, err := tokens.GenerateAuthTokenByUserId(userAccountModel.UserId, userAccountModel.Name)
	if err != nil {
		return nil, errors.NewMethodError("tokens.GenerateAuthTokenByUserId", err)
	}

	return SetAccountLoginResponse(token, result), nil
}

// Check ユーザーを確認する
func (s *accountService) Check(ctx context.Context, req *AccountCheckRequest) (*AccountCheckResponse, error) {
	userAccountModel, err := s.userAccountMysqlRepository.Find(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountMysqlRepository.Find", err)
	}

	return SetAccountCheckResponse(userAccountModel), err
}

// GenerateUserID ユーザーIDを生成する
func (s *accountService) GenerateUserID(ctx context.Context) (string, error) {
	shardKey, err := s.shardService.GetShardKey(ctx)
	if err != nil {
		return "", errors.NewMethodError("s.shardService.GetShardKey", err)
	}

	userId, err := keys.GenerateUserID(shardKey)
	if err != nil {
		return "", errors.NewMethodError("keys.GenerateUserID", err)
	}

	return userId, nil
}
