// entry point Execute() for cobra - read: https://github.com/spf13/cobra/blob/main/user_guide.md
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	Verbose 	bool

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
	Use:   "arch-sample-job",
	Short: "Sample application for building job services (UI-less with a queue as input)",
	Long: `This sample job service provides a standard way to do message-queue based background jobs 
and should be the prescriptive/opinionated starting point for new and upgraded 
Go and Kubernetes job's.  Use arch-sample-job help for detailed assistance or check online at
https://docs.chef.io`,
})

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Here you will define your flags and configuration settings. See https://github.com/spf13/cobra/blob/main/user_guide.md
	// and https://github.com/spf13/cobra/blob/main/user_guide.md 

	cobra.OnInitialize(initConfig)

	// persistent flags apply to root and all child commands (arch-sample-cli --verbose && arch-sample-cli exec --verbose, for example)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.chef.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	// rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")

	// to add more, https://github.com/spf13/cobra-cli/blob/main/README.md 
	// all these are added by discovery of the /cmd folder
	// rootCmd.AddCommand(domainAPICmd)
	// rootCmd.AddCommand(execCmd)
	// rootCmd.AddCommand(acceptEULACmd)
	// supplied automatically rootCmd.AddCommand(helpCmd), can customize per user guide
	// rootCmd.AddCommand(updateCmd)
	// rootCmd.AddCommand(versionCmd)	// can also customize
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".chef" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".chef")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func init() {}