package resume

// 経歴書リポジトリ
type ResumeRepository interface {
	// 経歴書を出力する
	Output(resume Resume) error
}
