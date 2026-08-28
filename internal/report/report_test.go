package report

import (
	"examarchive/internal/fixture"
	"strings"
	"testing"
)

func TestWorkflowReport(t *testing.T) {
	s := Table(fixture.Set())
	if !strings.Contains(s, "C003") || !strings.Contains(s, "ID") {
		t.Fatal(s)
	}
}
