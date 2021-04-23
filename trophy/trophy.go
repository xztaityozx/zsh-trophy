package trophy

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"xztaityozx/zsh-trophy/record"
)

const (
	Gold    string = "🥇"
	Silver  string = "🥈"
	Bronze  string = "🥉"
	Special string = "🏆"
)

type Trophy struct {
	Values  map[string]string
	Title   string
	Desc    string
	Grade   string
	Cleared bool
	Id      int
}

type ITrophy interface {
	Check(cmd string, record record.Record) (Trophy, error)
}

func (t Trophy) Print(width int) {

	fmt.Fprintf(os.Stderr, "\n%s zsh-trophy %s\n%s\n", Special, Special, strings.Repeat(".", width))
	fmt.Fprintf(os.Stderr, "..\t\033[48;2;66;66;66m\033[38;2;224;224;224m  %s %s  \033[0m\n", t.Grade, t.Title)

	sc := bufio.NewScanner(strings.NewReader(t.Desc))
	for sc.Scan() {
		fmt.Fprintf(os.Stderr, "..\t\t\033[48;2;66;66;66m\033[38;2;224;224;224m  %s  \033[0m\n", sc.Text())
	}
	fmt.Fprintf(os.Stderr, "%s\n\n", strings.Repeat(".", width))
}

var TrophyList = map[int]ITrophy{
	1:  FirstTime{Command: "ls", Comment: "lsにaliasが貼られててもいいけどね"},
	2:  NthCmd{Command: "ls", Count: 5, Key: "2::count", Grade: Bronze, Title: "5回目のls", Desc: fmt.Sprintf("lsコマンドを通算5回実行した\nとりあえずls打っちゃうことって有るよね")},
	3:  NthCmd{Command: "ls", Count: 50, Key: "3::count", Grade: Silver, Title: "50回目のls", Desc: fmt.Sprintf("lsコマンドを通算50回実行した\nslコマンドと間違えたりしないよね？")},
	4:  NthCmd{Command: "ls", Count: 100, Key: "4::count", Grade: Gold, Title: "lsマスター", Desc: fmt.Sprintf("lsコマンドを通算100回実行した\nおめでとうキミこそlsマスターだ")},
	5:  NthCmd{Command: "ls", Count: 200, Key: "4::count", Grade: Gold, Title: "キーボードにlsキーを作ろうと思います", Desc: fmt.Sprintf("lsコマンドを通算200回実行した\nワンタッチで入力されるlsキーが欲しい")},
	6:  NthCmd{Command: "ls", Count: 1000, Key: "4::count", Grade: Special, Title: "実質lsコマンド", Desc: fmt.Sprintf("lsコマンドを通算1000回実行した\nもうお前がlsでいいよ")},
	7:  FirstTime{Command: "cd", Comment: "まずはcdコマンドだよね"},
	8:  NthCmd{Command: "cd", Count: 5, Key: "8::count", Grade: Bronze, Title: "5回目のcd", Desc: fmt.Sprintf("cdコマンドを通算5回実行した\n移動にはもう慣れたかな？")},
	9:  NthCmd{Command: "cd", Count: 50, Key: "9::count", Grade: Silver, Title: "50回目のcd", Desc: fmt.Sprintf("cdコマンドを通算50回実行した")},
	10: NthCmd{Command: "cd", Count: 100, Key: "10::count", Grade: Gold, Title: "cdマスター", Desc: fmt.Sprintf("cdコマンドを通算100回実行した\nおめでとうキミこそcdマスターだ")},
	11: SimpleRegexp{Re: regexp.MustCompile(`xargs .*-P\d+`), Title: "へいれつ！", Desc: fmt.Sprintf("xargsの-Pオプションを使った\n並列！並列！", Grade: Silver)},
}
