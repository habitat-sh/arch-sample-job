// 'arch-sample-cli acceptEULA' should be done before any operations work
package cmd

import (
   "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(acceptEULACmd)
}

var acceptEULACmd = &cobra.Command{
  Use:   "acceptEULA",
  Short: "Accepts the Progress end-user license agreement for this and connected products",
  Long:  `Nothing works until you say yes...`,
  Run: func(cmd *cobra.Command, args []string) {
   //  fmt.Println("An upgrade to arch-sample-cli v1.0.2 is available")
  },
}