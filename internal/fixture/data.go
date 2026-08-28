package fixture

import "examarchive/internal/model"

func Record(id, name string, avg float64) model.ExamRecord {
	return model.ExamRecord{Candidate: model.NewCandidate(id, name, id+"@example.com"), Score: model.Score{CandidateID: id, Listening: avg, Reading: avg, Writing: avg, SpeakingPercent: &avg}, Average: avg}
}
func Set() []model.ExamRecord {
	return []model.ExamRecord{Record("C001", "Ada Chen", 88), Record("C002", "Bo Lin", 76), Record("C003", "Cy Rao", 94), Record("C004", "Dee Wu", 65)}
}
