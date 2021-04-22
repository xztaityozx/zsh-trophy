package trophy

import (
	"bufio"
	"fmt"
	"os"
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
	2: NthCmd{Command: "ls", Count: 5, Key: "2::count", Grade: Bronze, Title: "5回目のls", Desc: fmt.Sprintf("lsコマンドを通算5回実行した\nとりあえずls打っちゃうことって有るよね")},
	3: NthCmd{Command: "ls", Count: 50, Key: "3::count", Grade: Silver, Title: "50回目のls", Desc: fmt.Sprintf("lsコマンドを通算50回実行した\nslコマンドと間違えたりしないよね？")},
	4: NthCmd{Command: "ls", Count: 100, Key: "4::count", Grade: Gold, Title: "lsマスター", Desc: fmt.Sprintf("lsコマンドを通算100回実行した\nおめでとうキミこそlsマスターだ")},
	5: NthCmd{Command: "ls", Count: 200, Key: "4::count", Grade: Gold, Title: "キーボードにlsキーを作ろうと思います", Desc: fmt.Sprintf("lsコマンドを通算200回実行した\nワンタッチで入力されるlsキーが欲しい")},
	6: NthCmd{Command: "ls", Count: 1000, Key: "4::count", Grade: Special, Title: "実質lsコマンド", Desc: fmt.Sprintf("lsコマンドを通算1000回実行した\nもうお前がlsでいいよ")},
}
