// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package generated

import (
	"database/sql"
)

type Account struct {
	ID       int32
	Username sql.NullString
	CreatAt  sql.NullTime
}
