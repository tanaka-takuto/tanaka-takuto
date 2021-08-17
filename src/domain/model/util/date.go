package util

import (
	"strconv"
	"strings"
	"time"
)

// 日付
type Date struct {
	time.Time
}

// 文字列
func (d Date) String() string {
	if d.IsZero() {
		return ""
	}

	return d.Time.Format("2006/01")
}

// JSONからアンマーシャルする
func (d *Date) UnmarshalJSON(b []byte) error {
	str := string(b)
	split := strings.Split(strings.ReplaceAll(str, "\"", ""), "-")
	year, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}
	monthInt, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
	month := time.Month(monthInt)

	d.Time = time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

	return nil
}

// 月間の差分を取得する
// 終了日付が設定されていない場合は現在日時をまでの差分とする
func (d *Date) MonthDiff(d2 *Date) int {
	var date1, date2 time.Time

	if d.Time.IsZero() {
		date1 = time.Now()
	} else {
		date1 = d.Time
	}
	date2 = d2.Time

	year1 := date1.AddDate(0, 1, 0).Year()
	month1 := date1.AddDate(0, 1, 0).Month()

	year2 := date2.Year()
	month2 := date2.Month()

	year := year1 - year2
	month := month1 - month2

	return year*12 + int(month)
}
