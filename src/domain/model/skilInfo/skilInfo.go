package skilInfo

import (
	"app/domain/model/util"
)

// スキル情報
type SkilInfo struct {
	// 名称
	Name string

	// 色
	Color util.Color

	// アイコン
	Icon string
}
