package trophy

import (
	"errors"
	"os"
	"xztaityozx/zsh-trophy/record"
)

type EnvLookup struct {
	Name  string
	Title string
	Desc  string
	Grade string
}

func (e EnvLookup) Check(string, record.Record) (Trophy, error) {
	_, ok := os.LookupEnv(e.Name)

	return Trophy{
		Cleared: ok,
		Title:   e.Title,
		Grade:   e.Grade,
		Desc:    e.Desc,
	}, nil
}

type EnvValidate struct {
	Validator func(val string) bool
	EnvLookup
}

func (e EnvValidate) Check(string, record.Record) (Trophy, error) {
	if e.Validator == nil {
		return Trophy{}, errors.New("validator is nil")
	}

	val, ok := os.LookupEnv(e.Name)
	if !ok {
		return Trophy{
			Cleared: false,
		}, nil
	}

	return Trophy{
		Cleared: e.Validator(val),
		Title:   e.Title,
		Grade:   e.Grade,
		Desc:    e.Desc,
	}, nil
}
