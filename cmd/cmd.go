package cmd

import (
	"hello-world/cmd/root"
	_ "hello-world/cmd/version"

	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
