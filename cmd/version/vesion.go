package root

import (
	"fmt"
	"hello-world/cmd/root"
	"hello-world/version"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
