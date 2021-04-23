package trophy

import (
	"log"
	"regexp"
	"xztaityozx/zsh-trophy/record"
)

type Count struct {
	Re    *regexp.Regexp
	N     int
	Title string
	Desc  string
	Grade string
}

func (c Count) Check(cmd string, _ record.Record) (Trophy, error) {
	cnt := len(c.Re.FindAllString(cmd, -1))
	log.Println(cnt)
	return Trophy{
		Title:   c.Title,
		Desc:    c.Desc,
		Grade:   c.Grade,
		Cleared: c.N == cnt+1,
	}, nil
}
