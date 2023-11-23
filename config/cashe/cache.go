package cashe

import (
	"fmt"
)

// CreateCacheKey キャッシュキーを生成する
func CreateCacheKey(table, method, key string) string {
	return fmt.Sprintf("%s:%s:%s", table, method, key)
}
