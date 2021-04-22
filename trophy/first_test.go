package trophy_test

import (
	"strings"
	"testing"
	"xztaityozx/zsh-trophy/record"
	"xztaityozx/zsh-trophy/trophy"

	"github.com/stretchr/testify/assert"
)

func Test_FirstTime_Check(t *testing.T) {
	ft := trophy.FirstTime{Command: "cmd", Comment: "Comment"}
	as := assert.New(t)

	t.Run("コマンドがマッチしないなら", func(t *testing.T) {
		res, err := ft.Check("dmc", record.Record{})
		as.Nil(err)
		as.False(res.Cleared, "falseなべき")
	})

	t.Run("コマンドがマッチするなら", func(t *testing.T) {
		res, err := ft.Check("cmd", record.Record{})
		as.Nil(err)
		as.True(res.Cleared, "trueなべき")
		as.Equal(res.Title, "はじめてのcmdコマンド", "タイトルが一致するべき")
		as.Equal(res.Grade, trophy.Bronze, "グレードがBronzeであるべき")
		as.Equal(res.Desc, strings.Join([]string{"はじめてcmdコマンドを実行した", "Comment"}, "\n"), "説明が一致するべき")
	})
}
