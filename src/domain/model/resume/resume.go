package resume

import (
	"app/domain/model/basic"
	"app/domain/model/project"
	"app/domain/model/skilInfo"
	skilset "app/domain/model/skilSet"
)

// 経歴書
type Resume struct {
	// エイリアス
	Aliases []skilInfo.SkilInfo

	// 基本情報
	Basic basic.Basic

	// スキルセット
	SkilSet skilset.SkilSet

	// プロジェクト
	Projects []project.Project
}
