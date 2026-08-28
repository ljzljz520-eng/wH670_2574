package scoring

import "examarchive/internal/model"

type Summary struct {
	Count                 int
	Mean, Highest, Lowest float64
}

func Summarize(rs []model.ExamRecord) Summary {
	if len(rs) == 0 {
		return Summary{}
	}
	s := Summary{Count: len(rs), Lowest: rs[0].Average}
	for _, r := range rs {
		s.Mean += r.Average
		if r.Average > s.Highest {
			s.Highest = r.Average
		}
		if r.Average < s.Lowest {
			s.Lowest = r.Average
		}
	}
	s.Mean /= float64(s.Count)
	return s
}
func PassRate(rs []model.ExamRecord) float64 {
	if len(rs) == 0 {
		return 0
	}
	n := 0
	for _, r := range rs {
		if r.Average >= 60 {
			n++
		}
	}
	return float64(n) / float64(len(rs))
}
func Distribution(rs []model.ExamRecord) map[string]int {
	d := map[string]int{}
	for _, r := range rs {
		d[Grade(r.Average)]++
	}
	return d
}
