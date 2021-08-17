package skilInfo

// スキル情報リポジトリ
type SkilInfoRepository interface {
	// スキル情報を読み込む
	LoadSkilInfo() []SkilInfo
}
