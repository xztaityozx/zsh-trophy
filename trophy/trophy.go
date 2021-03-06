package trophy

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"xztaityozx/zsh-trophy/record"

	"github.com/mitchellh/go-homedir"
)

const (
	Gold    string = "ð¥"
	Silver  string = "ð¥"
	Bronze  string = "ð¥"
	Special string = "ð"
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
		1:  FirstTime{Command: "ls", Comment: "lsã«aliasãè²¼ããã¦ã¦ããããã©ã­"},
		2:  NthCmd{Command: "ls", Count: 5, Key: "2::count", Grade: Bronze, Title: "5åç®ã®ls", Desc: fmt.Sprintf("lsã³ãã³ããéç®5åå®è¡ãã\nã¨ããããlsæã£ã¡ãããã¨ã£ã¦æããã­")},
		3:  NthCmd{Command: "ls", Count: 50, Key: "3::count", Grade: Silver, Title: "50åç®ã®ls", Desc: fmt.Sprintf("lsã³ãã³ããéç®50åå®è¡ãã\nslã³ãã³ãã¨ééãããããªããã­ï¼")},
		4:  NthCmd{Command: "ls", Count: 100, Key: "4::count", Grade: Gold, Title: "lsãã¹ã¿ã¼", Desc: fmt.Sprintf("lsã³ãã³ããéç®100åå®è¡ãã\nããã§ã¨ãã­ãããlsãã¹ã¿ã¼ã ")},
		5:  NthCmd{Command: "ls", Count: 200, Key: "4::count", Grade: Gold, Title: "ã­ã¼ãã¼ãã«lsã­ã¼ãä½ããã¨æãã¾ã", Desc: fmt.Sprintf("lsã³ãã³ããéç®200åå®è¡ãã\nã¯ã³ã¿ããã§å¥åãããlsã­ã¼ãæ¬²ãã")},
		6:  NthCmd{Command: "ls", Count: 1000, Key: "4::count", Grade: Special, Title: "å®è³ªlsã³ãã³ã", Desc: fmt.Sprintf("lsã³ãã³ããéç®1000åå®è¡ãã\nãããåãlsã§ããã")},
		7:  FirstTime{Command: "cd", Comment: "ã¾ãã¯cdã³ãã³ãã ãã­"},
		8:  NthCmd{Command: "cd", Count: 5, Key: "8::count", Grade: Bronze, Title: "5åç®ã®cd", Desc: fmt.Sprintf("cdã³ãã³ããéç®5åå®è¡ãã\nç§»åã«ã¯ããæ£ããããªï¼")},
		9:  NthCmd{Command: "cd", Count: 50, Key: "9::count", Grade: Silver, Title: "50åç®ã®cd", Desc: fmt.Sprintf("cdã³ãã³ããéç®50åå®è¡ãã")},
		10: NthCmd{Command: "cd", Count: 100, Key: "10::count", Grade: Gold, Title: "cdãã¹ã¿ã¼", Desc: fmt.Sprintf("cdã³ãã³ããéç®100åå®è¡ãã\nããã§ã¨ãã­ãããcdãã¹ã¿ã¼ã ")},
		11: SimpleRegexp{Re: regexp.MustCompile(`xargs .*-P\d+`), Title: "ã¸ããã¤ï¼", Desc: fmt.Sprintf("xargsã®-Pãªãã·ã§ã³ãä½¿ã£ã\nä¸¦åï¼ä¸¦åï¼"), Grade: Silver},
		12: SimpleRegexp{Re: regexp.MustCompile(`xargs .*-P0`), Title: "ãã«ãã¯ã¼ï¼", Desc: fmt.Sprintf("xargsã®-P0ãªãã·ã§ã³ãä½¿ã£ã"), Grade: Silver},
		13: SimpleRegexp{Re: echoUnkoRegexp, Title: "ð©", Desc: fmt.Sprintf("echo unkoãã\nã¾ããããã¡ã ãã­"), Grade: Bronze},
		14: NthRegexp{Re: echoUnkoRegexp, Title: strings.Repeat("ð©", 5), Desc: fmt.Sprint("echo unkoãç´¯è¨5åå®è¡ãã\nunkoãå¥½ããªãã ã­"), Count: 5, Key: "14::count", Grade: Bronze},
		15: NthRegexp{Re: echoUnkoRegexp, Title: strings.Repeat("ð©", 10), Desc: fmt.Sprint("echo unkoãç´¯è¨10åå®è¡ãã\nããhogeã®ä»£ããã«ä½¿ã£ã¦ãã§ãã"), Count: 10, Key: "15::count", Grade: Bronze},
		16: SimpleRegexp{Re: regexp.MustCompile(`^unko.shout`), Title: "å«ã¹ãð©", Desc: fmt.Sprintf("unko.shoutãã\nãããã¯å«ã°ãªãã¨æããã©ã­"), Grade: Silver},
		17: SimpleRegexp{Re: regexp.MustCompile(`^echo-sd`), Title: "çªç¶ã®æ­»", Desc: fmt.Sprintf("echo-sdãã\næ­»ãªãªãã§ï¼ï¼ï¼"), Grade: Silver},
		18: NthRegexp{Re: echoUnkoRegexp, Title: "unkoãã¹ã¿ã¼", Desc: fmt.Sprint("echo unkoãç´¯è¨100åå®è¡ãã\nãããªã®éæãã¦ãå ´åãã"), Count: 100, Key: "16::count", Grade: Special},
		19: Sequence{Seq: []string{"cd", "ls"}, Title: "ããã¯ã©ãããªï¼", Grade: Bronze, Desc: fmt.Sprintf("cdããå¾lsãã\nåºæ¬è¡åã ãã­"), Id: 19},
		20: Sequence{Seq: []string{"cd", "cd", "cd"}, Title: "ããï½¥ï½¥ï½¥ï¼ã©ãã ããã¯ï½¥ï½¥ï½¥ï¼", Grade: Bronze, Desc: fmt.Sprintf("3åé£ç¶cdãã\nãã¾ã©ãã«ãããããã£ã¦ã¾ããï¼"), Id: 20},
		21: Sequence{
			Seq:   []string{"cd", "cd", "cd", "cd", "cd", "cd", "cd", "cd", "cd", "cd"},
			Title: "å®å¨ã«è¿·å­",
			Grade: Silver,
			Desc:  fmt.Sprintf("10åé£ç¶cdãã\nä¸æ¦HOMEã«æ»ã£ã¦ã¯ã©ãã§ããï¼"),
			Id:    21,
		},
		22: SimpleRegexp{Re: regexp.MustCompile(`cd +\$HOME`), Title: "å®¶", Grade: Bronze, Desc: fmt.Sprintf("cd $HOMEãã\nããã¡ãä¸çªï¼")},
		23: OneChanceSequence{
			Sequence: Sequence{
				Seq:   []string{"cd", "cd", "cd", "cd", "cd", "cd", "cd", "cd", "cd", "cd"},
				Title: "ãããã",
				Grade: Gold,
				Desc:  fmt.Sprintf("10åé£ç¶cdãããã¨cd $HOMEãã\nç´ ç´ï¼"),
				Id:    23,
			},
			Re: regexp.MustCompile(`cd \$HOME`),
		},
		24: CatCheck{
			Title:     "zsh-trophyã£ã¦ãªãã ã",
			FileName:  filepath.Join(ztd, "bin", "zt.zsh"),
			Desc:      fmt.Sprintf("zt.zshãcatãã\nãããªçãããã¨ã¯æ¸ãã¦ãªãã"),
			Grade:     Bronze,
			OneChance: false,
		},
		25: CatCheck{
			Title:     "ð¿",
			FileName:  filepath.Join(ztd, ".secret", "moai"),
			Desc:      fmt.Sprintf("moaiãcatãã\nãªãã§ã¢ã¢ã¤ï¼"),
			Grade:     Bronze,
			OneChance: false,
		},
		26: CatCheck{
			Title:     "ã¯ã³ã¡ãã",
			FileName:  filepath.Join(ztd, ".secret", "wan-chan"),
			Desc:      fmt.Sprintf("ãªã«ãããåã«wan-chanãcatãã\nè¯ãæ°ä»ããã­"),
			Grade:     Silver,
			OneChance: true,
		},
		27: SimpleRegexp{Re: regexp.MustCompile("^git +.*init"), Grade: Bronze, Title: "ããããªãã¸ããªã¨ããï¼", Desc: "ã¯ããã¦git initãã"},
		28: SimpleRegexp{Re: regexp.MustCompile("^git +.*add"), Grade: Bronze, Title: "ããè¿½å ãã¨ãã¾ã", Desc: "ã¯ããã¦git addãã"},
		29: SimpleRegexp{Re: regexp.MustCompile("^git +.*commit"), Grade: Bronze, Title: "ãã»ãã»ãã»ã", Desc: "ã¯ããã¦git commitãã"},
		30: SimpleRegexp{Re: regexp.MustCompile("^git +.*push"), Grade: Bronze, Title: "è¡ã£ã¦ãã£ããã", Desc: fmt.Sprintf("ã¯ããã¦git pushãã\npushãå¿ããã¨æ²ãããã¨ããªããã­")},
		31: SimpleRegexp{Re: regexp.MustCompile("^git +.*pull"), Grade: Bronze, Title: "ããã£ããã", Desc: fmt.Sprintf("ã¯ããã¦git pullãã\nããããï½¥ï½¥ï½¥pullãå¿ãããï½¥ï½¥ï½¥")},
		32: SimpleRegexp{Re: regexp.MustCompile("^git +.*((-f|--force)?.+ push|push .*(-f|--force))"), Grade: Silver, Title: "æ­´å²ããæãã®ã¾ã¾ã", Desc: fmt.Sprintf("ã¯ããã¦git push -fãã\nç¡ãã£ããã¨ã«ã§ãããããgit push -fãªãã­")},
		33: SimpleRegexp{Re: pipeRegexp, Grade: Bronze, Title: "ã¯ããã¦ã®ãã¤ã", Desc: "ã¯ããã¦ãã¤ã(|)ãä½¿ã£ãã³ãã³ããå®è¡ãã"},
		34: Count{Re: pipeRegexp, N: 4, Title: "ã¯ã³ã©ã¤ãã¼ã¯ããã¾ãã", Desc: "ãã¤ã(|)ã4ã¤å«ã¾ããã³ãã³ããå®è¡ãã", Grade: Silver},
		35: Count{Re: pipeRegexp, N: 10, Title: "ç¹ãéãããªï½¥ï½¥ï½¥", Desc: "ãã¤ã(|)ã10ã¤å«ã¾ããã³ãã³ããå®è¡ãã", Grade: Silver},
		36: SimpleRegexp{Re: regexp.MustCompile(`^echo \|{10}$`), Title: "ãããã", Desc: fmt.Sprintf("ãç¹ãéãããªï½¥ï½¥ï½¥ããæ¥½ã«ã¯ãªã¢ãããã¨ãã¦echo ||||||||||ãã\nãã¦ãªãï¼ãã³ãï¼ã¾ããããã©ã­"), Grade: Silver},
		37: FirstTime{Command: "find", Comment: "ãªãã·ã§ã³ãè¦ãããã¾ããï¼ï¼ï¼ï¼"},
		38: NthCmd{
			Command: "find",
			Key:     "38::count",
			Grade:   Silver,
			Title:   "findãã¹ã¿ã¼",
			Desc:    fmt.Sprintf("findã³ãã³ããéç®50åå®è¡ãã"),
			Count:   50,
		},
		39: NthCmd{
			Command: "cd",
			Key:     "39::count",
			Grade:   Special,
			Title:   "ãã¡ã¤ã«æ¢ãã¦3000ãã£ã¬ã¯ããª",
			Desc:    fmt.Sprintf("cdã³ãã³ããéç®3000åå®è¡ãã\nã©ã®ãã£ã¬ã¯ããªã«ä¸çªè¨ªããããªï¼"),
			Count:   3000,
		},
		40: Sequence{
			Seq: []string{
				"unko.king", "unko.shout",
			},
			Title: "ã¦ã³ã³ã³ã(åç´)",
			Desc:  "ã¦ã³ã³ã³ã(åç´)ãçºåãã\nããã¯ã¦ã³ã³ä½¿ãåç´ã£ã¦ãã¨ã ãª",
			Grade: Bronze,
			Id:    40,
		},
		41: Sequence{
			Seq: []string{
				"unko.king", "unko.shout", "unko.shout", "bigunko.show",
			},
			Title: "ã¦ã³ã³ã³ã(ä¸­ç´)",
			Desc:  "ã¦ã³ã³ã³ã(ä¸­ç´)ãçºåãã\nã¦ã³ã³ã®æ±ããããããä¸æããªã£ã¦ãããª",
			Grade: Silver,
			Id:    41,
		},
		42: Sequence{
			Seq: []string{
				"unko.king", "unko.shout", "unko.shout", "unko.think", "unko.pyramid", "unko.tower", "unko.ls", "super_unko",
			},
			Title: "ã¦ã³ã³ã³ã(ä¸ç´)",
			Desc:  "ã¦ã³ã³ã³ã(ä¸ç´)ãçºåãã\nç´ æ´ãããéé¬ã ",
			Grade: Gold,
			Id:    42,
		},
		43: Sequence{
			Seq: []string{
				"unko.king", "unko.shout",
				"unko.king", "unko.shout", "unko.shout", "bigunko.show",
				"unko.king", "unko.shout", "unko.shout", "unko.think", "unko.pyramid", "unko.tower", "unko.ls", "super_unko",
			},
			Title: "ã¦ã³ã³ã³ããã¹ã¿ã¼",
			Desc:  "ã¦ã³ã³ã³ããåç´~ä¸ç´ã®é çªã§çºåãã\nãã«ã³ã³ãï¼",
			Grade: Special,
			Id:    43,
		},

		44: FirstTime{Command: "rm", Comment: "åé¤ãããï¼"},
		45: NthCmd{Count: 5, Command: "rm", Title: "5åç®ã®rm", Desc: fmt.Sprintf("rmã³ãã³ããéç®5åå®è¡ãã"), Grade: Bronze, Key: "45::count"},
		46: NthCmd{Count: 10, Command: "rm", Title: "10åç®ã®rm", Desc: fmt.Sprintf("rmã³ãã³ããéç®10åå®è¡ãã"), Grade: Bronze, Key: "46::count"},
		47: NthCmd{Count: 50, Command: "rm", Title: "50åç®ã®rm", Desc: fmt.Sprintf("rmã³ãã³ããéç®50åå®è¡ãã"), Grade: Silver, Key: "47::count"},
		48: NthCmd{Count: 100, Command: "rm", Title: "100åç®ã®rm", Desc: fmt.Sprintf("rmã³ãã³ããéç®100åå®è¡ãã"), Grade: Gold, Key: "48::count"},
		49: NthCmd{Count: 100, Command: "rm", Title: "1000åç®ã®rm", Desc: fmt.Sprintf("rmã³ãã³ããéç®1000åå®è¡ãã"), Grade: Special, Key: "49::count"},
		51: NthCmd{Count: 5, Command: "mv", Title: "5åç®ã®mv", Desc: fmt.Sprintf("mvã³ãã³ããéç®5åå®è¡ãã"), Grade: Bronze, Key: "51::count"},
		52: NthCmd{Count: 10, Command: "mv", Title: "10åç®ã®mv", Desc: fmt.Sprintf("mvã³ãã³ããéç®10åå®è¡ãã"), Grade: Bronze, Key: "52::count"},
		53: NthCmd{Count: 50, Command: "mv", Title: "50åç®ã®mv", Desc: fmt.Sprintf("mvã³ãã³ããéç®50åå®è¡ãã"), Grade: Silver, Key: "53::count"},
		54: NthCmd{Count: 100, Command: "mv", Title: "100åç®ã®mv", Desc: fmt.Sprintf("mvã³ãã³ããéç®100åå®è¡ãã"), Grade: Gold, Key: "54::count"},
		55: NthCmd{Count: 1000, Command: "mv", Title: "1000åç®ã®mv", Desc: fmt.Sprintf("mvã³ãã³ããéç®1000åå®è¡ãã"), Grade: Special, Key: "55::count"},
		56: NthCmd{Count: 5, Command: "date", Title: "5åç®ã®date", Desc: fmt.Sprintf("dateã³ãã³ããéç®5åå®è¡ãã"), Grade: Bronze, Key: "56::count"},
		57: NthCmd{Count: 10, Command: "date", Title: "10åç®ã®date", Desc: fmt.Sprintf("dateã³ãã³ããéç®10åå®è¡ãã"), Grade: Bronze, Key: "57::count"},
		58: NthCmd{Count: 50, Command: "date", Title: "50åç®ã®date", Desc: fmt.Sprintf("dateã³ãã³ããéç®50åå®è¡ãã"), Grade: Silver, Key: "58::count"},
		59: NthCmd{Count: 100, Command: "date", Title: "100åç®ã®date", Desc: fmt.Sprintf("dateã³ãã³ããéç®100åå®è¡ãã"), Grade: Gold, Key: "59::count"},
		60: NthCmd{Count: 1000, Command: "date", Title: "1000åç®ã®date", Desc: fmt.Sprintf("dateã³ãã³ããéç®1000åå®è¡ãã"), Grade: Special, Key: "60::count"},
		61: Progress{N: 10, Grade: Bronze, Comment: "ã¾ã ã¾ã ããã"},
		62: Progress{N: 20, Grade: Silver, Comment: "ãã®èª¿å­"},
		63: Progress{N: 50, Grade: Gold, Comment: "ãããããã"},
		64: Progress{N: 64, Grade: Gold, Comment: "ã©ãã§ãããæå ±ã ãã©ãã®ãã­ãã£ã¼ã®åé¨IDã64ãªãã ã"},
		65: FirstTime{Command: "date", Comment: "ä»ä½æï¼"},
		66: SimpleRegexp{Re: regexp.MustCompile(`date .?\+%s.?`), Title: "2038", Desc: fmt.Sprintf("dateã³ãã³ãã§unixtimeãåºåãã\nå°½ããæ¥ã®ãã¨ãæãã¨ï½¥ï½¥ï½¥"), Grade: Silver},
		67: StrictEqual{B: "dir", Title: "DOSçª", Desc: fmt.Sprintf("dirãããã¨ãã\né»ãç»é¢ã ããã­"), Grade: Bronze},
		68: StrictEqual{B: "ipconfig", Title: "IPã¢ãã¬ã¹ã¯ã£ã¨ï½¥ï½¥ï½¥", Desc: fmt.Sprintf("ipconfigãããã¨ãã\nãã®ã³ãã³ãã§ãã¤é¡ãããã¨ãªãï¼ããã§ããï¼"), Grade: Bronze},
		69: SimpleRegexp{Re: regexp.MustCompile(`^powershell\.exe`), Title: "powershell.exe", Grade: Silver, Desc: fmt.Sprintf("PowerShellãèµ·åãã\nããããã¦wslåã«å±ãããããªãã§ããï¼")},
		70: SimpleRegexp{Re: regexp.MustCompile(`^pwsh`), Title: "pwsh", Grade: Bronze, Desc: fmt.Sprintf("PowerShellãèµ·åãã\npwshã£ã¦ãªãã¦èª­ãã®ï¼")},
		71: SimpleRegexp{Re: regexp.MustCompile(`^zsh`), Title: "zsh", Grade: Bronze, Desc: fmt.Sprintf("zshãèµ·åãã\nzshããzshãèµ·åï¼ã¦ã£ã¼ã³ï¼ï¼")},
		72: SimpleRegexp{Re: regexp.MustCompile(`^bash`), Title: "bash", Grade: Bronze, Desc: fmt.Sprintf("bashãèµ·åãã\nbashããããã­")},
		73: SimpleRegexp{Re: regexp.MustCompile(`^dash`), Title: "dash", Grade: Bronze, Desc: fmt.Sprintf("dashãèµ·åãã\nãªããdashã£ã¦å¸¸ã«èµ°ã£ã¦ãã")},
		74: SimpleRegexp{Re: regexp.MustCompile(`^sh`), Title: "sh", Grade: Bronze, Desc: fmt.Sprintf("shãèµ·åãã\nï½¼ï½­")},
		75: SimpleRegexp{Re: regexp.MustCompile(`^fish`), Title: "fish", Grade: Bronze, Desc: fmt.Sprintf("fishãèµ·åãã\nð¡ð¡ð¡ð¡ð¡ð¡ð¡ð¡ð¡ð¡ð¡ð¡ð¡ð¡ð¡")},
		76: SimpleRegexp{Re: regexp.MustCompile(`^tcsh`), Title: "tcsh", Grade: Bronze, Desc: fmt.Sprintf("tcshãèµ·åãã")},
		77: SimpleRegexp{Re: regexp.MustCompile(`^csh`), Title: "csh", Grade: Bronze, Desc: fmt.Sprintf("cshãèµ·åãã")},
		78: FirstTime{Command: "docker", Comment: "ãã«ãã«ããã«ã¼ï½"},
		79: FirstTime{Command: "ssh", Comment: "ä½ã®ç¥ãç¥ã£ã¦ããããï¼"},
		80: EnvLookup{Name: "TMUX", Desc: fmt.Sprintf("ã¯ããã¦tmuxãèµ·åãã\nãã¤ãã¨éãç°å¢ããã¤ãã¨éãããªãã£ãã¯ã¹ã­ã¼"), Title: "tmux", Grade: Bronze},
		81: EnvValidate{
			Validator: func(val string) bool {
				home, _ := homedir.Dir()
				return val != home
			},
			EnvLookup: EnvLookup{
				Name:  "HOME",
				Title: "ã¢ã³ã¿ãèª°ï¼",
				Desc:  "zt-trophyãåå¾ããHOMEã¨$HOMEãä¸è´ããªãã£ã",
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
				Title: "å¥ã®è¨é²",
				Desc:  "è¨é²ãã¡ã¤ã«ãããã©ã«ãã®å ´æã§ãªãã£ã",
				Grade: Gold,
			},
		},
		83: EnvLookup{Name: "UNKO", Desc: "UNKOã¨ããç°å¢å¤æ°ãæã£ã", Title: "ð© on env", Grade: Silver},
		84: EnvLookup{Name: "SUPER_UNKO", Desc: "SUPER_UNKOã¨ããç°å¢å¤æ°ãæã£ã\nsuper_unkoã¸ã®ã³ã³ããªãã¥ã¼ãããå¾ã¡ãã¦ããã¾ã", Title: "super_unkoã¸ã®æå¾", Grade: Silver},
		85: EnvLookup{Name: "ZDOTDIR", Desc: "ZDOTDIRãè¨­å®ããã¦ãã\nããããå ´æã¯å¤ãããããã ãã­", Title: ".zshrcãä¿ºã¯ããã ï¼ï¼ã", Grade: Silver},
		86: SimpleRegexp{
			Re:    regexp.MustCompile("^zcompile"),
			Grade: Silver,
			Title: "ã³ã³ãã¤ã«",
			Desc:  "zcompileãå®è¡ãã\nzwcã£ã¦ãªãã ã",
		},
		87: FirstTime{Comment: "alias", Command: "ä½¿ã£ã¦ãªãaliasããããããªãï¼"},
		88: EnvValidate{
			Validator: func(val string) bool {
				return len(val) == 0
			},
			EnvLookup: EnvLookup{
				Name:  "PATH",
				Title: "ãªãã«ããªãã",
				Desc:  "PATHãç©ºæå­ã«ãã¦ãã¾ã£ã\nãããããã£ã¡ã¾ã£ãã­ã",
				Grade: Gold,
			},
		},
		89: SimpleRegexp{Re: regexp.MustCompile("^apt"), Title: "apt", Desc: "aptãå®è¡ãã", Grade: Bronze},
		90: SimpleRegexp{Re: regexp.MustCompile("^brew"), Title: "brew", Desc: "brewãå®è¡ãã", Grade: Bronze},
		91: SimpleRegexp{Re: regexp.MustCompile("^pacman"), Title: "pacman", Desc: "pacmanãå®è¡ãã", Grade: Bronze},
		92: SimpleRegexp{Re: regexp.MustCompile("^yum"), Title: "yum", Desc: "yumãå®è¡ãã", Grade: Bronze},
		93: SimpleRegexp{Re: regexp.MustCompile("^aptitude"), Title: "aptitude", Desc: "aptitudeãå®è¡ãã", Grade: Bronze},
		94: SimpleRegexp{Re: regexp.MustCompile("^rpm"), Title: "rpm", Desc: "rpmãå®è¡ãã", Grade: Bronze},
		95: SimpleRegexp{Re: regexp.MustCompile("^pushd"), Title: "æµ®æ°ãï¼ï¼", Desc: "pushdãä½¿ã£ã\ncdãåãè¦ã¦ããã", Grade: Silver},
		96: SimpleRegexp{Re: regexp.MustCompile("^popd"), Title: "æ»ãã¾ã¼ã", Desc: "popdãä½¿ã£ã\nãããä½¿ããªãã¨pushdã®æå³ããªã", Grade: Silver},
		97: CatCheck{
			Title:     "èª­ãã§ããã¦ãããã¨ã",
			FileName:  filepath.Join(ztd, "README.md"),
			Desc:      "READMEãcatãã\nããããããspecialãªãã¤ãããã¡ãã",
			Grade:     Special,
			OneChance: false,
		},
		98: FirstTime{Command: "curl", Comment: "wgetã£ã¦ã®ããããã ã"},
		99: FirstTime{Command: "wget", Comment: "curlã£ã¦ã®ããããã ã"},
		100: Progress{
			N:       100,
			Grade:   Special,
			Comment: "ãã®ãã­ãã£ã¼ã§101åç®ï¼ããã§ã¨ãã ã­ï¼",
		},
		101: SimpleRegexp{Re: regexp.MustCompile("unko"), Title: "unko is here", Grade: Bronze, Desc: "unkoãå«ãã³ãã³ããå®è¡ãã"},
		102: FirstTime{Command: "grep", Comment: "grepã§ggãï¼"},
		103: FirstTime{Command: "sed", Comment: ""},
		104: FirstTime{Command: "awk", Comment: ""},
		105: FirstTime{Command: "vim", Comment: "emacs"},
		106: FirstTime{Command: "emacs", Comment: "vim"},
		107: Sequence{
			Seq:   []string{"ls", "ls"},
			Title: "ãã£ããè¦ãã§ããï¼",
			Desc:  "é£ç¶ã§lsãã\n",
			Grade: Bronze,
			Id:    107,
		},
		108: Progress{N: 108, Grade: Special, Comment: "ããã£ï¼ããããªåã¯ï¼ããã§å®å¨ã¯ãªã¢ã ï¼"},
	}
}
