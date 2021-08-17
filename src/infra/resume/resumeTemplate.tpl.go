package infra_resume

// テンプレート
const tpl = `<!-- 画像登録 -->
{{range $i, $v := .Aliases}}[{{$v.Name}}]: https://img.shields.io/badge/-{{$v.Name}}-555555?logo={{$v.Icon}}&logoColor=ffffff&labelColor={{$v.Color.ImageColor}}
{{end}}

# 基本情報
|  key  |          value          |
| :---: | :---------------------: |
| 氏名  | {{.Basic.Name}}({{.Basic.AlpName}}) |

![](https://github-profile-summary-cards.vercel.app/api/cards/profile-details?username=tanaka-takuto)
![](https://github-profile-summary-cards.vercel.app/api/cards/repos-per-language?username=tanaka-takuto)
![](https://github-profile-summary-cards.vercel.app/api/cards/most-commit-language?username=tanaka-takuto)
![](https://github-profile-summary-cards.vercel.app/api/cards/stats?username=tanaka-takuto)

{{if gt (len .Basic.Likes) 0}}## 好きなこと
{{range $i, $v := .Basic.Likes}}- {{$v}}
{{end}}{{end}}

## 得意なこと
{{if gt (len .Basic.GoodAt) 0}}{{range $i, $v := .Basic.GoodAt}}- {{$v}}
{{end}}{{end}}

## したいこと
{{if gt (len .Basic.WantToDo) 0}}{{range $i, $v := .Basic.WantToDo}}- {{$v}}
{{end}}{{end}}

# スキル

<!-- https://shields.io/ -->

## 言語
{{range $i, $v := .SkilSet.Languages}}![](https://img.shields.io/badge/{{$v.SkilInfo.Name}}-{{$v.ExperienceMonth}}-555555?logo={{$v.SkilInfo.Icon}}&logoColor={{$v.SkilInfo.Color.FontColor}}&labelColor={{$v.SkilInfo.Color.ImageColor}})
{{end}}

## フレームワーク
{{range $i, $v := .SkilSet.Framework}}![](https://img.shields.io/badge/{{$v.SkilInfo.Name}}-{{$v.ExperienceMonth}}-555555?logo={{$v.SkilInfo.Icon}}&logoColor={{$v.SkilInfo.Color.FontColor}}&labelColor={{$v.SkilInfo.Color.ImageColor}})
{{end}}

## DB
{{range $i, $v := .SkilSet.DB}}![](https://img.shields.io/badge/{{$v.SkilInfo.Name}}-{{$v.ExperienceMonth}}-555555?logo={{$v.SkilInfo.Icon}}&logoColor={{$v.SkilInfo.Color.FontColor}}&labelColor={{$v.SkilInfo.Color.ImageColor}})
{{end}}

## ツール
{{range $i, $v := .SkilSet.Tool}}![](https://img.shields.io/badge/{{$v.SkilInfo.Name}}-{{$v.ExperienceMonth}}-555555?logo={{$v.SkilInfo.Icon}}&logoColor={{$v.SkilInfo.Color.FontColor}}&labelColor={{$v.SkilInfo.Color.ImageColor}})
{{end}}


# 職務経歴書
{{range $i, $project := .Projects}}
## 【{{$project.Title}}】({{$project.From}} ~ {{$project.To}})
{{$project.Description}}

| 担当業務 | フロントエンド | バックエンド | その他 |
| :------: | :------------: | :----------: | :----: |
| {{range $i, $operation := $project.Operations}}{{if ne $i 0}} / {{end}}{{$operation}}{{end}} ` +
	`| {{if eq (len $project.Frontend) 0}}-{{end}}{{range $i, $experience := $project.Frontend}}![{{$experience.Name}}] {{end}} ` +
	`| {{if eq (len $project.Backend) 0}}-{{end}}{{range $i, $experience := $project.Backend}}![{{$experience.Name}}] {{end}} ` +
	`| {{if eq (len $project.Other) 0}}-{{end}}{{range $i, $experience := $project.Other}}![{{$experience.Name}}] {{end}} |
{{end}}
`
