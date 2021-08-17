package skilset

import (
	"app/domain/model/project"
	"app/domain/model/skil"
	"app/domain/model/skilInfo"
	"app/domain/model/skilType"
	"app/domain/model/util"
)

// スキルセット
type SkilSet struct {
	// 言語一覧
	Languages []skil.Skil
	// フレームワーク一覧
	Framework []skil.Skil
	// DB一覧
	DB []skil.Skil
	// 言語一覧
	Tool []skil.Skil
}

// プロジェクトをもとに更新する
func (s *SkilSet) Update(project project.Project, skilInfos []skilInfo.SkilInfo) SkilSet {
	experiencesMonth := project.To.MonthDiff(&project.From)

	for _, currentExperience := range project.Experiences {
		for _, currentSkilType := range currentExperience.SkilTypes {
			switch currentSkilType {
			case skilType.Language:
				s.Languages = updateSkils(s.Languages, currentExperience.Name, experiencesMonth, skilInfos)
				break
			case skilType.Framework:
				s.Framework = updateSkils(s.Framework, currentExperience.Name, experiencesMonth, skilInfos)
				break
			case skilType.DB:
				s.DB = updateSkils(s.DB, currentExperience.Name, experiencesMonth, skilInfos)
				break
			case skilType.Tool:
				s.Tool = updateSkils(s.Tool, currentExperience.Name, experiencesMonth, skilInfos)
				break
			default:
				break
			}
		}
	}

	return *s
}

// スキルを更新する
func updateSkils(skils []skil.Skil, name string, experiencesMonth int, skilInfos []skilInfo.SkilInfo) []skil.Skil {

	targetSkil := findSkil(skils, name)
	if targetSkil != nil {
		targetSkil.ExperienceMonth += skil.ExperienceMonth(experiencesMonth)
		return skils
	}

	targetSkilInfo := findSkilInfo(skilInfos, name)
	if targetSkilInfo == nil {
		panic("存在しないスキルで更新しようとしてます" + name)
	}

	newSkil := skil.Skil{
		SkilInfo:        *targetSkilInfo,
		ExperienceMonth: skil.ExperienceMonth(experiencesMonth),
	}
	skils = append(skils, newSkil)

	return skils
}

// スキルを探す
func findSkil(skils []skil.Skil, name string) *skil.Skil {
	skil := util.Find(skils, name, func(v1 skil.Skil, v2 string) bool {
		return v1.SkilInfo.Name == v2
	})

	return skil
}

// スキル情報を探す
func findSkilInfo(skilInfos []skilInfo.SkilInfo, name string) *skilInfo.SkilInfo {
	skilInfo := util.Find(skilInfos, name, func(v1 skilInfo.SkilInfo, v2 string) bool {
		return v1.Name == v2
	})

	return skilInfo
}
