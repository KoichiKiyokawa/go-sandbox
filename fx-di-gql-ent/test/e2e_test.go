package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestE2e(t *testing.T) {
	json, err := json.Marshal(map[string]string{"query": "{health}"})
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8080/query", "application/json", bytes.NewBuffer(json))
	if err != nil {
		t.Fatal(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	got := string(b)
	want := `{"data":{"health":"ok"}}`
	assert.Equal(t, want, got)
}
