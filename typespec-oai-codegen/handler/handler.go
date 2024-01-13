package handler

import (
	"typespec-oai-codegen/generated"
	"typespec-oai-codegen/generated/db"
)

type handler struct {
	queries db.Queries
}

func NewHandler(queries db.Queries) generated.StrictServerInterface {
	return &handler{queries: queries}
}
