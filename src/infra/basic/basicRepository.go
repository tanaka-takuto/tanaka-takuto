package infra_basic

import (
	"app/domain/model/basic"
	"app/domain/model/util"
)

// 基本情報リポジトリ
type BasicRepositoryImpl struct {
	// ファイルパス
	path string
}

func (b BasicRepositoryImpl) LoadBasic() basic.Basic {
	return util.FromJson[basic.Basic](b.path)
}

// 基本情報リポジトリを作成する
func NewBasicRepository(path string) basic.BasicRepository {
	return BasicRepositoryImpl{
		path: path,
	}
}
