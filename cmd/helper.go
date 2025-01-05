package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func Printers(cmd *cobra.Command, args []string) {
	// check if flag has been set by user
	if cmd.Flags().Changed("message") {
		message, _ := cmd.Flags().GetString("message")

		fmt.Print("Hi lol lol ", message)
	}

}

// save file in root with main.go

func NumberofBytes(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("count") {
		file_name, _ := cmd.Flags().GetString("count")

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

}
