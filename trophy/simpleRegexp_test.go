package trophy_test

import (
	"regexp"
	"testing"
	"xztaityozx/zsh-trophy/record"
	"xztaityozx/zsh-trophy/trophy"

	"github.com/stretchr/testify/assert"
)

func Test_SimpleRegexp_Check(t *testing.T) {
	as := assert.New(t)

	sr := trophy.SimpleRegexp{
		Re:    regexp.MustCompile("^cmd$"),
		Title: "Title",
		Desc:  "Desc",
		Grade: "Grade",
	}

	t.Run("マッチしないなら", func(t *testing.T) {
		res, err := sr.Check("dmc", record.Record{})
		as.Nil(err)
		as.False(res.Cleared, "falseが返ってくるべき")
	})
	t.Run("マッチするなら", func(t *testing.T) {
		res, err := sr.Check("cmd", record.Record{})
		as.Nil(err)
		as.True(res.Cleared, "trueが返ってくるべき")
		as.Equal(res.Title, "Title")
		as.Equal(res.Grade, "Grade")
		as.Equal(res.Desc, "Desc")
	})
	t.Run("Reがnilなら", func(t *testing.T) {
		_, err := trophy.SimpleRegexp{Re: nil}.Check("dmc", record.Record{})
		as.NotNil(err)
	})
}
