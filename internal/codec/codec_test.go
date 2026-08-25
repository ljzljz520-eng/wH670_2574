package codec

import (
	"examarchive/internal/fixture"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	r := fixture.Record("x", "X", 70)
	b, e := Encode(r)
	if e != nil {
		t.Fatal(e)
	}
	got, e := Decode(b)
	if e != nil || got.Candidate.ID != "x" {
		t.Fatal(e)
	}
}
