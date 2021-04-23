package trophy

import (
	"errors"
	"regexp"
	"xztaityozx/zsh-trophy/record"
)

type SimpleRegexp struct {
	Re    *regexp.Regexp
	Grade string
	Title string
	Desc  string
}

func (sr SimpleRegexp) Check(cmd string, r record.Record) (Trophy, error) {
	if sr.Re == nil {
		return Trophy{}, errors.New("Regexp is nil")
	}

	return Trophy{
		Cleared: sr.Re.MatchString(cmd),
		Title:   sr.Title,
		Desc:    sr.Desc,
		Grade:   sr.Grade,
	}, nil
}
