package experience

import "app/domain/model/skilType"

// 経験
type Experience struct {
	// 名称
	Name string

	// スキルタイプ一覧
	SkilTypes []skilType.SkilType
}
