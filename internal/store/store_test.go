package store

import (
	"examarchive/internal/fixture"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "scores.db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := fixture.Record("P1", "Persist", 80)
	if e = s.Put(r); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.Get("P1")
	if e != nil || got.Candidate.Name != "Persist" {
		t.Fatalf("reopen failed: %v", e)
	}
}
