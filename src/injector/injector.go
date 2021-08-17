package injector

import (
	"app/domain/model/basic"
	"app/domain/model/project"
	"app/domain/model/resume"
	"app/domain/model/skilInfo"

	infra_basic "app/infra/basic"
	infra_project "app/infra/project"
	infra_resume "app/infra/resume"
	infra_SkilInfo "app/infra/skilInfo"
)

type Injector struct {
	BasicRepository    basic.BasicRepository
	SkilInfoRepository skilInfo.SkilInfoRepository
	ProjectRepository  project.ProjectRepository
	ResumeRepository   resume.ResumeRepository
}

func NewInjector(
	basicPath string,
	skilInfoPath string,
	projectPath string,
	outputPaht string,
) Injector {
	basicRepository := infra_basic.NewBasicRepository(basicPath)
	skilInfoRpository := infra_SkilInfo.NewSkilInfoRepository(skilInfoPath)
	projectRepository := infra_project.NewProjectRepository(projectPath)
	resumeRepository := infra_resume.NewResumeRepository(outputPaht)

	return Injector{
		BasicRepository:    basicRepository,
		SkilInfoRepository: skilInfoRpository,
		ProjectRepository:  projectRepository,
		ResumeRepository:   resumeRepository,
	}
}
