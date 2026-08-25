package main

import (
	"context"
	"examarchive/internal/fixture"
	"examarchive/internal/service"
	"examarchive/internal/store"
	"examarchive/internal/workflow"
	"path/filepath"
	"testing"
)

func TestBusinessChain15(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "bug.db"))
	defer s.Close()
	a := service.New(s)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if e := workflow.BuildImportFlow(a).Execute(ctx, fixture.Set()[0]); e != nil {
		t.Fatal(e)
	}
	if _, e := s.Get("C001"); e == nil {
		t.Fatal("cancelled workflow wrote a result")
	}
}
