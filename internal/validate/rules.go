package validate

import (
	"examarchive/internal/model"
	"fmt"
	"regexp"
)

var emailRE = regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)

func Candidate(c model.Candidate) error {
	if c.ID == "" {
		return fmt.Errorf("id required")
	}
	if c.Name == "" {
		return fmt.Errorf("name required")
	}
	if !emailRE.MatchString(c.Email) {
		return fmt.Errorf("invalid email")
	}
	return nil
}
func Score(s model.Score) error {
	if !s.Valid() {
		return fmt.Errorf("score out of range")
	}
	if s.SpeakingLevel == "" && s.SpeakingPercent == nil {
		return fmt.Errorf("speaking required")
	}
	return nil
}
func UniqueIDs(records []model.ExamRecord) error {
	seen := map[string]bool{}
	for _, r := range records {
		if seen[r.Candidate.ID] {
			return fmt.Errorf("duplicate candidate %s", r.Candidate.ID)
		}
		seen[r.Candidate.ID] = true
	}
	return nil
}
func SearchTerm(term string) error {
	if len(term) > 80 {
		return fmt.Errorf("search term too long")
	}
	return nil
}
