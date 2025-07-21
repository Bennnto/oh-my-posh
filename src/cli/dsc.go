package cli

import (
	"fmt"

	"github.com/jandedobbeleer/oh-my-posh/src/cli/dsc"
	"github.com/jandedobbeleer/oh-my-posh/src/runtime"

	"github.com/spf13/cobra"
)

var dscCmd = &cobra.Command{
	Use:   "dsc [set|get|schema|export]",
	Short: "Desired State Configuration",
	Long: `Desired State Configuration.

Available commands:

- set: Set the desired state configuration
- get: Get the desired state configuration
- schema: Get the desired state configuration schema
- export: Export the desired state configuration`,
	ValidArgs: []string{
		"set",
		"get",
		"schema",
		"export",
	},
	Args: NoArgsOrOneValidArg,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}

		flags := &runtime.Flags{}

		env := &runtime.Terminal{}
		env.Init(flags)
		defer env.Close()

		switch args[0] {
		case "set":
			if err := dsc.Set(env.Cache(), args[1]); err != nil {
				fmt.Printf("Error setting desired state configuration: %v", err)
			}
		case "schema":
			fmt.Print(dsc.Schema)
		case "export", "get":
			state := dsc.Get(env.Cache())
			fmt.Print(state.String())
		default:
			_ = cmd.Help()
		}
	},
}

func init() {
	RootCmd.AddCommand(dscCmd)
}
