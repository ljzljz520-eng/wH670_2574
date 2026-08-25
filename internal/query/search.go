package query

import (
	"examarchive/internal/model"
	"examarchive/internal/scoring"
	"sort"
	"strings"
)

func ByID(records []model.ExamRecord, id string) (model.ExamRecord, bool) {
	for _, r := range records {
		if r.Candidate.ID == id {
			return r, true
		}
	}
	return model.ExamRecord{}, false
}
func ByName(records []model.ExamRecord, name string) []model.ExamRecord {
	needle := strings.ToLower(strings.TrimSpace(name))
	out := []model.ExamRecord{}
	for _, r := range records {
		if strings.Contains(strings.ToLower(r.Candidate.Name), needle) {
			out = append(out, r)
		}
	}
	return out
}
func Ranked(records []model.ExamRecord) []model.ExamRecord { return scoring.Rank(records) }
func Top(records []model.ExamRecord, n int) []model.ExamRecord {
	r := Ranked(records)
	if n < 0 {
		n = 0
	}
	if n > len(r) {
		n = len(r)
	}
	return r[:n]
}
func FilterGrade(records []model.ExamRecord, grade string) []model.ExamRecord {
	out := []model.ExamRecord{}
	for _, r := range records {
		if scoring.Grade(r.Average) == grade {
			out = append(out, r)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Average > out[j].Average })
	return out
}
