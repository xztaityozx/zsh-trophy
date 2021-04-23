package trophy

import (
	"fmt"
	"regexp"
	"strings"
	"xztaityozx/zsh-trophy/record"
)

type FirstTime struct {
	Command string
	Comment string
}

func (ft FirstTime) Check(cmd string, _ record.Record) (Trophy, error) {
	re, err := regexp.Compile(fmt.Sprintf("^%s ?", ft.Command))
	if err != nil {
		return Trophy{}, err
	}

	return Trophy{
		Cleared: re.MatchString(cmd),
		Title:   fmt.Sprintf("はじめての%sコマンド", ft.Command),
		Desc:    strings.Join([]string{fmt.Sprintf("はじめて%sコマンドを実行した", ft.Command), ft.Comment}, "\n"),
		Grade:   Bronze,
	}, nil
}
