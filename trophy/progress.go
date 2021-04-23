package trophy

import (
	"fmt"
	"xztaityozx/zsh-trophy/record"
)

type Progress struct {
	N       int
	Grade   string
	Comment string
}

func (p Progress) Check(_ string, r record.Record) (Trophy, error) {
	v, ok := r.Args["progress"]

	return Trophy{
		Cleared: ok && v == fmt.Sprint(p.N),
		Grade:   p.Grade,
		Title:   fmt.Sprintf("%d個目のトロフィー", p.N),
		Desc:    fmt.Sprintf("%d個目のトロフィーを獲得した\n%s", p.N, p.Comment),
	}, nil
}
