package infra_resume

import (
	"app/domain/model/experience"
	"app/domain/model/project"
	"app/domain/model/resume"
	"app/domain/model/skilType"
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"text/template"
)

// 経歴書リポジトリ
type ResumeRepositoryImpl struct {
	// 経歴書出力先
	path string

	// Json出力先
	jsonPath string
}

// 経歴書
type ResumeTmp struct {
	resume.Resume

	// プロジェクト一覧
	Projects []ProjectTmp
}

// プロジェクト
type ProjectTmp struct {
	project.Project

	// フロントエンド
	Frontend []experience.Experience

	// バックエンド
	Backend []experience.Experience

	// その他
	Other []experience.Experience
}

// 経歴書リポジトリを作成する
func NewResumeRepository(path string) resume.ResumeRepository {
	return ResumeRepositoryImpl{
		path:     path,
		jsonPath: "resume.json",
	}

}

func (r ResumeRepositoryImpl) Output(resume resume.Resume) error {

	resumeTmp := createResume(resume)

	outputResume(r.path, resumeTmp)

	outputJson(r.jsonPath, resumeTmp)

	return nil
}

// 出力用の経歴書オブジェクトを作成する
func createResume(original resume.Resume) ResumeTmp {
	projects := []ProjectTmp{}

	for _, currentProject := range original.Projects {
		project := ProjectTmp{
			Project:  currentProject,
			Frontend: []experience.Experience{},
			Backend:  []experience.Experience{},
			Other:    []experience.Experience{},
		}

		for _, currentExperience := range currentProject.Experiences {
			isOther := true

			for _, currentSkilType := range currentExperience.SkilTypes {

				if currentSkilType == skilType.Frontend {
					project.Frontend = append(project.Frontend, currentExperience)
					isOther = false
				}

				if currentSkilType == skilType.Backend {
					project.Backend = append(project.Backend, currentExperience)
					isOther = false
				}
			}

			if isOther {
				project.Other = append(project.Other, currentExperience)
			}
		}

		projects = append(projects, project)
	}

	resumeTmp := ResumeTmp{
		Resume:   original,
		Projects: projects,
	}

	return resumeTmp
}

// 経歴書を出力する
func outputResume(path string, resume ResumeTmp) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tplStr, err := template.New("tpl").Parse(tpl)
	if err != nil {
		panic(err)
	}

	writer := new(strings.Builder)
	if err := tplStr.Execute(writer, resume); err != nil {
		panic(err)
	}

	file.WriteString(writer.String())
}

// JSONファイルを出力する
func outputJson(path string, resume ResumeTmp) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jsonStr, err := json.Marshal(resume)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	json.Indent(&buf, jsonStr, "", "  ")

	file.WriteString(buf.String())
}
