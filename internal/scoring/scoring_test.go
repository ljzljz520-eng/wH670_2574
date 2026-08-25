package scoring

import (
	"examarchive/internal/fixture"
	"testing"
)

func TestRanking(t *testing.T) {
	rs := Rank(fixture.Set())
	if rs[0].Candidate.ID != "C003" || rs[0].Rank != 1 {
		t.Fatal(rs)
	}
}
func TestGrades(t *testing.T) {
	if Grade(91) != "excellent" || Grade(40) != "needs-review" {
		t.Fatal("grade")
	}
}
