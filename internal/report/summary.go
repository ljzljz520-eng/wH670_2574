package report

import (
	"examarchive/internal/model"
	"examarchive/internal/scoring"
	"fmt"
)

func Summary(records []model.ExamRecord) string {
	s := scoring.Summarize(records)
	return fmt.Sprintf("count=%d mean=%.2f highest=%.2f lowest=%.2f pass_rate=%.2f", s.Count, s.Mean, s.Highest, s.Lowest, scoring.PassRate(records))
}
func GradeCounts(records []model.ExamRecord) map[string]int { return scoring.Distribution(records) }
func Empty(records []model.ExamRecord) bool                 { return len(records) == 0 }
