package trophy

import (
	"fmt"
	"regexp"
	"strconv"
	"xztaityozx/zsh-trophy/record"
)

type NthCmd struct {
	Command string
	Key     string
	Grade   string
	Title   string
	Desc    string
	Count   int
}

func (n NthCmd) Check(cmd string, r record.Record) (Trophy, error) {
	if !regexp.MustCompile(fmt.Sprintf("^%s ?", n.Command)).MatchString(cmd) {
		return Trophy{Cleared: false}, nil
	}

	cnt := func() int {
		if val, ok := r.Args[n.Key]; ok {
			if rt, err := strconv.Atoi(val); err == nil {
				return rt
			}
		}
		return 0
	}()

	return Trophy{
		Cleared: n.Count-1 == cnt,
		Title:   n.Title,
		Desc:    n.Desc,
		Grade:   n.Grade,
		Values: map[string]string{
			n.Key: fmt.Sprintf("%d", cnt+1),
		},
	}, nil
}
