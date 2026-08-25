package report

import (
	"examarchive/internal/model"
	"examarchive/internal/scoring"
	"fmt"
	"strings"
)

func Header() string { return "ID\tName\tListening\tReading\tWriting\tSpeaking\tAverage\tGrade\tRank" }
func Line(r model.ExamRecord) string {
	return fmt.Sprintf("%s\t%s\t%.1f\t%.1f\t%.1f\t%.1f\t%.2f\t%s\t%d", r.Candidate.ID, r.Candidate.Name, r.Score.Listening, r.Score.Reading, r.Score.Writing, r.Score.SpeakingValue(), r.Average, scoring.Grade(r.Average), r.Rank)
}
func Table(records []model.ExamRecord) string {
	rows := []string{Header()}
	for _, r := range scoring.Rank(records) {
		rows = append(rows, Line(r))
	}
	return strings.Join(rows, "\n")
}
func CSV(records []model.ExamRecord) string {
	rows := []string{"id,name,listening,reading,writing,speaking,average,grade,rank"}
	for _, r := range scoring.Rank(records) {
		rows = append(rows, fmt.Sprintf("%s,%s,%.1f,%.1f,%.1f,%.1f,%.2f,%s,%d", r.Candidate.ID, r.Candidate.Name, r.Score.Listening, r.Score.Reading, r.Score.Writing, r.Score.SpeakingValue(), r.Average, scoring.Grade(r.Average), r.Rank))
	}
	return strings.Join(rows, "\n")
}
