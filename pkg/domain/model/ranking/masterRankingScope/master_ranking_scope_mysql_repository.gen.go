// Package masterRankingScope ランキング範囲
//
//go:generate mockgen -source=./master_ranking_scope_mysql_repository.gen.go -destination=./master_ranking_scope_mysql_repository_mock.gen.go -package=masterRankingScope
package masterRankingScope

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/pkg/domain/enum"
)

type MasterRankingScopeMysqlRepository interface {
	Find(ctx context.Context, id int64) (*MasterRankingScope, error)
	FindOrNil(ctx context.Context, id int64) (*MasterRankingScope, error)
	FindByRankingScopeType(ctx context.Context, rankingScopeType enum.RankingScopeType) (*MasterRankingScope, error)
	FindOrNilByRankingScopeType(ctx context.Context, rankingScopeType enum.RankingScopeType) (*MasterRankingScope, error)
	FindList(ctx context.Context) (MasterRankingScopes, error)
	FindListByRankingScopeType(ctx context.Context, rankingScopeType enum.RankingScopeType) (MasterRankingScopes, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterRankingScope) (*MasterRankingScope, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterRankingScopes) (MasterRankingScopes, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterRankingScope) (*MasterRankingScope, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterRankingScope) error
}
