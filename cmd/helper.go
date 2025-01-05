package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// function to open and read files
func FileHelper() {

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

func NumberofLines(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("lines") {

		file_name, _ := cmd.Flags().GetString("lines")
		// returns an io.Reader for later
		file, err := os.Open(file_name)
		if err != nil {
			panic(err)
		}

		// create a buffer with a counter and separator by new line
		buffer := make([]byte, 32*1024)
		count := 0
		line_separator := []byte{'\n'}

		r := io.Reader(file)

		for {

			c, err := r.Read(buffer)
			count += bytes.Count(buffer[:c], line_separator)
			// fmt.Print("Oh yeahL", count)

			switch {
			case err == io.EOF:
				fmt.Print("Heres':", count, nil)

			case err != nil:
				fmt.Print("Yapa", count, err)

			}

		}
	}

}
