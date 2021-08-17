package skilType

// スキルタイプ
type SkilType string

// スキルタイプ一覧
const (
	// バックエンド
	Backend SkilType = "backend"

	// フロントエンド
	Frontend SkilType = "frontend"

	// 言語
	Language SkilType = "language"

	// フレームワーク
	Framework SkilType = "framework"

	// DB・データストア
	DB SkilType = "db"

	// プラットフォーム
	Platform SkilType = "platform"

	// ツール
	Tool SkilType = "tool"
)
