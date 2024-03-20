# 自動生成関連
自動生成に関連するドキュメントを記述
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
- コマンド
```
make docker_gen_api
```
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
- コマンド
```
make docker_gen_di
```
- 下記のようにusecaseやserviceに依存しているパッケージを引数にしたfuncを定義すると、wireで記述されるdiを自動生成してくれる
  - 記述の際の命名規則
    - メソッド名: New + パッケージ名 + レイヤー名をアッパーキャメルケースで記述する
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

### Domain
- コマンド
```
make docker_gen_domain
```
- `./docs/yaml/pkg/domain/model`配下にyamlファイルを定義する
  - Domainサービス毎にディレクトリを作成してその中にyamlを定義する
  - Infrastructure層のDaoとRepositoryを繋ぎこみたい場合は、さらにディレクトリ内にInfrastructureと同名のディレクトリとyamlを定義する(([例：ログインボーナス](https://github.com/game-core/gocrafter/tree/main/docs/yaml/pkg/domain/model/loginBonus)))
  - 繋ぎこみを行わないModelを定義したい場合はDomainサービスのディレクトリ直下に定義する
```yaml
name: MasterLoginBonus
package: masterLoginBonus
comment: "ログインボーナス"

structure:
  Id:
    name: id
    type: int64
    nullable: false
    number: 1
    comment: "ID"
  MasterLoginBonusEventId:
    name: master_login_bonus_event_id
    type: int64
    nullable: false
    number: 2
    comment: "ログボーナスイベントID"
  Name:
    name: name
    type: string
    nullable: false
    number: 3
    comment: "ログインボーナス名"
```
- Modelが生成される
```go
// Package masterLoginBonus ログインボーナス
package masterLoginBonus

type MasterLoginBonuses []*MasterLoginBonus

type MasterLoginBonus struct {
	Id                      int64
	MasterLoginBonusEventId int64
	Name                    string
}

func NewMasterLoginBonus() *MasterLoginBonus {
	return &MasterLoginBonus{}
}

func NewMasterLoginBonuses() MasterLoginBonuses {
	return MasterLoginBonuses{}
}

func SetMasterLoginBonus(id int64, masterLoginBonusEventId int64, name string) *MasterLoginBonus {
	return &MasterLoginBonus{
		Id:                      id,
		MasterLoginBonusEventId: masterLoginBonusEventId,
		Name:                    name,
	}
}
```
- Daoと繋ぎこみの設定を行っている場合、Infrastructure層の自動生成を行ったタイミングでRepositoryが生成される
  - CreatedAtとUpdatedAtはmodelでは除外されるため、コード内で日時系の値を参照・更新したい場合は別途追加する
  - CreatedAtとUpdatedAtはあくまでDB側でのみ取り扱うカラムとして定義する
```go
// Package masterLoginBonus ログインボーナス
//
//go:generate mockgen -source=./master_login_bonus_repository.gen.go -destination=./master_login_bonus_repository_mock.gen.go -package=masterLoginBonus
package masterLoginBonus

import (
	context "context"

	"gorm.io/gorm"
)

type MasterLoginBonusRepository interface {
	Find(ctx context.Context, id int64) (*MasterLoginBonus, error)
	FindOrNil(ctx context.Context, id int64) (*MasterLoginBonus, error)
	FindByMasterLoginBonusEventId(ctx context.Context, masterLoginBonusEventId int64) (*MasterLoginBonus, error)
	FindOrNilByMasterLoginBonusEventId(ctx context.Context, masterLoginBonusEventId int64) (*MasterLoginBonus, error)
	FindList(ctx context.Context) (MasterLoginBonuses, error)
	FindListByMasterLoginBonusEventId(ctx context.Context, masterLoginBonusEventId int64) (MasterLoginBonuses, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) (*MasterLoginBonus, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonuses) (MasterLoginBonuses, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) (*MasterLoginBonus, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) error
}
```

### Infrastructure
- コマンド
```
make docker_gen_infra
```
- `./docs/yaml/pkg/domain/infrastructure`配下にyamlファイルを定義する
  - 定義したフィールドやキーに応じたTable, Dao, SQLを生成してくれる
    - 現状、新規カラムを追加したくなった場合は対応しているSQLを削除して新規生成するか、AlterのSQLを自前で追加する必要がある
    - Dao, Tableに関しては上書き生成できる
```yaml
name: MasterLoginBonus
package: masterLoginBonus
comment: "ログインボーナス"

structure:
  Id:
    name: id
    type: int64
    nullable: false
    number: 1
    comment: "ID"
  MasterLoginBonusEventId:
    name: master_login_bonus_event_id
    type: int64
    nullable: false
    number: 2
    comment: "ログボーナスイベントID"
  Name:
    name: name
    type: string
    nullable: false
    number: 3
    comment: "ログインボーナス名"
primary:
  - Id
unique:
  - Id
index:
  - MasterLoginBonusEventId
```
- Tableが生成される
```go
// Package masterLoginBonus ログインボーナス
package masterLoginBonus

type MasterLoginBonuses []*MasterLoginBonus

type MasterLoginBonus struct {
	Id                      int64
	MasterLoginBonusEventId int64
	Name                    string
}

func NewMasterLoginBonus() *MasterLoginBonus {
	return &MasterLoginBonus{}
}

func NewMasterLoginBonuses() MasterLoginBonuses {
	return MasterLoginBonuses{}
}

func SetMasterLoginBonus(id int64, masterLoginBonusEventId int64, name string) *MasterLoginBonus {
	return &MasterLoginBonus{
		Id:                      id,
		MasterLoginBonusEventId: masterLoginBonusEventId,
		Name:                    name,
	}
}

func (t *MasterLoginBonus) TableName() string {
	return "master_login_bonus"
}
```
- Daoが生成される
```go
// Package masterLoginBonus ログインボーナス
package masterLoginBonus

import (
	"context"
	"fmt"

	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"

	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/pkg/domain/model/loginBonus/masterLoginBonus"
)

type masterLoginBonusDao struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
	Cache     *cache.Cache
}

func NewMasterLoginBonusDao(conn *database.SqlHandler) masterLoginBonus.MasterLoginBonusRepository {
	return &masterLoginBonusDao{
		ReadConn:  conn.Master.ReadConn,
		WriteConn: conn.Master.WriteConn,
		Cache:     cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (s *masterLoginBonusDao) Find(ctx context.Context, id int64) (*masterLoginBonus.MasterLoginBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "Find", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonus.MasterLoginBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonus()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "Find", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusDao) FindOrNil(ctx context.Context, id int64) (*masterLoginBonus.MasterLoginBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "FindOrNil", fmt.Sprintf("%d_", id)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonus.MasterLoginBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonus()
	res := s.ReadConn.WithContext(ctx).Where("id = ?", id).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "FindOrNil", fmt.Sprintf("%d_", id)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusDao) FindByMasterLoginBonusEventId(ctx context.Context, masterLoginBonusEventId int64) (*masterLoginBonus.MasterLoginBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "FindByMasterLoginBonusEventId", fmt.Sprintf("%d_", masterLoginBonusEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonus.MasterLoginBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonus()
	res := s.ReadConn.WithContext(ctx).Where("master_login_bonus_event_id = ?", masterLoginBonusEventId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.NewError("record does not exist")
	}

	m := masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "FindByMasterLoginBonusEventId", fmt.Sprintf("%d_", masterLoginBonusEventId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusDao) FindOrNilByMasterLoginBonusEventId(ctx context.Context, masterLoginBonusEventId int64) (*masterLoginBonus.MasterLoginBonus, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "FindOrNilByMasterLoginBonusEventId", fmt.Sprintf("%d_", masterLoginBonusEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(*masterLoginBonus.MasterLoginBonus); ok {
			return cachedEntity, nil
		}
	}

	t := NewMasterLoginBonus()
	res := s.ReadConn.WithContext(ctx).Where("master_login_bonus_event_id = ?", masterLoginBonusEventId).Find(t)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}

	m := masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name)
	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "FindOrNilByMasterLoginBonusEventId", fmt.Sprintf("%d_", masterLoginBonusEventId)), m, cache.DefaultExpiration)
	return m, nil
}

func (s *masterLoginBonusDao) FindList(ctx context.Context) (masterLoginBonus.MasterLoginBonuses, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "FindList", ""))
	if found {
		if cachedEntity, ok := cachedResult.(masterLoginBonus.MasterLoginBonuses); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterLoginBonuses()
	res := s.ReadConn.WithContext(ctx).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterLoginBonus.NewMasterLoginBonuses()
	for _, t := range ts {
		ms = append(ms, masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "FindList", ""), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterLoginBonusDao) FindListByMasterLoginBonusEventId(ctx context.Context, masterLoginBonusEventId int64) (masterLoginBonus.MasterLoginBonuses, error) {
	cachedResult, found := s.Cache.Get(cashes.CreateCacheKey("master_login_bonus", "FindListByMasterLoginBonusEventId", fmt.Sprintf("%d_", masterLoginBonusEventId)))
	if found {
		if cachedEntity, ok := cachedResult.(masterLoginBonus.MasterLoginBonuses); ok {
			return cachedEntity, nil
		}
	}

	ts := NewMasterLoginBonuses()
	res := s.ReadConn.WithContext(ctx).Where("master_login_bonus_event_id = ?", masterLoginBonusEventId).Find(&ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	ms := masterLoginBonus.NewMasterLoginBonuses()
	for _, t := range ts {
		ms = append(ms, masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name))
	}

	s.Cache.Set(cashes.CreateCacheKey("master_login_bonus", "FindListByMasterLoginBonusEventId", fmt.Sprintf("%d_", masterLoginBonusEventId)), ms, cache.DefaultExpiration)
	return ms, nil
}

func (s *masterLoginBonusDao) Create(ctx context.Context, tx *gorm.DB, m *masterLoginBonus.MasterLoginBonus) (*masterLoginBonus.MasterLoginBonus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterLoginBonus{
		Id:                      m.Id,
		MasterLoginBonusEventId: m.MasterLoginBonusEventId,
		Name:                    m.Name,
	}
	res := conn.Model(NewMasterLoginBonus()).WithContext(ctx).Create(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name), nil
}

func (s *masterLoginBonusDao) CreateList(ctx context.Context, tx *gorm.DB, ms masterLoginBonus.MasterLoginBonuses) (masterLoginBonus.MasterLoginBonuses, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	ts := NewMasterLoginBonuses()
	for _, m := range ms {
		t := &MasterLoginBonus{
			Id:                      m.Id,
			MasterLoginBonusEventId: m.MasterLoginBonusEventId,
			Name:                    m.Name,
		}
		ts = append(ts, t)
	}

	res := conn.Model(NewMasterLoginBonus()).WithContext(ctx).Create(ts)
	if err := res.Error; err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *masterLoginBonusDao) Update(ctx context.Context, tx *gorm.DB, m *masterLoginBonus.MasterLoginBonus) (*masterLoginBonus.MasterLoginBonus, error) {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	t := &MasterLoginBonus{
		Id:                      m.Id,
		MasterLoginBonusEventId: m.MasterLoginBonusEventId,
		Name:                    m.Name,
	}
	res := conn.Model(NewMasterLoginBonus()).WithContext(ctx).Where("id = ?", m.Id).Updates(t)
	if err := res.Error; err != nil {
		return nil, err
	}

	return masterLoginBonus.SetMasterLoginBonus(t.Id, t.MasterLoginBonusEventId, t.Name), nil
}

func (s *masterLoginBonusDao) Delete(ctx context.Context, tx *gorm.DB, m *masterLoginBonus.MasterLoginBonus) error {
	var conn *gorm.DB
	if tx != nil {
		conn = tx
	} else {
		conn = s.WriteConn
	}

	res := conn.Model(NewMasterLoginBonus()).WithContext(ctx).Where("id = ?", m.Id).Delete(NewMasterLoginBonus())
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
```
- SQlが生成される
```sql
CREATE TABLE master_login_bonus
(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT "ID",
    master_login_bonus_event_id BIGINT NOT NULL COMMENT "ログボーナスイベントID",
    name VARCHAR(255) NOT NULL COMMENT "ログインボーナス名",
    PRIMARY KEY(id),
    UNIQUE KEY(id),
    INDEX(master_login_bonus_event_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```
