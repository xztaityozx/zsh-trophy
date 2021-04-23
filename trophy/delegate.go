package trophy

import "xztaityozx/zsh-trophy/record"

type Delegate struct {
	Del   func(cmd string, r record.Record) bool
	Title string
	Desc  string
	Grade string
}

func (d Delegate) Check(cmd string, r record.Record) (Trophy, error) {
	return Trophy{
		Cleared: d.Del(cmd, r),
		Title:   d.Title,
		Grade:   d.Grade,
		Desc:    d.Desc,
	}, nil
}
