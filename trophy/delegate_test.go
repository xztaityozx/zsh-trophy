package trophy_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"xztaityozx/zsh-trophy/record"
	"xztaityozx/zsh-trophy/trophy"
)

func TestDelegate_Check(t *testing.T) {
	as := assert.New(t)

	d := trophy.Delegate{
		Del: func(cmd string, r record.Record) bool {
			return cmd == "cmd"
		},
		Title: "Title",
		Grade: "Grade",
		Desc:  "Desc",
	}

	t.Run("評価結果がfalseなら", func(t *testing.T) {
		res, err := d.Check("abc", record.Record{})
		as.Nil(err)
		as.False(res.Cleared, "falseが返される")
		as.Equal("Title", res.Title)
		as.Equal("Desc", res.Desc)
		as.Equal("Grade", res.Grade)
	})
	t.Run("評価結果がtrueなら", func(t *testing.T) {
		res, err := d.Check("cmd", record.Record{})
		as.Nil(err)
		as.True(res.Cleared, "trueが返される")
		as.Equal("Title", res.Title)
		as.Equal("Desc", res.Desc)
		as.Equal("Grade", res.Grade)
	})
}
