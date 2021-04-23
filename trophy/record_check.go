package trophy

import (
	"strconv"
	"xztaityozx/zsh-trophy/record"
)

type RecordCheck struct {
	Ztd string
}

func (r RecordCheck) Check(_ string, record record.Record) (Trophy, error) {
	str, ok := record.Args["progress"]
	if !ok {
		return Trophy{Cleared: false}, nil
	}

	t := Trophy{
		Cleared: true,
		Title:   "不正してないよね･･･？",
		Desc:    "記録チェック中に変な点が有った",
		Grade:   Silver,
	}

	if val, err := strconv.Atoi(str); err != nil {
		return t, nil
	} else {
		cnt := 0
		for _, v := range record.Status {
			if v {
				cnt++
			}
		}

		t.Cleared = cnt != val
		return t, nil
	}
}
