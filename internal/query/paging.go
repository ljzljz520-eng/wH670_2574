package query

import "examarchive/internal/model"

type Page struct {
	Items                []model.ExamRecord
	Offset, Limit, Total int
}

func Paginate(records []model.ExamRecord, offset, limit int) Page {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 20
	}
	if offset > len(records) {
		offset = len(records)
	}
	end := offset + limit
	if end > len(records) {
		end = len(records)
	}
	return Page{Items: records[offset:end], Offset: offset, Limit: limit, Total: len(records)}
}
func IDs(records []model.ExamRecord) []string {
	out := make([]string, 0, len(records))
	for _, r := range records {
		out = append(out, r.Candidate.ID)
	}
	return out
}
func Names(records []model.ExamRecord) []string {
	out := make([]string, 0, len(records))
	for _, r := range records {
		out = append(out, r.Candidate.Name)
	}
	return out
}
