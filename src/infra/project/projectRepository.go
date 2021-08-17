package infra_project

import (
	"app/domain/model/project"
	"app/domain/model/util"
)

// プロジェクトリポジトリ
type ProjectRepositoryImpl struct {
	// ファイルパス
	path string
}

// プロジェクト一覧
type projects struct {
	// プロジェクト一覧
	Projects []project.Project
}

func (p ProjectRepositoryImpl) LoadProject() []project.Project {
	return util.FromJson[projects](p.path).Projects
}

// プロジェクトリポジトリを作成する
func NewProjectRepository(path string) project.ProjectRepository {
	return ProjectRepositoryImpl{
		path: path,
	}
}
