package trophy_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"xztaityozx/zsh-trophy/record"
	"xztaityozx/zsh-trophy/trophy"
)

func TestEnvLookup_Check(t *testing.T) {
	as := assert.New(t)

	el := trophy.EnvLookup{
		Name:  "ZT_TEST_LOOK_UP",
		Title: "Title",
		Desc:  "Desc",
		Grade: "Grade",
	}

	t.Run("Envにない場合", func(t *testing.T) {
		res, err := el.Check("", record.Record{})
		as.Nil(err)
		as.False(res.Cleared, "falseが返される")
	})

	t.Run("Envにある場合", func(t *testing.T) {
		os.Setenv("ZT_TEST_LOOK_UP", "zt_test_look_up")
		res, err := el.Check("", record.Record{})
		as.Nil(err)
		as.True(res.Cleared, "trueが返される")
		as.Equal("Title", res.Title)
		as.Equal("Grade", res.Grade)
		as.Equal("Desc", res.Desc)
	})
}

func TestEnvValidate_Check(t *testing.T) {
	as := assert.New(t)
	k := "ZT_TEST_ENV_VALIDATE"
	ev := trophy.EnvValidate{
		Validator: func(val string) bool {
			return val == "ABC"
		},
		EnvLookup: trophy.EnvLookup{
			Name:  k,
			Title: "Title",
			Desc:  "Desc",
			Grade: "Grade",
		},
	}

	t.Run("バリデート成功なら", func(t *testing.T) {
		os.Setenv(k, "ABC")
		res, err := ev.Check("", record.Record{})
		as.Nil(err)
		as.True(res.Cleared, "trueが返される")
		as.Equal("Title", res.Title)
		as.Equal("Grade", res.Grade)
		as.Equal("Desc", res.Desc)
	})
	t.Run("バリデート失敗なら", func(t *testing.T) {
		os.Setenv(k, "DEF")
		res, err := ev.Check("", record.Record{})
		as.Nil(err)
		as.False(res.Cleared, "trueが返される")
	})
	t.Run("バリデーターがnilなら", func(t *testing.T) {
		os.Setenv(k, "ABC")
		ev.Validator = nil
		_, err := ev.Check("", record.Record{})
		as.NotNil(err)
	})
}
