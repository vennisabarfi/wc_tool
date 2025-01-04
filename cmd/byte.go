/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

func NumberofBytes(cmd *cobra.Command, args []string) {
	file_name, _ := cmd.Flags().GetString("file")
	print("Hello yayy", file_name)
}

// byteCmd represents the byte command
var byteCmd = &cobra.Command{
	Use:   "byte",
	Short: "This outputs the number of bytes in a file ",
	Long:  `This command receives a file and then outputs the number of bytes in said file.`,
	Run:   NumberofBytes,
}

func init() {
	rootCmd.AddCommand(byteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// byteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	rootCmd.Flags().StringP("file", "c", "", "Output number of bytes in a file") //attached to root so I can use as a global flag instead
	// byteCmd.Flags().StringP("file", "c", "", "Output number of bytes in a file")

}
