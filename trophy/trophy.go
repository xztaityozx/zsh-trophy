package trophy

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
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

func GenerateTrophyList(ztd string) map[int]ITrophy {

	echoUnkoRegexp := regexp.MustCompile("^echo unko$")
	pipeRegexp := regexp.MustCompile(`\|`)

	return map[int]ITrophy{
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
		11: SimpleRegexp{Re: regexp.MustCompile(`xargs .*-P\d+`), Title: "へいれつ！", Desc: fmt.Sprintf("xargsの-Pオプションを使った\n並列！並列！"), Grade: Silver},
		12: SimpleRegexp{Re: regexp.MustCompile(`xargs .*-P0`), Title: "フルパワー！", Desc: fmt.Sprintf("xargsの-P0オプションを使った"), Grade: Silver},
		13: SimpleRegexp{Re: echoUnkoRegexp, Title: "💩", Desc: fmt.Sprintf("echo unkoした\nまぁやりがちだよね"), Grade: Bronze},
		14: NthRegexp{Re: echoUnkoRegexp, Title: strings.Repeat("💩", 5), Desc: fmt.Sprint("echo unkoを累計5回実行した\nunko、好きなんだね"), Count: 5, Key: "14::count", Grade: Bronze},
		15: NthRegexp{Re: echoUnkoRegexp, Title: strings.Repeat("💩", 10), Desc: fmt.Sprint("echo unkoを累計10回実行した\nもうhogeの代わりに使ってるでしょ"), Count: 10, Key: "15::count", Grade: Bronze},
		16: SimpleRegexp{Re: regexp.MustCompile(`^unko.shout`), Title: "叫べ、💩", Desc: fmt.Sprintf("unko.shoutした\nうんこは叫ばないと思うけどね"), Grade: Silver},
		17: SimpleRegexp{Re: regexp.MustCompile(`^echo-sd`), Title: "突然の死", Desc: fmt.Sprintf("echo-sdした\n死なないで！！！"), Grade: Silver},
		18: NthRegexp{Re: echoUnkoRegexp, Title: "unkoマスター", Desc: fmt.Sprint("echo unkoを累計100回実行した\nこんなの達成してる場合かよ"), Count: 100, Key: "16::count", Grade: Special},
		19: Sequence{Seq: []string{"cd", "ls"}, Title: "あれはどこかな？", Grade: Bronze, Desc: fmt.Sprintf("cdした後lsした\n基本行動だよね"), Id: 19},
		20: Sequence{Seq: []string{"cd", "cd", "cd"}, Title: "あれ･･･？どこだここは･･･？", Grade: Bronze, Desc: fmt.Sprintf("3回連続cdした\nいまどこにいるかわかってますか？"), Id: 20},
		21: Sequence{
			Seq:   []string{"cd", "cd", "cd", "cd", "cd", "cd", "cd", "cd", "cd", "cd"},
			Title: "完全に迷子",
			Grade: Silver,
			Desc:  fmt.Sprintf("10回連続cdした\n一旦HOMEに戻ってはどうですか？"),
			Id:    21,
		},
		22: SimpleRegexp{Re: regexp.MustCompile(`cd +\$HOME`), Title: "家", Grade: Bronze, Desc: fmt.Sprintf("cd $HOMEした\nおうちが一番！")},
		23: OneChanceSequence{
			Sequence: Sequence{
				Seq:   []string{"cd", "cd", "cd", "cd", "cd", "cd", "cd", "cd", "cd", "cd"},
				Title: "おかえり",
				Grade: Gold,
				Desc:  fmt.Sprintf("10回連続cdしたあとcd $HOMEした\n素直！"),
				Id:    23,
			},
			Re: regexp.MustCompile(`cd \$HOME`),
		},
		24: CatCheck{
			Title:     "zsh-trophyってなんだよ",
			FileName:  filepath.Join(ztd, "bin", "zt.zsh"),
			Desc:      fmt.Sprintf("zt.zshをcatした\nそんな珍しいことは書いてないよ"),
			Grade:     Bronze,
			OneChance: false,
		},
		25: CatCheck{
			Title:     "🗿",
			FileName:  filepath.Join(ztd, ".secret", "moai"),
			Desc:      fmt.Sprintf("moaiをcatした\nなんでモアイ？"),
			Grade:     Bronze,
			OneChance: false,
		},
		26: CatCheck{
			Title:     "ワンちゃん",
			FileName:  filepath.Join(ztd, ".secret", "wan-chan"),
			Desc:      fmt.Sprintf("なによりも先にwan-chanをcatした\n良く気付いたね"),
			Grade:     Silver,
			OneChance: true,
		},
		27: SimpleRegexp{Re: regexp.MustCompile("^git +.*init"), Grade: Bronze, Title: "ここをリポジトリとする！", Desc: "はじめてgit initした"},
		28: SimpleRegexp{Re: regexp.MustCompile("^git +.*add"), Grade: Bronze, Title: "これ追加しときます", Desc: "はじめてgit addした"},
		29: SimpleRegexp{Re: regexp.MustCompile("^git +.*commit"), Grade: Bronze, Title: "や・く・そ・く", Desc: "はじめてgit commitした"},
		30: SimpleRegexp{Re: regexp.MustCompile("^git +.*push"), Grade: Bronze, Title: "行ってらっしゃい", Desc: fmt.Sprintf("はじめてgit pushした\npushし忘れると悲しいこと、なるよね")},
		31: SimpleRegexp{Re: regexp.MustCompile("^git +.*pull"), Grade: Bronze, Title: "いらっしゃい", Desc: fmt.Sprintf("はじめてgit pullした\nあ、あれ･･･pullし忘れたぞ･･･")},
		32: SimpleRegexp{Re: regexp.MustCompile("^git +.*((-f|--force)?.+ push|push .*(-f|--force))"), Grade: Silver, Title: "歴史すら思いのままさ", Desc: fmt.Sprintf("はじめてgit push -fした\n無かったことにできる。そうgit push -fならね")},
		33: SimpleRegexp{Re: pipeRegexp, Grade: Bronze, Title: "はじめてのパイプ", Desc: "はじめてパイプ(|)を使ったコマンドを実行した"},
		34: Count{Re: pipeRegexp, N: 4, Title: "ワンライナーはじめました", Desc: "パイプ(|)が4つ含まれるコマンドを実行した", Grade: Silver},
		35: Count{Re: pipeRegexp, N: 10, Title: "繋げ過ぎたな･･･", Desc: "パイプ(|)が10つ含まれるコマンドを実行した", Grade: Silver},
		36: SimpleRegexp{Re: regexp.MustCompile(`^echo \|{10}$`), Title: "ずるいぞ", Desc: fmt.Sprintf("「繋げ過ぎたな･･･」を楽にクリアしようとしてecho ||||||||||した\nしてない？ホント？まぁいいけどね"), Grade: Silver},
		37: FirstTime{Command: "find", Comment: "オプションが覚えられません！！！！"},
		38: NthCmd{
			Command: "find",
			Key:     "38::count",
			Grade:   Silver,
			Title:   "findマスター",
			Desc:    fmt.Sprintf("findコマンドを通算50回実行した"),
			Count:   50,
		},
		39: NthCmd{
			Command: "cd",
			Key:     "39::count",
			Grade:   Special,
			Title:   "ファイル探して3000ディレクトリ",
			Desc:    fmt.Sprintf("cdコマンドを通算3000回実行した\nどのディレクトリに一番訪れたかな？"),
			Count:   3000,
		},
		40: Sequence{
			Seq: []string{
				"unko.king", "unko.shout",
			},
			Title: "ウンコンボ(初級)",
			Desc:  "ウンコンボ(初級)を発動した\nようはウンコ使い初級ってことだな",
			Grade: Bronze,
			Id:    40,
		},
		41: Sequence{
			Seq: []string{
				"unko.king", "unko.shout", "unko.shout", "bigunko.show",
			},
			Title: "ウンコンボ(中級)",
			Desc:  "ウンコンボ(中級)を発動した\nウンコの扱いがそこそこ上手くなってきたな",
			Grade: Silver,
			Id:    41,
		},
		42: Sequence{
			Seq: []string{
				"unko.king", "unko.shout", "unko.shout", "unko.think", "unko.pyramid", "unko.tower", "unko.ls", "super_unko",
			},
			Title: "ウンコンボ(上級)",
			Desc:  "ウンコンボ(上級)を発動した\n素晴らしい鍛錬だ",
			Grade: Gold,
			Id:    42,
		},
		43: Sequence{
			Seq: []string{
				"unko.king", "unko.shout",
				"unko.king", "unko.shout", "unko.shout", "bigunko.show",
				"unko.king", "unko.shout", "unko.shout", "unko.think", "unko.pyramid", "unko.tower", "unko.ls", "super_unko",
			},
			Title: "ウンコンボマスター",
			Desc:  "ウンコンボを初級~上級の順番で発動した\nフルコンボ！",
			Grade: Special,
			Id:    43,
		},

		44: FirstTime{Command: "rm", Comment: "削除するお！"},
		45: NthCmd{Command: "rm", Title: "5回目のrm", Desc: fmt.Sprintf("rmコマンドを通算5回実行した"), Grade: Bronze, Key: "45::count"},
		46: NthCmd{Command: "rm", Title: "10回目のrm", Desc: fmt.Sprintf("rmコマンドを通算10回実行した"), Grade: Bronze, Key: "46::count"},
		47: NthCmd{Command: "rm", Title: "50回目のrm", Desc: fmt.Sprintf("rmコマンドを通算50回実行した"), Grade: Silver, Key: "47::count"},
		48: NthCmd{Command: "rm", Title: "100回目のrm", Desc: fmt.Sprintf("rmコマンドを通算100回実行した"), Grade: Gold, Key: "48::count"},
		49: NthCmd{Command: "rm", Title: "1000回目のrm", Desc: fmt.Sprintf("rmコマンドを通算1000回実行した"), Grade: Special, Key: "49::count"},
		51: NthCmd{Command: "mv", Title: "5回目のmv", Desc: fmt.Sprintf("mvコマンドを通算5回実行した"), Grade: Bronze, Key: "51::count"},
		52: NthCmd{Command: "mv", Title: "10回目のmv", Desc: fmt.Sprintf("mvコマンドを通算10回実行した"), Grade: Bronze, Key: "52::count"},
		53: NthCmd{Command: "mv", Title: "50回目のmv", Desc: fmt.Sprintf("mvコマンドを通算50回実行した"), Grade: Silver, Key: "53::count"},
		54: NthCmd{Command: "mv", Title: "100回目のmv", Desc: fmt.Sprintf("mvコマンドを通算100回実行した"), Grade: Gold, Key: "54::count"},
		55: NthCmd{Command: "mv", Title: "1000回目のmv", Desc: fmt.Sprintf("mvコマンドを通算1000回実行した"), Grade: Special, Key: "55::count"},
		56: NthCmd{Command: "date", Title: "5回目のdate", Desc: fmt.Sprintf("dateコマンドを通算5回実行した"), Grade: Bronze, Key: "56::count"},
		57: NthCmd{Command: "date", Title: "10回目のdate", Desc: fmt.Sprintf("dateコマンドを通算10回実行した"), Grade: Bronze, Key: "57::count"},
		58: NthCmd{Command: "date", Title: "50回目のdate", Desc: fmt.Sprintf("dateコマンドを通算50回実行した"), Grade: Silver, Key: "58::count"},
		59: NthCmd{Command: "date", Title: "100回目のdate", Desc: fmt.Sprintf("dateコマンドを通算100回実行した"), Grade: Gold, Key: "59::count"},
		60: NthCmd{Command: "date", Title: "1000回目のdate", Desc: fmt.Sprintf("dateコマンドを通算1000回実行した"), Grade: Special, Key: "60::count"},
		61: Progress{N: 10, Grade: Bronze, Comment: "まだまだあるぞ"},
		62: Progress{N: 20, Grade: Silver, Comment: "その調子"},
		63: Progress{N: 50, Grade: Gold, Comment: "すごいじゃん"},
		64: Progress{N: 64, Grade: Gold, Comment: "どうでもいい情報だけどこのトロフィーの内部IDも64なんだよ"},
		65: RecordCheck{Ztd: ztd},
		66: SimpleRegexp{Re: regexp.MustCompile(`date .?\+%s.?`), Title: "2038", Desc: fmt.Sprintf("dateコマンドでunixtimeを出力した\n尽きる日のことを思うと･･･"), Grade: Silver},
		67: StrictEqual{B: "dir", Title: "DOS窓", Desc: fmt.Sprintf("dirしようとした\n黒い画面だもんね"), Grade: Bronze},
		68: StrictEqual{B: "ipconfig", Title: "IPアドレスはっと･･･", Desc: fmt.Sprintf("ipconfigしようとした\nこのコマンドでドヤ顔したことない？あるでしょ？"), Grade: Bronze},
		69: SimpleRegexp{Re: regexp.MustCompile(`^powershell\.exe`), Title: "powershell.exe", Grade: Silver, Desc: fmt.Sprintf("PowerShellを起動した\nもしかしてwsl内に居るんじゃないですか？")},
		70: SimpleRegexp{Re: regexp.MustCompile(`^pwsh`), Title: "pwsh", Grade: Bronze, Desc: fmt.Sprintf("PowerShellを起動した\npwshってなんて読むの？")},
		71: SimpleRegexp{Re: regexp.MustCompile(`^zsh`), Title: "zsh", Grade: Bronze, Desc: fmt.Sprintf("zshを起動した\nzshからzshを起動！ウィーン！！")},
		72: SimpleRegexp{Re: regexp.MustCompile(`^bash`), Title: "bash", Grade: Bronze, Desc: fmt.Sprintf("bashを起動した\nbashもいいよね")},
		73: SimpleRegexp{Re: regexp.MustCompile(`^dash`), Title: "dash", Grade: Bronze, Desc: fmt.Sprintf("dashを起動した\nなんかdashって常に走ってそう")},
		74: SimpleRegexp{Re: regexp.MustCompile(`^sh`), Title: "sh", Grade: Bronze, Desc: fmt.Sprintf("shを起動した\nｼｭ")},
		75: SimpleRegexp{Re: regexp.MustCompile(`^fish`), Title: "fish", Grade: Bronze, Desc: fmt.Sprintf("fishを起動した\n🐡🐡🐡🐡🐡🐡🐡🐡🐡🐡🐡🐡🐡🐡🐡")},
		76: SimpleRegexp{Re: regexp.MustCompile(`^tcsh`), Title: "tcsh", Grade: Bronze, Desc: fmt.Sprintf("tcshを起動した")},
		77: SimpleRegexp{Re: regexp.MustCompile(`^csh`), Title: "csh", Grade: Bronze, Desc: fmt.Sprintf("cshを起動した")},
		78: FirstTime{Command: "docker", Comment: "ドカドカドッカーｗ"},
		79: FirstTime{Command: "ssh", Comment: "何の略か知ってるかしら？"},
		80: EnvLookup{Name: "TMUX", Desc: fmt.Sprintf("はじめてtmuxを起動した\nいつもと違う環境。いつもと違うプリフィックスキー"), Title: "tmux", Grade: Bronze},
		81: EnvValidate{
			Validator: func(val string) bool {
				home, _ := homedir.Dir()
				return val != home
			},
			EnvLookup: EnvLookup{
				Name:  "HOME",
				Title: "アンタ、誰？",
				Desc:  "zt-trophyが取得したHOMEと$HOMEが一致しなかった",
				Grade: Gold,
			},
		},
		82: EnvValidate{
			Validator: func(val string) bool {
				a, _ := filepath.Abs(val)
				b, _ := filepath.Abs(filepath.Join(ztd, ".record", "record.json"))
				return a != b
			},
			EnvLookup: EnvLookup{
				Name:  "ZT_RECORD",
				Title: "別の記録",
				Desc:  "記録ファイルがデフォルトの場所でなかった",
				Grade: Gold,
			},
		},
		83: EnvLookup{Name: "UNKO", Desc: "UNKOという環境変数が有った", Title: "💩 on env", Grade: Silver},
		84: EnvLookup{Name: "SUPER_UNKO", Desc: "SUPER_UNKOという環境変数が有った\nsuper_unkoへのコントリビュート、お待ちしております", Title: "super_unkoへの招待", Grade: Silver},
		85: EnvLookup{Name: "ZDOTDIR", Desc: "ZDOTDIRが設定されていた\nそうそう場所は変えられるんだよね", Title: ".zshrc「俺はここだ！！」", Grade: Silver},
		86: SimpleRegexp{
			Re:    regexp.MustCompile("^zcompile"),
			Grade: Silver,
			Title: "コンパイル",
			Desc:  "zcompileを実行した\nzwcってなんだよ",
		},
		87: FirstTime{Comment: "alias", Command: "使ってないaliasあるんじゃない？"},
		88: EnvValidate{
			Validator: func(val string) bool {
				return len(val) == 0
			},
			EnvLookup: EnvLookup{
				Name:  "PATH",
				Title: "なんにもないよ",
				Desc:  "PATHを空文字にしてしまった\nあらあらやっちまったねぇ",
				Grade: Gold,
			},
		},
		89: SimpleRegexp{Re: regexp.MustCompile("^apt"), Title: "apt", Desc: "aptを実行した", Grade: Bronze},
		90: SimpleRegexp{Re: regexp.MustCompile("^brew"), Title: "brew", Desc: "brewを実行した", Grade: Bronze},
		91: SimpleRegexp{Re: regexp.MustCompile("^pacman"), Title: "pacman", Desc: "pacmanを実行した", Grade: Bronze},
		92: SimpleRegexp{Re: regexp.MustCompile("^yum"), Title: "yum", Desc: "yumを実行した", Grade: Bronze},
		93: SimpleRegexp{Re: regexp.MustCompile("^aptitude"), Title: "aptitude", Desc: "aptitudeを実行した", Grade: Bronze},
		94: SimpleRegexp{Re: regexp.MustCompile("^rpm"), Title: "rpm", Desc: "rpmを実行した", Grade: Bronze},
		95: SimpleRegexp{Re: regexp.MustCompile("^pushd"), Title: "浮気か！？", Desc: "pushdを使った\ncdが君を見ているぞ", Grade: Silver},
		96: SimpleRegexp{Re: regexp.MustCompile("^popd"), Title: "戻りまーす", Desc: "popdを使った\nこれを使わないとpushdの意味がない", Grade: Silver},
		97: CatCheck{
			Title:     "読んでくれてありがとう",
			FileName:  filepath.Join(ztd, "README.md"),
			Desc:      "READMEをcatした\nうれしいからspecialなやつをあげちゃう",
			Grade:     Special,
			OneChance: false,
		},
		98: FirstTime{Command: "curl", Comment: "wgetってのもあるんだぜ"},
		99: FirstTime{Command: "wget", Comment: "curlってのもあるんだぜ"},
		100: Progress{
			N:       99,
			Grade:   Special,
			Comment: "このトロフィーで100個目！おめでとうだね！",
		},
	}
}
