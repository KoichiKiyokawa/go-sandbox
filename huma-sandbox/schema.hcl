schema "public" {}

table "users" {
  schema = schema.public
  column "id" {
    type = char(26)
  }
  column "name" {
    type = varchar(255)
  }
  column "nickname" {
    type = varchar(255)
    null = true
  }
  column "created_at" {
    type = timestamp
  }
  column "updated_at" {
    type = timestamp
  }
  primary_key {
    columns = [column.id]
  }
}

table "user_secrets" {
  schema = schema.public
  column "user_id" {
    type = char(26)
  }
  column "password_hash" {
    type = varchar(255)
  }
  foreign_key "user_fk" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
  }
}