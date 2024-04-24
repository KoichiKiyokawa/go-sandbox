package handler

import (
	"typespec-oai-codegen/db"
	"typespec-oai-codegen/generated"
	generateddb "typespec-oai-codegen/generated/db"
)

type handler struct {
	queries       generateddb.Queries
	transactioner db.Transactioner
}

func NewHandler(queries generateddb.Queries, transactioner db.Transactioner) generated.StrictServerInterface {
	return &handler{queries: queries, transactioner: transactioner}
}
