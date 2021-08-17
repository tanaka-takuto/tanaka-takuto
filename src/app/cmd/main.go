package main

import (
	"app/domain/model/util"
	"app/injector"
	"app/usecase"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 5 {
		panic("引数が不正です。")
	}
	basicPath := os.Args[1]
	skilInfoPath := os.Args[2]
	projectsPath := os.Args[3]
	outputPath := os.Args[4]
	if !util.IsExistsFile(basicPath) ||
		!util.IsExistsFile(skilInfoPath) ||
		!util.IsExistsFile(projectsPath) {
		panic("ファイルが存在しません。")
	}

	injector := injector.NewInjector(basicPath, skilInfoPath, projectsPath, outputPath)

	fmt.Println("-- Start --")
	generateREADME := usecase.NewGenerateREADME(injector.BasicRepository, injector.SkilInfoRepository, injector.ProjectRepository, injector.ResumeRepository)
	fmt.Println("-- End --")

	generateREADME.Execute()
}
