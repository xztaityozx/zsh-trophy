package trophy_test

import (
	"regexp"
	"testing"
	"xztaityozx/zsh-trophy/record"
	"xztaityozx/zsh-trophy/trophy"

	"github.com/stretchr/testify/assert"
)

func Test_NthCmd_Check(t *testing.T) {
	as := assert.New(t)

	nthCmd := trophy.NthCmd{Command: "cmd", Key: "key", Grade: "Grade", Title: "Title", Desc: "Desc", Count: 10}

	t.Run("cmdが違うなら", func(t *testing.T) {
		res, err := nthCmd.Check("dmc", record.Record{})
		as.Nil(err)
		as.False(res.Cleared, "falseなべき")
	})

	t.Run("回数が足りないなら", func(t *testing.T) {
		res, err := nthCmd.Check("cmd", record.Record{Args: map[string]string{"key": "1"}})
		as.Nil(err)
		as.False(res.Cleared, "falseなべき")
	})

	t.Run("回数が足りているなら", func(t *testing.T) {
		res, err := nthCmd.Check("cmd", record.Record{Args: map[string]string{"key": "9"}})
		as.Nil(err)
		as.True(res.Cleared, "trueなべき")
		as.Equal("Grade", res.Grade)
		as.Equal("Desc", res.Desc)
		as.Equal("Title", res.Title)
	})
}

func TestNthRegexp_Check(t *testing.T) {
	as := assert.New(t)

	nr := trophy.NthRegexp{Re: regexp.MustCompile("^echo cmd$"), Title: "Title", Grade: "Grade", Desc: "Desc", Key: "Key", Count: 10}
	r0 := record.Record{
		Args: map[string]string{
			"Key": "0",
		},
	}
	r9 := record.Record{
		Args: map[string]string{
			"Key": "9",
		},
	}

	t.Run("マッチしないとき", func(t *testing.T) {
		t.Run("Key: 0のなら", func(t *testing.T) {
			res, err := nr.Check("dmc", r0)
			as.Nil(err)
			as.False(res.Cleared, "falseが返ってくるべき")
		})
		t.Run("Key: 9なら", func(t *testing.T) {
			res, err := nr.Check("dmc", r9)
			as.Nil(err)
			as.False(res.Cleared, "falseが返ってくるべき")
		})
	})
	t.Run("マッチしないとき", func(t *testing.T) {
		t.Run("Key: 0のなら", func(t *testing.T) {
			res, err := nr.Check("echo cmd", r0)
			as.Nil(err)
			as.False(res.Cleared, "falseが返ってくるべき")
			as.Equal("Title", res.Title)
			as.Equal("Desc", res.Desc)
			as.Equal("Grade", res.Grade)
		})
		t.Run("Key: 9なら", func(t *testing.T) {
			res, err := nr.Check("echo cmd", r9)
			as.Nil(err)
			as.True(res.Cleared, "trueが返ってくるべき")
			as.Equal("Title", res.Title)
			as.Equal("Desc", res.Desc)
			as.Equal("Grade", res.Grade)
		})
	})

	t.Run("n.Reがnilなら", func(t *testing.T) {
		_, err := trophy.NthRegexp{Re: nil}.Check("cmd", record.Record{})
		as.NotNil(err)
	})
}
