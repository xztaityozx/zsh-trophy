package trophy_test

import (
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
	})
}
