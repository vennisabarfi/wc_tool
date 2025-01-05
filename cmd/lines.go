/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Printer(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("name") {
		name, _ := cmd.Flags().GetString("name")
		fmt.Print("Hi", name)
	}
}

// linesCmd represents the lines command
var linesCmd = &cobra.Command{
	Use:   "lines",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lines called")
	},
}

func init() {
	rootCmd.AddCommand(linesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// linesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// linesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
