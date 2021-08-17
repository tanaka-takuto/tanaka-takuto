package infra_SkilInfo

import (
	"app/domain/model/skilInfo"
	"app/domain/model/util"
	"net/url"
)

// スキル情報リポジトリ
type SkilInfoRepositoryImpl struct {
	// ファイルパス
	path string
}

// スキル情報一覧
type skilInfos struct {
	// スキル情報一覧
	SkilInfos []skilInfo.SkilInfo
}

// スキル情報を読み込む
func (s SkilInfoRepositoryImpl) LoadSkilInfo() []skilInfo.SkilInfo {
	skilInfos := util.FromJson[skilInfos](s.path).SkilInfos

	for i, skilInfo := range skilInfos {
		skilInfos[i].Icon = url.PathEscape(skilInfo.Icon)
	}

	return skilInfos
}

// 基本情報リポジトリを作成する
func NewSkilInfoRepository(path string) skilInfo.SkilInfoRepository {
	return SkilInfoRepositoryImpl{
		path: path,
	}
}
