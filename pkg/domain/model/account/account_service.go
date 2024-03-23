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
	shardService          shard.ShardService
	userAccountRepository userAccount.UserAccountRepository
}

func NewAccountService(
	shardService shard.ShardService,
	userAccountRepository userAccount.UserAccountRepository,
) AccountService {
	return &accountService{
		shardService:          shardService,
		userAccountRepository: userAccountRepository,
	}
}

// FindByUserId ユーザーIDから取得する
func (s *accountService) FindByUserId(ctx context.Context, userId string) (*userAccount.UserAccount, error) {
	userAccountModel, err := s.userAccountRepository.Find(ctx, userId)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountRepository.Find", err)
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

	userAccountModel, err := s.userAccountRepository.Create(ctx, tx, userAccount.SetUserAccount(req.UserId, req.Name, hashPassword, times.Now(), times.Now()))
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountRepository.Create", err)
	}

	userAccountModel.Password = password

	return SetAccountCreateResponse(userAccountModel), nil
}

// Login ログインする
func (s *accountService) Login(ctx context.Context, tx *gorm.DB, req *AccountLoginRequest) (*AccountLoginResponse, error) {
	userAccountModel, err := s.userAccountRepository.Find(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountRepository.Find", err)
	}

	if !keys.CheckPassword(req.Password, userAccountModel.Password) {
		return nil, errors.NewError("invalid password")
	}

	userAccountModel.LoginAt = times.Now()
	result, err := s.userAccountRepository.Update(ctx, tx, userAccountModel)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountRepository.Update", err)
	}

	token, err := tokens.GenerateAuthTokenByUserId(userAccountModel.UserId, userAccountModel.Name)
	if err != nil {
		return nil, errors.NewMethodError("tokens.GenerateAuthTokenByUserId", err)
	}

	return SetAccountLoginResponse(token, result), nil
}

// Check ユーザーを確認する
func (s *accountService) Check(ctx context.Context, req *AccountCheckRequest) (*AccountCheckResponse, error) {
	userAccountModel, err := s.userAccountRepository.Find(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountRepository.Find", err)
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
