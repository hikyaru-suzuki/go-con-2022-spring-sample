// Code generated by protoc-gen-sample. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package transaction

import "time"

// ユーザーのカード
type UserCard struct {
	// ユーザーID
	UserID string
	// カードID
	CardID string
	// レベル
	Level int64
	// 作成日時
	CreatedTime time.Time
	// 更新日時
	UpdatedTime time.Time
}