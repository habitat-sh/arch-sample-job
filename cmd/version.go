// 'arch-sample-cli version' lists this executable's VERSION, and domain API versions supported
package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of arch-sample-cli",
  Long:  `Lists this executable's VERSION, and domain API versions supported`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("arch-sample-cli v1.0.1")
    fmt.Println("  supports local-licensing-service (1.0.0, )")
  },
}