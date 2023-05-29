package root

import (
	"hello-world/pkg/hello"

	"github.com/spf13/cobra"
)

var FlagName string

var Cmd = &cobra.Command{
	Use: "hello-world",
	Run: func(c *cobra.Command, args []string) {
		hello.PrintHello(FlagName)
	},
}

func init() {
	Cmd.Flags().StringVarP(
		&FlagName,
		"name",
		"n",
		"World",
		"Name to say hello to",
	)
}
