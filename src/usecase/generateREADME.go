package usecase

import (
	"app/domain/model/basic"
	"app/domain/model/project"
	"app/domain/model/resume"
	"app/domain/model/skilInfo"
	skilset "app/domain/model/skilSet"
)

type GenerateREADME struct {
	basicRepository    basic.BasicRepository
	skilInfoRepository skilInfo.SkilInfoRepository
	projectRepository  project.ProjectRepository
	resumeRepository   resume.ResumeRepository
}

func NewGenerateREADME(
	basicRepository basic.BasicRepository,
	skilInfoRepository skilInfo.SkilInfoRepository,
	projectRepository project.ProjectRepository,
	resumeRepository resume.ResumeRepository,
) GenerateREADME {
	return GenerateREADME{
		basicRepository:    basicRepository,
		skilInfoRepository: skilInfoRepository,
		projectRepository:  projectRepository,
		resumeRepository:   resumeRepository,
	}
}

func (g GenerateREADME) Execute() error {
	// 基本情報取得
	basic := g.basicRepository.LoadBasic()

	// スキル情報取得
	skilInfos := g.skilInfoRepository.LoadSkilInfo()

	// プロジェクト情報取得
	projects := g.projectRepository.LoadProject()

	// プロジェクトをもとにスキル情報更新
	skilSet := skilset.SkilSet{}
	for _, project := range projects {
		skilSet = skilSet.Update(project, skilInfos)
	}

	// 経歴書情報設定
	resume := resume.Resume{
		Aliases:  skilInfos,
		Basic:    basic,
		SkilSet:  skilSet,
		Projects: projects,
	}

	// 経歴書生成
	if err := g.resumeRepository.Output(resume); err != nil {
		return err
	}

	return nil
}
