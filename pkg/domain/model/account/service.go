package account

import (
	"context"

	"github.com/game-core/gocrafter/internal/tokens"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/keys"
	"github.com/game-core/gocrafter/internal/times"
	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccount"
	shardService "github.com/game-core/gocrafter/pkg/domain/model/shard"
)

type AccountService interface {
	FindByUserId(ctx context.Context, userId string) (*userAccount.UserAccount, error)
	Create(ctx context.Context, tx *gorm.DB, req *AccountCreateRequest) (*AccountCreateResponse, error)
	Login(ctx context.Context, tx *gorm.DB, req *AccountLoginRequest) (*AccountLoginResponse, error)
	Check(ctx context.Context, req *AccountCheckRequest) (*AccountCheckResponse, error)
	GenerateUserID(ctx context.Context) (string, error)
}

type accountService struct {
	shardService          shardService.ShardService
	userAccountRepository userAccount.UserAccountRepository
}

func NewAccountService(
	shardService shardService.ShardService,
	userAccountRepository userAccount.UserAccountRepository,
) AccountService {
	return &accountService{
		shardService:          shardService,
		userAccountRepository: userAccountRepository,
	}
}

// FindByUserId ユーザーIDから取得する
func (s *accountService) FindByUserId(ctx context.Context, userId string) (*userAccount.UserAccount, error) {
	userAccount, err := s.userAccountRepository.Find(ctx, userId)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountRepository.Find", err)
	}

	return userAccount, err
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

	userAccount, err := s.userAccountRepository.Create(ctx, tx, userAccount.SetUserAccount(req.UserId, req.Name, hashPassword, times.Now(), times.Now()))
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountRepository.Create", err)
	}

	userAccount.Password = password

	return SetAccountCreateResponse(userAccount), nil
}

// Login ログインする
func (s *accountService) Login(ctx context.Context, tx *gorm.DB, req *AccountLoginRequest) (*AccountLoginResponse, error) {
	userAccount, err := s.userAccountRepository.Find(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountRepository.Find", err)
	}

	if !keys.CheckPassword(req.Password, userAccount.Password) {
		return nil, errors.NewError("invalid password")
	}

	userAccount.LoginAt = times.Now()
	result, err := s.userAccountRepository.Update(ctx, tx, userAccount)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountRepository.Update", err)
	}

	token, err := tokens.GenerateAuthTokenByUserId(userAccount.UserId, userAccount.Name)
	if err != nil {
		return nil, errors.NewMethodError("tokens.GenerateAuthTokenByUserId", err)
	}

	return SetAccountLoginResponse(token, result), nil
}

// Check ユーザーを確認する
func (s *accountService) Check(ctx context.Context, req *AccountCheckRequest) (*AccountCheckResponse, error) {
	userAccount, err := s.userAccountRepository.Find(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountRepository.Find", err)
	}

	return SetAccountCheckResponse(userAccount), err
}

// GenerateUserID ユーザーIDを生成する
func (s *accountService) GenerateUserID(ctx context.Context) (string, error) {
	shardKey, err := s.shardService.GetShardKeyAndUpdate(ctx, nil)
	if err != nil {
		return "", errors.NewMethodError("s.shardService.GetShardKeyAndUpdate", err)
	}

	userId, err := keys.GenerateUserID(shardKey)
	if err != nil {
		return "", errors.NewMethodError("keys.GenerateUserID", err)
	}

	return userId, nil
}
