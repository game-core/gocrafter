# 自動生成関連

## 自動生成可能な範囲
- API
  - request
  - response
  - handler
  - di


- Domain
  - model
  - repository
  - enum


- Infrastructure
  - dao
  - table
  - sql


## API
### request, responseの生成
  - `./docs/yaml/api`配下にyamlファイルを定義する
  

- 定義するyamlファイル
```yaml
name: LoginBonusReceiveRequest
package: loginBonus
comment: "ログインボーナス受け取りリクエスト"

structure:
  UserId:
    name: user_id
    type: string
    nullable: false
    number: 1
    comment: "ユーザーID"
  MasterLoginBonusId:
    name: master_login_bonus_id
    type: int64
    nullable: false
    number: 2
    comment: "ログインボーナスID"
```
- protobufとgoのSetterが生成される
```protobuf
// ログインボーナス受け取りリクエスト
syntax = "proto3";

package proto;

option go_package = "api/game/presentation/server/loginBonus";

message LoginBonusReceiveRequest {
  string user_id = 1;
  int64 master_login_bonus_id = 2;
}
```

```go
// Package loginBonus ログインボーナス受け取りリクエスト
package loginBonus

func SetLoginBonusReceiveRequest(userId string, masterLoginBonusId int64) *LoginBonusReceiveRequest {
	return &LoginBonusReceiveRequest{
		UserId:             userId,
		MasterLoginBonusId: masterLoginBonusId,
	}
}
```

### handlerの生成
- `./docs/yaml/api`配下にyamlファイルを定義する

- 定義するyamlファイル
  - 複数ある場合は続けて記述する
```yaml
name: LoginBonus
package: loginBonus
comment: "ログインボーナス"

structure:
  Receive:
    name: receive
    method: unary
    auth: false
    request: LoginBonusReceiveRequest
    response: LoginBonusReceiveResponse
    number: 1
    comment: "ログインボーナスを受け取る"

```
- protobufとgoのhandlerが生成される
```protobuf
// ログインボーナス
syntax = "proto3";

package proto;

option go_package = "api/game/presentation/server/loginBonus";

import "login_bonus_receive_request.proto";
import "login_bonus_receive_response.proto";

service LoginBonus {
  rpc Receive (LoginBonusReceiveRequest) returns (LoginBonusReceiveResponse);
}
```

```go
package loginBonus

import (
	"context"

	"github.com/game-core/gocrafter/api/game/presentation/server/loginBonus"
	loginBonusUsecase "github.com/game-core/gocrafter/api/game/usecase/loginBonus"
	"github.com/game-core/gocrafter/internal/errors"
)

type LoginBonusHandler interface {
	loginBonus.LoginBonusServer
}

type loginBonusHandler struct {
	loginBonus.UnimplementedLoginBonusServer
	loginBonusUsecase loginBonusUsecase.LoginBonusUsecase
}

func NewLoginBonusHandler(
	loginBonusUsecase loginBonusUsecase.LoginBonusUsecase,
) LoginBonusHandler {
	return &loginBonusHandler{
		loginBonusUsecase: loginBonusUsecase,
	}
}

// Receive ログインボーナスを受け取る
func (s *loginBonusHandler) Receive(ctx context.Context, req *loginBonus.LoginBonusReceiveRequest) (*loginBonus.LoginBonusReceiveResponse, error) {
	res, err := s.loginBonusUsecase.Receive(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.loginBonusUsecase.Receive", err)
	}

	return res, nil
}
```

### diの生成
- 下記のようにusecaseやserviceに依存しているパッケージを引数にしたfuncを定義すると、wireで記述されるdiを自動生成してくれる
  - 記述の際の命名規則
    - handler: 自動生成されるため不要
    - usecase
      - usecaseからservice依存のみ許可する
      - usecase同士の依存は不可
      - パッケージ名 + Serviceをキャメルケースで記述する
    - service
      - serviceからrepository、またはservice同士の依存を許可する
      - service同士依存は依存関係の循環が発生しないように注意する
      - パッケージ名 + Repositoryをキャメルケースで記述する
      - パッケージ名 + Serviceをキャメルケースで記述する
```go
func NewLoginBonusService(
	itemService item.ItemService,
	userLoginBonusRepository userLoginBonus.UserLoginBonusRepository,
	masterLoginBonusRepository masterLoginBonus.MasterLoginBonusRepository,
	masterLoginBonusEventRepository masterLoginBonusEvent.MasterLoginBonusEventRepository,
	masterLoginBonusItemRepository masterLoginBonusItem.MasterLoginBonusItemRepository,
	masterLoginBonusScheduleRepository masterLoginBonusSchedule.MasterLoginBonusScheduleRepository,
) LoginBonusService {
	return &loginBonusService{
		itemService:                        itemService,
		userLoginBonusRepository:           userLoginBonusRepository,
		masterLoginBonusRepository:         masterLoginBonusRepository,
		masterLoginBonusEventRepository:    masterLoginBonusEventRepository,
		masterLoginBonusItemRepository:     masterLoginBonusItemRepository,
		masterLoginBonusScheduleRepository: masterLoginBonusScheduleRepository,
	}
}
```
- 生成されるwireコード
```go
func InitializeLoginBonusService() loginBonusService.LoginBonusService {
	wire.Build(
		database.NewDB,
		loginBonusService.NewLoginBonusService,
		InitializeItemService,
		userLoginBonusDao.NewUserLoginBonusDao,
		masterLoginBonusDao.NewMasterLoginBonusDao,
		masterLoginBonusEventDao.NewMasterLoginBonusEventDao,
		masterLoginBonusItemDao.NewMasterLoginBonusItemDao,
		masterLoginBonusScheduleDao.NewMasterLoginBonusScheduleDao,
	)
	return nil
}
```
