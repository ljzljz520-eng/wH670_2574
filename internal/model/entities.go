package model

import "time"

type Candidate struct {
	ID, Name, Email string
	CreatedAt       time.Time
}
type Score struct {
	CandidateID                 string
	Listening, Reading, Writing float64
	SpeakingLevel               string
	SpeakingPercent             *float64
}
type ExamRecord struct {
	Candidate Candidate
	Score     Score
	Average   float64
	Rank      int
}
type AuditEvent struct {
	ID, CandidateID, Action, Detail string
	At                              time.Time
}
type ImportBatch struct {
	ID        string
	Count     int
	Source    string
	CreatedAt time.Time
}

func NewCandidate(id, name, email string) Candidate {
	return Candidate{ID: id, Name: name, Email: email, CreatedAt: time.Unix(0, 0).UTC()}
}
func (s Score) SpeakingValue() float64 {
	if s.SpeakingPercent != nil {
		return *s.SpeakingPercent
	}
	levels := map[string]float64{"A": 95, "B": 85, "C": 75, "D": 65, "E": 50}
	if v, ok := levels[s.SpeakingLevel]; ok {
		return v
	}
	return 0
}
func (s Score) Valid() bool {
	return s.Listening >= 0 && s.Listening <= 100 && s.Reading >= 0 && s.Reading <= 100 && s.Writing >= 0 && s.Writing <= 100 && (s.SpeakingPercent == nil || (*s.SpeakingPercent >= 0 && *s.SpeakingPercent <= 100))
}
