package account

import (
	"context"
	"fmt"
	"github.com/game-core/gocrafter/internal"
	"time"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccount"
	"github.com/game-core/gocrafter/pkg/domain/model/shard"
)

type AccountService interface {
	Create(ctx context.Context, tx *gorm.DB, uam *AccountCreateRequest) (*AccountCreateResponse, error)
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

// Create アカウントを作成する
func (s *accountService) Create(ctx context.Context, tx *gorm.DB, req *AccountCreateRequest) (*AccountCreateResponse, error) {
	shardKey, err := s.shardService.GetShardKeyAndUpdate(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to s.shardService.GetShardKeyAndUpdate: %s", err)
	}

	userId, err := internal.GenerateUserID(shardKey)
	if err != nil {
		return nil, fmt.Errorf("failed to internal.GenerateUserID: %s", err)
	}

	password, err := internal.GeneratePassword()
	if err != nil {
		return nil, fmt.Errorf("failed to internal.GeneratePassword: %s", err)
	}

	hashPassword, err := internal.GenerateHashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to internal.GenerateHashPassword: %s", err)
	}

	userAccount, err := s.userAccountRepository.Create(ctx, tx, userAccount.SetUserAccount(userId, req.Name, hashPassword, time.Now(), time.Now()))
	if err != nil {
		return nil, fmt.Errorf("failed to s.userAccountRepository.Create: %s", err)
	}
	userAccount.Password = hashPassword

	return SetAccountCreateResponse(userAccount), nil
}
