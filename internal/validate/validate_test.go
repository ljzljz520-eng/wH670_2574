package validate

import (
	"examarchive/internal/fixture"
	"testing"
)

func TestValidBatch(t *testing.T) {
	if e := Batch(fixture.Set()); e != nil {
		t.Fatal(e)
	}
}
func TestRejectBadEmail(t *testing.T) {
	r := fixture.Record("x", "X", 70)
	r.Candidate.Email = "bad"
	if Candidate(r.Candidate) == nil {
		t.Fatal("accepted")
	}
}
