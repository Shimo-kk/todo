package user

import (
	"time"
	"todo/app/core"

	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	id        int
	createdAt time.Time
	updatedAt time.Time
	name      string
	email     string
	password  string
}

// エンティティの作成
func NewEntity(name string, email string, password string) (*UserEntity, error) {
	// バリデーション
	if err := validateName(name); err != nil {
		return nil, err
	}
	if err := validateEmail(email); err != nil {
		return nil, err
	}
	if err := validatePassword(password); err != nil {
		return nil, err
	}

	// パスワードの暗号化
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, core.NewError(core.SystemError, "パスワードの暗号化に失敗しました。->"+err.Error())
	}

	return &UserEntity{name: name, email: email, password: string(hashed)}, nil
}

// パスワードの検証
func (e *UserEntity) VerifyPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(e.password), []byte(password)); err != nil {
		return core.NewError(core.BadRequestError, "パスワードが正しくありません。")
	}

	return nil
}

// ユーザー名の変更
func (e *UserEntity) ChangeName(name string) error {
	// バリデーション
	if err := validateName(name); err != nil {
		return err
	}

	e.name = name
	return nil
}

// パスワードの変更
func (e *UserEntity) ChangePassword(password string) error {
	// バリデーション
	if err := validatePassword(password); err != nil {
		return err
	}

	// パスワードの暗号化
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return core.NewError(core.SystemError, "パスワードの暗号化に失敗しました。->"+err.Error())
	}

	e.password = string(hashed)
	return nil
}

// ゲッター　ID
func (e *UserEntity) GetId() int {
	return e.id
}

// ゲッター　作成日時
func (e *UserEntity) GetCreatedAt() time.Time {
	return e.createdAt
}

// ゲッター　更新日時
func (e *UserEntity) GetUpdatedAt() time.Time {
	return e.updatedAt
}

// ゲッター　名称
func (e *UserEntity) GetName() string {
	return e.name
}

// ゲッター　E-mailアドレス
func (e *UserEntity) GetEmail() string {
	return e.email
}

// ゲッター　パスワード
func (e *UserEntity) GetPassword() string {
	return e.password
}
