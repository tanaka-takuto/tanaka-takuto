package util

import (
	"encoding/json"
	"os"
)

// ファイルが存在するか
func IsExistsFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// JSONから取得する
func FromJson[T any](path string) T {
	var t T
	b, _ := os.ReadFile(path)
	err := json.Unmarshal(b, &t)
	if err != nil {
		panic(err)
	}
	return t

}

// 重複を削除する
func Uniq[T comparable](array []T) []T {
	result := []T{}
	m := make(map[T]struct{})
	for _, v := range array {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

// 存在するか
func Find[T1 any, T2 any](array []T1, value T2, match func(v1 T1, v2 T2) bool) *T1 {
	for i, current := range array {
		if match(current, value) {
			return &array[i]
		}
	}

	return nil
}
