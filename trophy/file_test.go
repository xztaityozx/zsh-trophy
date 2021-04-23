package trophy_test

import (
	"crypto/sha256"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"xztaityozx/zsh-trophy/record"
	"xztaityozx/zsh-trophy/trophy"
)

func TestCatCheck_Check(t *testing.T) {
	as := assert.New(t)

	tmp := os.TempDir()
	basedir := filepath.Join(tmp, "zsh-trophy", "test", "cat_check")
	filename := filepath.Join(basedir, "filename")

	cc := trophy.CatCheck{
		Title:     "Title",
		FileName:  filename,
		Desc:      "Desc",
		Grade:     "Grade",
		OneChance: false,
	}

	t.Run("catじゃないなら", func(t *testing.T) {
		res, err := cc.Check("echo abc", record.Record{})
		as.Nil(err)
		as.False(res.Cleared, "falseが返される")
	})

	t.Run("catだけどfileが違うなら", func(t *testing.T) {
		res, err := cc.Check("cat abc", record.Record{})
		as.Nil(err)
		as.False(res.Cleared, "falseが返される")
	})

	t.Run("catでファイルが一致するなら", func(t *testing.T) {
		res, err := cc.Check("cat "+filename, record.Record{})
		as.Nil(err)
		as.True(res.Cleared, "trueが返される")
	})

	t.Run("OneChanceなら", func(t *testing.T) {
		cc.OneChance = true
		res, err := cc.Check("cat abc", record.Record{})
		as.Nil(err)
		as.False(res.Cleared, "falseが返される")
		nk := fmt.Sprintf("%x", sha256.Sum256([]byte(cc.Title+cc.Desc+cc.Grade+filename)))
		as.Equal("true", res.Values[nk])

		res, err = cc.Check("cat "+filename, record.Record{Args: map[string]string{nk: "true"}})
		as.Nil(err)
		as.False(res.Cleared, "falseが返される")
	})
}
