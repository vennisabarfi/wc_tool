/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wctool",
	Short: "Go Version of the Iconic wc in Unix",
	Long:  `This is a clone of the wc tool in Unix but in Go! Run go run main.go -flag file_name`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// Call multiple functions here

		NumberofBytes(cmd, args)
		NumberofLines(cmd, args)
		NumberofWords(cmd, args)
		NumberofCharacters(cmd, args)
	},
}

// default command with no flags
var defaultCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "Default command to instantly find count of bytes, number of words and number of lines in a file!",
	Long:  `run ccwc file_name`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// Call multiple functions here

		NoFlagOption(cmd, args)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(defaultCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wctool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("count", "c", "", "Count number of bytes in a file")
	rootCmd.Flags().StringP("words", "w", "", "Find number of words in a file")
	rootCmd.Flags().StringP("lines", "l", "", "Find number of lines in a file")
	rootCmd.Flags().StringP("char", "m", "", "Find number of characters in a file")
	// rootCmd.Flags().StringP("main", "", "", "Main. Outputs -c, -l  and -w options")
	// something wrong with not using flags
}
