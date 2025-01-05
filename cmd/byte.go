/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// save file in root with main.go

func NumberofBytes(cmd *cobra.Command, args []string) {
	file_name, _ := cmd.Flags().GetString("file")
	fmt.Println("File is: ", file_name)

	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	// check for error to be returned and exit if file not found
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	buffer, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}
	byte_length := len(buffer)

	fmt.Println(byte_length, file_name)
	fmt.Printf("Byte length is %d for %s", byte_length, file_name)

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
	byteCmd.PersistentFlags().StringP("file", "c", "", "Output number of bytes in a file")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rootCmd.Flags().StringP("file", "c", "", "Output number of bytes in a file")
	//attached to root so I can use as a global flag instead
	byteCmd.Flags().StringP("file", "c", "", "Output number of bytes in a file")

}
