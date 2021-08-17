package project

import (
	"app/domain/model/experience"
	"app/domain/model/util"
)

// 業務
type Operation = string

// プロジェクト
type Project struct {
	// タイトル
	Title string

	// 開始年月
	From util.Date

	// 終了年月
	To util.Date

	// 詳細
	Description string

	// 経験一覧
	Experiences []experience.Experience

	// 業務一覧
	Operations []Operation
}
