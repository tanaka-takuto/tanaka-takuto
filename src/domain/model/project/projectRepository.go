package project

// プロジェクトリポジトリ
type ProjectRepository interface {
	// プロジェクトをロードする
	LoadProject() []Project
}
