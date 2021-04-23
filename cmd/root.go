package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"xztaityozx/zsh-trophy/record"
	"xztaityozx/zsh-trophy/trophy"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "zsh-tropy",
	Run: func(cmd *cobra.Command, args []string) {
		ztd, _ := cmd.PersistentFlags().GetString("ztd")
		rd, _ := cmd.PersistentFlags().GetString("rd")
		command, _ := cmd.PersistentFlags().GetString("cmd")
		width, _ := cmd.PersistentFlags().GetInt("width")

		if len(ztd) == 0 {
			log.Fatal("ztd(arg[0]) is required")
		}

		if len(rd) == 0 {
			rd = filepath.Join(ztd, ".record", "record.json")
		}

		if _, err := os.Stat(rd); err != nil {
			fp, err := os.Create(rd)
			if err != nil {
				log.Fatal(err)
			}
			fp.Close()
		}

		var record record.Record
		b, err := ioutil.ReadFile(rd)
		if err == nil {
			json.Unmarshal(b, &record)
		}

		if record.Args == nil {
			record.Args = map[string]string{}
		}

		if record.Status == nil {
			record.Status = map[int]bool{}
		}

		progress := 0
		{
			if str, ok := record.Args["progress"]; ok {
				if v, err := strconv.Atoi(str); err == nil {
					progress = v
				}
			}
		}

		for id, itrophy := range trophy.GenerateTrophyList(ztd) {
			if val, ok := record.Status[id]; ok && val {
				continue
			}

			t, err := itrophy.Check(command, record)
			if err != nil {
				continue
			}

			if t.Cleared {
				t.Print(width)
				record.Status[id] = true
				progress++
			}

			for k, v := range t.Values {
				record.Args[k] = v
			}
		}

		record.Args["progress"] = fmt.Sprint(progress)
		b, err = json.Marshal(record)
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile(rd, b, 0644)
	},
}

func init() {
	rootCmd.PersistentFlags().Int("width", 10, "")
	rootCmd.PersistentFlags().String("ztd", "", "")
	rootCmd.PersistentFlags().String("rd", "", "")
	rootCmd.PersistentFlags().String("cmd", "", "")
}

func Execute() {
	rootCmd.Execute()
}
