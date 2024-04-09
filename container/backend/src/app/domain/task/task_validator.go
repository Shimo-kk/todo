package task

import (
	"time"
	"todo/app/core"
	"unicode/utf8"
)

// タイトルのバリデーション
func validateTitle(value string) error {
	if value == "" {
		return core.NewError(core.ValidationError, "タイトルは空の値を入力することはできません。")
	}

	if utf8.RuneCountInString(value) > 50 {
		return core.NewError(core.ValidationError, "タイトルは50文字より大きい値を入力することはできません。")
	}

	return nil
}

// 開始日のバリデーション
func validateStartDate(value time.Time) error {
	if value.Before(time.Now()) {
		return core.NewError(core.ValidationError, "開始日は過去の日付を入力することはできません。")
	}
	return nil
}
