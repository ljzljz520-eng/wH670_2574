package validate

import (
	"examarchive/internal/model"
	"fmt"
)

func Batch(records []model.ExamRecord) error {
	if len(records) == 0 {
		return fmt.Errorf("empty batch")
	}
	if len(records) > 10000 {
		return fmt.Errorf("batch too large")
	}
	if err := UniqueIDs(records); err != nil {
		return err
	}
	for _, r := range records {
		if err := Candidate(r.Candidate); err != nil {
			return err
		}
		if err := Score(r.Score); err != nil {
			return err
		}
	}
	return nil
}
func EmailDomain(email string) string {
	for i, c := range email {
		if c == '@' {
			return email[i+1:]
		}
	}
	return ""
}
func IsInstitutional(email string) bool {
	d := EmailDomain(email)
	return d == "university.example" || d == "school.example"
}
