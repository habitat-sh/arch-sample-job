// arch-sample-cli exec 128.0.0.1 --powershell <file>
// exec (passthrough to cmdline from file or single cmd line in quotes, to powershell, cmd.exe, bash, sh, or Habitat shell)
// --ssh "commands to pass through" to a node
package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(execCmd)
}

var execCmd = &cobra.Command{
  Use:   "exec",
  Short: "CLI to shell executor",
  Long:  `When the CLI needs to call powershell, bash, or whatever... takes a file and funnels to os.exec`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("PS>")
  },
}