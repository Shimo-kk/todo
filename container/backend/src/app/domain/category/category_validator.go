package category

import (
	"todo/app/core"
	"unicode/utf8"
)

// カテゴリ名のバリデーション
func validateName(value string) error {
	if value == "" {
		return core.NewError(core.ValidationError, "カテゴリ名は空の値を入力することはできません。")
	}

	if utf8.RuneCountInString(value) > 50 {
		return core.NewError(core.ValidationError, "カテゴリ名は50文字より大きい値を入力できません。")
	}

	return nil
}
