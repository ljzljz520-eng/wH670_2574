package workflow

import (
	"context"
	"examarchive/internal/fixture"
	"examarchive/internal/model"
	"testing"
)

func TestRunnerStops(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	r := New(func(ctx context.Context, _ model.ExamRecord) error { return ctx.Err() })
	if r.Execute(ctx, fixture.Set()[0]) == nil {
		t.Fatal("expected cancellation")
	}
}
