package scoring

import (
	"examarchive/internal/model"
	"sort"
)

func Average(s model.Score) float64 {
	return (s.Listening + s.Reading + s.Writing + s.SpeakingValue()) / 4
}
func Rank(records []model.ExamRecord) []model.ExamRecord {
	out := append([]model.ExamRecord(nil), records...)
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Average == out[j].Average {
			return out[i].Candidate.ID < out[j].Candidate.ID
		}
		return out[i].Average > out[j].Average
	})
	for i := range out {
		out[i].Rank = i + 1
	}
	return out
}
func Grade(avg float64) string {
	switch {
	case avg >= 90:
		return "excellent"
	case avg >= 75:
		return "strong"
	case avg >= 60:
		return "pass"
	default:
		return "needs-review"
	}
}
func Weighted(s model.Score, weights [4]float64) float64 {
	return s.Listening*weights[0] + s.Reading*weights[1] + s.Writing*weights[2] + s.SpeakingValue()*weights[3]
}
