package trophy_test

import (
	"regexp"
	"testing"
	"xztaityozx/zsh-trophy/record"
	"xztaityozx/zsh-trophy/trophy"

	"github.com/stretchr/testify/assert"
)

func TestCount_Check(t *testing.T) {
	as := assert.New(t)

	c := trophy.Count{
		Re:    regexp.MustCompile(`\|`),
		N:     10,
		Title: "Title",
		Desc:  "Desc",
		Grade: "Grade",
	}
	t.Run("10回あるとき", func(t *testing.T) {
		res, err := c.Check("||||||||||", record.Record{})
		as.Nil(err)
		as.True(res.Cleared, "trueが返される")
	})

	t.Run("3回しかないとき", func(t *testing.T) {
		res, err := c.Check("|||", record.Record{})
		as.Nil(err)
		as.False(res.Cleared, "falseが返される")
	})
	t.Run("11回もあるとき", func(t *testing.T) {
		res, err := c.Check("|||||||||||", record.Record{})
		as.Nil(err)
		as.False(res.Cleared, "falseが返される")
	})
}
