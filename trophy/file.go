package trophy

import (
	"crypto/sha256"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"xztaityozx/zsh-trophy/record"
)

type CatCheck struct {
	Title     string
	FileName  string
	Desc      string
	Grade     string
	OneChance bool
}

var catRegexp = regexp.MustCompile("^cat [^ ]+")

func (c CatCheck) Check(cmd string, r record.Record) (Trophy, error) {
	nk := fmt.Sprintf("%x", sha256.Sum256([]byte(c.Title+c.Desc+c.Grade+c.FileName)))

	if !catRegexp.MatchString(cmd) {
		return Trophy{Cleared: false}, nil
	}

	if never, ok := r.Args[nk]; ok && never == "true" {
		return Trophy{Cleared: false}, nil
	}

	a, _ := filepath.Abs(c.FileName)
	b, _ := filepath.Abs(strings.Split(cmd, " ")[1])

	return Trophy{
		Cleared: a == b,
		Title:   c.Title,
		Desc:    c.Desc,
		Grade:   c.Grade,
		Values: map[string]string{
			nk: strings.ToLower(fmt.Sprint(c.OneChance)),
		},
	}, nil
}
