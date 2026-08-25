package exam

import (
	"examarchive/internal/fixture"
	"examarchive/internal/report"
	"examarchive/internal/service"
	"examarchive/internal/store"
	"fmt"
	"os"
	"path/filepath"
)

func Run() {
	path := filepath.Join(os.TempDir(), "examarchive.db")
	s, err := store.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer s.Close()
	a := service.New(s)
	for _, r := range fixture.Set() {
		if _, err := a.Create(r.Candidate, r.Score); err != nil {
			fmt.Println(err)
		}
	}
	records, _ := a.List()
	fmt.Println(report.Table(records))
}
