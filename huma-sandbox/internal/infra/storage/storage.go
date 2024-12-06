package storage

import "database/sql"

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

// インターフェイスを使わない。管理下にある依存なので、テストはインテグレーションテストを行うのがリファクタへの耐性が高くなる
// by 『単体テストの考え方/使い方』
