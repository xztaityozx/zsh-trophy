package trophy

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"xztaityozx/zsh-trophy/record"
)

type Sequence struct {
	Seq   []string
	Title string
	Desc  string
	Grade string
	Id    int
}

func (s Sequence) Check(cmd string, r record.Record) (Trophy, error) {
	if len(s.Seq) == 0 {
		return Trophy{Cleared: false}, errors.New("seq size is 0")
	}

	sep := ","
	key := fmt.Sprint(s.Id, "::seq")
	seq, ok := r.Args[key]
	var hist []string
	if ok {
		hist = strings.Split(seq, sep)
	} else {
		hist = []string{}
	}

	if len(cmd) == 0 {
		hist = append(hist, "")
	} else {
		hist = append(hist, strings.Split(cmd, " ")[0])
	}

	if len(hist) == len(s.Seq)+1 {
		hist = hist[1:]
	}

	joined := strings.Join(hist, sep)
	return Trophy{
		Cleared: joined == strings.Join(s.Seq, sep),
		Title:   s.Title,
		Grade:   s.Grade,
		Desc:    s.Desc,
		Values:  map[string]string{key: joined},
	}, nil
}

type OneChanceSequence struct {
	Sequence
	Re *regexp.Regexp
}

func (ocs OneChanceSequence) Check(cmd string, r record.Record) (Trophy, error) {
	nk := fmt.Sprint(ocs.Id, "::never")
	never, ok := r.Args[nk]
	if ok && never == "true" {
		return Trophy{Cleared: false}, nil
	}

	if len(ocs.Seq) == 0 {
		return Trophy{}, errors.New("seq size is 0")
	}

	k := fmt.Sprint(ocs.Id, "::seq")
	sep := ","
	seq, ok := r.Args[k]
	var hist []string
	if ok {
		hist = strings.Split(seq, sep)
	} else {
		hist = []string{}
	}

	ok = false
	if len(hist) == len(ocs.Seq) {
		ok = true
		for i, v := range hist {
			if v != ocs.Seq[i] {
				ok = false
			}
		}
	}

	matched := ocs.Re.MatchString(cmd)

	if ok && !matched {
		never = "true"
	} else {
		never = "false"
	}

	if len(cmd) == 0 {
		hist = append(hist, "")
	} else {
		hist = append(hist, strings.Split(cmd, " ")[0])
	}

	if len(hist) == len(ocs.Seq)+1 {
		hist = hist[1:]
	}

	return Trophy{
		Cleared: ok && matched,
		Title:   ocs.Title,
		Grade:   ocs.Grade,
		Desc:    ocs.Desc,
		Values:  map[string]string{nk: never, k: strings.Join(hist, sep)},
	}, nil
}
