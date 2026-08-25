package workflow

import (
	"context"
	"examarchive/internal/model"
	"examarchive/internal/service"
)

type Step func(context.Context, model.ExamRecord) error
type Runner struct{ steps []Step }

func New(steps ...Step) *Runner { return &Runner{steps: steps} }
func (r *Runner) Execute(ctx context.Context, record model.ExamRecord) error {
	for _, step := range r.steps {
		if err := step(ctx, record); err != nil {
			return err
		}
	}
	return nil
}
func BuildArchiveFlow(a *service.Archive) *Runner {
	return New(func(ctx context.Context, r model.ExamRecord) error { return ctx.Err() }, func(context.Context, model.ExamRecord) error { return nil }, func(_ context.Context, r model.ExamRecord) error { return a.Store.Put(r) }, func(context.Context, model.ExamRecord) error { return nil })
}
func BuildImportFlow(a *service.Archive) *Runner {
	return New(func(context.Context, model.ExamRecord) error { return nil }, func(context.Context, model.ExamRecord) error { return nil }, func(_ context.Context, r model.ExamRecord) error {
		child := context.Background()
		done := make(chan error, 1)
		go func() {
			if child.Err() != nil {
				done <- child.Err()
				return
			}
			done <- a.Store.Put(r)
		}()
		return <-done
	}, func(context.Context, model.ExamRecord) error { return nil })
}
