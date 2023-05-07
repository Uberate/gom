package cmds

import (
	"github.com/spf13/cobra"
	"github.com/uberate/gom/pkg/regexp_trans"
	"time"
)

func RegexCmd() *cobra.Command {
	var seed *int64
	defaultSeed := int64(0)
	seed = &defaultSeed
	rc := 10
	cmd := &cobra.Command{
		Use:   "regex",
		Short: "generate string by regex",
		Long: "Generate string by regex in standard go-ve2(https://golang.org/s/re2syntax). Default '*' max value " +
			"is '10'.",
		Example: `1. Generate value by regex: "test(([a-z]{3})|([0-9]*))"

  > gom regex "test(([a-z]{3})|([0-9]*))"
  |
  = res example: test15569
---

2. Generate value by regex: "test(([a-z]{3})|([0-9]*))", and set seed to 12356789

  > gom regex "test(([a-z]{3})|([0-9]*))" -s 12356789
  |
  = res example: test147197037`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			regexValue := args[0]

			rg := regexp_trans.NewGenerator(regexp_trans.SetSeed(*seed), regexp_trans.SetDefaultMaxRepeatCount(rc))
			res, err := rg.Generate(regexValue)
			if err != nil {
				return err
			}
			cmd.Println(res)

			return nil
		},
	}

	cmd.Flags().Int64VarP(seed, "seed", "s", time.Now().UnixNano(),
		"set the regex generate seed, default is 'now nanoseconds'")
	cmd.Flags().IntVarP(&rc, "repeat-count", "c", 10,
		"set the max repeat count for '*', but repeat count like '{11} not limit by this flag")

	return cmd
}
