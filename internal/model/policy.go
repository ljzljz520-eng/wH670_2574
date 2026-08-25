package model

type SearchPolicy struct {
	CaseSensitive bool
	Exact         bool
	Limit         int
}

func (p SearchPolicy) Normalize() SearchPolicy {
	if p.Limit < 0 {
		p.Limit = 0
	}
	if p.Limit > 100 {
		p.Limit = 100
	}
	return p
}
func (p SearchPolicy) Match(value, term string) bool {
	if !p.CaseSensitive {
		value, term = lower(value), lower(term)
	}
	if p.Exact {
		return value == term
	}
	return contains(value, term)
}
func lower(s string) string {
	if s >= "A" && s <= "Z" {
		return string(s[0]+32) + s[1:]
	}
	return s
}
func contains(a, b string) bool {
	for i := 0; i+len(b) <= len(a); i++ {
		if a[i:i+len(b)] == b {
			return true
		}
	}
	return b == ""
}
