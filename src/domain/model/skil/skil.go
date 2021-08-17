package skil

import (
	"app/domain/model/skilInfo"
	"fmt"
	"math"
)

// 経験年数
type ExperienceMonth int

// スキル
type Skil struct {
	// スキル情報
	SkilInfo skilInfo.SkilInfo

	// 経験年数
	ExperienceMonth ExperienceMonth
}

// 文字列
func (e ExperienceMonth) String() string {
	year := int(math.Floor(float64(e) / 12))
	month := int(e) - year*12

	var yearStr, monthStr string
	switch year {
	case 0:
		yearStr = ""
	case 1:
		yearStr = "1year"
	default:
		yearStr = fmt.Sprintf("%dyears", year)
	}

	switch month {
	case 0:
		monthStr = ""
	case 1:
		monthStr = "1month"
	default:
		monthStr = fmt.Sprintf("%dmonths", month)
	}

	return yearStr + monthStr
}
