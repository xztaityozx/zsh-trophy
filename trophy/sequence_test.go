package trophy_test

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"xztaityozx/zsh-trophy/record"
	"xztaityozx/zsh-trophy/trophy"
)

func TestSequence_Check(t *testing.T) {
	as := assert.New(t)
	s := trophy.Sequence{Seq: []string{"a", "b", "c"}, Title: "Title", Grade: "Grade", Desc: "Desc", Id: 0}

	t.Run("1回目はa", func(t *testing.T) {
		res, err := s.Check("a", record.Record{})
		as.Nil(err)
		as.False(res.Cleared, "falseが返ってくるべき")
		as.Equal("a", res.Values["0::seq"], "履歴が正しい")
	})

	t.Run("2回目はb", func(t *testing.T) {
		res, err := s.Check("b", record.Record{Args: map[string]string{"0::seq": "a"}})
		as.Nil(err)
		as.False(res.Cleared, "falseが返ってくるべき")
		as.Equal("a,b", res.Values["0::seq"], "履歴が正しい")
	})

	t.Run("3回目はc", func(t *testing.T) {
		res, err := s.Check("c", record.Record{Args: map[string]string{"0::seq": "a,b"}})
		as.Nil(err)
		as.True(res.Cleared, "trueが返ってくるべき")
		as.Equal("a,b,c", res.Values["0::seq"], "履歴が正しい")
	})

	t.Run("4回目はd", func(t *testing.T) {
		res, err := s.Check("d", record.Record{Args: map[string]string{"0::seq": "a,b,c"}})
		as.Nil(err)
		as.False(res.Cleared, "falseが返ってくるべき")
		as.Equal("b,c,d", res.Values["0::seq"], "履歴が正しい")
	})
}

func TestOneChanceSequence_Check(t *testing.T) {
	as := assert.New(t)
	s := trophy.OneChanceSequence{
		Sequence: trophy.Sequence{
			Seq:   []string{"a", "b", "c"},
			Id:    0,
			Title: "Title",
			Grade: "Grade",
			Desc:  "Desc",
		},
		Re: regexp.MustCompile("^d HOME"),
	}

	t.Run("1回目はa", func(t *testing.T) {
		res, err := s.Check("a", record.Record{})
		as.Nil(err)
		as.False(res.Cleared, "falseが返ってくるべき")
		as.Equal("a", res.Values["0::seq"], "履歴が正しい")
	})

	t.Run("2回目はb", func(t *testing.T) {
		res, err := s.Check("b", record.Record{Args: map[string]string{"0::seq": "a"}})
		as.Nil(err)
		as.False(res.Cleared, "falseが返ってくるべき")
		as.Equal("a,b", res.Values["0::seq"], "履歴が正しい")
	})

	t.Run("3回目はc", func(t *testing.T) {
		res, err := s.Check("c", record.Record{Args: map[string]string{"0::seq": "a,b"}})
		as.Nil(err)
		as.False(res.Cleared, "falseが返ってくるべき")
		as.Equal("a,b,c", res.Values["0::seq"], "履歴が正しい")
	})

	t.Run("4回目がdでneverがnil", func(t *testing.T) {
		res, err := s.Check("d HOME", record.Record{Args: map[string]string{"0::seq": "a,b,c"}})
		as.Nil(err)
		as.True(res.Cleared, "trueが返ってくるべき")
		as.Equal("b,c,d", res.Values["0::seq"], "履歴が正しい")
	})
	t.Run("4回目がdでneverがfalse", func(t *testing.T) {
		res, err := s.Check("d HOME", record.Record{Args: map[string]string{"0::seq": "a,b,c", "0::never": "false"}})
		as.Nil(err)
		as.True(res.Cleared, "trueが返ってくるべき")
		as.Equal("b,c,d", res.Values["0::seq"], "履歴が正しい")
	})
	t.Run("4回目がdでneverがtrue", func(t *testing.T) {
		res, err := s.Check("d HOME", record.Record{Args: map[string]string{"0::seq": "a,b,c", "0::never": "true"}})
		as.Nil(err)
		as.False(res.Cleared, "falseが返ってくるべき")
		as.Equal("", res.Values["0::seq"], "履歴が正しい")
	})

	t.Run("4回目がe", func(t *testing.T) {
		res, err := s.Check("e", record.Record{Args: map[string]string{"0::seq": "a,b,c"}})
		as.Nil(err)
		as.False(res.Cleared, "falseが返ってくるべき")
		as.Equal("b,c,e", res.Values["0::seq"], "履歴が正しい")
		as.Equal("true", res.Values["0::never"], "二度とクリアできない")
	})

}
