package account

import (
	"context"
	"fmt"
	"time"

	"github.com/game-core/gocrafter/internal/keys"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/pkg/domain/model/account/userAccount"
	shardService "github.com/game-core/gocrafter/pkg/domain/model/shard"
)

type AccountService interface {
	Create(ctx context.Context, tx *gorm.DB, uam *AccountCreateRequest) (*AccountCreateResponse, error)
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

// Create アカウントを作成する
func (s *accountService) Create(ctx context.Context, tx *gorm.DB, req *AccountCreateRequest) (*AccountCreateResponse, error) {
	password, err := keys.GeneratePassword()
	if err != nil {
		return nil, fmt.Errorf("failed to internal.GeneratePassword: %s", err)
	}

	hashPassword, err := keys.GenerateHashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to internal.GenerateHashPassword: %s", err)
	}

	userAccount, err := s.userAccountRepository.Create(ctx, tx, userAccount.SetUserAccount(req.UserId, req.Name, hashPassword, time.Now(), time.Now()))
	if err != nil {
		return nil, fmt.Errorf("failed to s.userAccountRepository.Create: %s", err)
	}
	userAccount.Password = hashPassword

	return SetAccountCreateResponse(userAccount), nil
}
