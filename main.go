package main

import (
	"github.com/adhara-tech/kubesu/cmds"
	"github.com/spf13/cobra"
)

func main() {
	cmds.InitCmdConnectNodes()

	var rootCmd = &cobra.Command{}
	rootCmd.AddCommand(cmds.CmdConnectNodes)
	rootCmd.Execute()
}
