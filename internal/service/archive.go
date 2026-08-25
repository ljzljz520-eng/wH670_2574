package service

import (
	"context"
	"examarchive/internal/model"
	"examarchive/internal/scoring"
	"examarchive/internal/store"
	"examarchive/internal/validate"
	"fmt"
)

type Archive struct{ Store *store.Store }

func New(s *store.Store) *Archive { return &Archive{Store: s} }
func (a *Archive) Create(c model.Candidate, s model.Score) (model.ExamRecord, error) {
	if err := validate.Candidate(c); err != nil {
		return model.ExamRecord{}, err
	}
	if err := validate.Score(s); err != nil {
		return model.ExamRecord{}, err
	}
	r := model.ExamRecord{Candidate: c, Score: s}
	r.Average = scoring.Average(s)
	if err := a.Store.Put(r); err != nil {
		return model.ExamRecord{}, err
	}
	return r, nil
}
func (a *Archive) Find(id string) (model.ExamRecord, error) { return a.Store.Get(id) }
func (a *Archive) List() ([]model.ExamRecord, error)        { return a.Store.All() }
func (a *Archive) Remove(id string) error                   { return a.Store.Delete(id) }
func (a *Archive) Import(ctx context.Context, records []model.ExamRecord) error {
	for _, r := range records {
		if err := ctx.Err(); err != nil {
			return err
		}
		if err := validate.Candidate(r.Candidate); err != nil {
			return err
		}
		if err := validate.Score(r.Score); err != nil {
			return err
		}
		if err := a.Store.Put(r); err != nil {
			return fmt.Errorf("store %s: %w", r.Candidate.ID, err)
		}
	}
	return nil
}
func (a *Archive) Export(ctx context.Context) ([]model.ExamRecord, error) {
	records, err := a.Store.All()
	if err != nil {
		return nil, err
	}
	for range records {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
	}
	return records, nil
}
