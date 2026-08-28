package service

import (
	"context"
	"examarchive/internal/fixture"
	"examarchive/internal/store"
	"path/filepath"
	"testing"
)

func TestWorkflowCreateAndRank(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "a.db"))
	defer s.Close()
	a := New(s)
	for _, r := range fixture.Set() {
		if _, e := a.Create(r.Candidate, r.Score); e != nil {
			t.Fatal(e)
		}
	}
	rs, _ := a.List()
	if len(rs) != 4 {
		t.Fatal(len(rs))
	}
}
func TestWorkflowExport(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "a.db"))
	defer s.Close()
	a := New(s)
	if e := a.Import(context.Background(), fixture.Set()); e != nil {
		t.Fatal(e)
	}
	rs, e := a.Export(context.Background())
	if e != nil || len(rs) != 4 {
		t.Fatal(e, len(rs))
	}
}
