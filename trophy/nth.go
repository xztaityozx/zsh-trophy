package trophy

import (
	"errors"
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

func parse(m map[string]string, k string) int {
	if val, ok := m[k]; ok {
		if rt, err := strconv.Atoi(val); err == nil {
			return rt
		}
	}

	return 0
}

func (n NthCmd) Check(cmd string, r record.Record) (Trophy, error) {
	if !regexp.MustCompile(fmt.Sprintf("^%s ?", n.Command)).MatchString(cmd) {
		return Trophy{Cleared: false}, nil
	}

	cnt := parse(r.Args, n.Key)

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

type NthRegexp struct {
	Re    *regexp.Regexp
	Key   string
	Grade string
	Title string
	Desc  string
	Count int
}

func (n NthRegexp) Check(cmd string, record record.Record) (Trophy, error) {
	if n.Re == nil {
		return Trophy{}, errors.New("regexp is nil")
	}

	if !n.Re.MatchString(cmd) {
		return Trophy{Cleared: false}, nil
	}

	cnt := parse(record.Args, n.Key)

	return Trophy{
		Cleared: n.Count-1 == cnt,
		Title:   n.Title,
		Grade:   n.Grade,
		Desc:    n.Desc,
		Values:  map[string]string{n.Key: fmt.Sprint(cnt + 1)},
	}, nil
}
