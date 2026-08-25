package query

import (
	"examarchive/internal/fixture"
	"testing"
)

func TestWorkflowQueryByName(t *testing.T) {
	rs := fixture.Set()
	got := ByName(rs, "ada")
	if len(got) != 1 || got[0].Candidate.ID != "C001" {
		t.Fatal(got)
	}
}
func TestPaging(t *testing.T) {
	p := Paginate(fixture.Set(), 1, 2)
	if len(p.Items) != 2 || p.Total != 4 {
		t.Fatal(p)
	}
}
