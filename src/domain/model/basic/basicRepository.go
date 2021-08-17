package basic

// 基本情報リポジトリ
type BasicRepository interface {
	// 基本情報を読み込む
	LoadBasic() Basic
}
