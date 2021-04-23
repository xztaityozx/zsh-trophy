package trophy

import (
	"xztaityozx/zsh-trophy/record"
)

type StrictEqual struct {
	B     string
	Title string
	Desc  string
	Grade string
}

func (s StrictEqual) Check(cmd string, _ record.Record) (Trophy, error) {
	return Trophy{
		Title:   s.Title,
		Desc:    s.Desc,
		Grade:   s.Grade,
		Cleared: s.B == cmd,
	}, nil
}
