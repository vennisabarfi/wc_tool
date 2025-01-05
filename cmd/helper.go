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
		file_name, err := cmd.Flags().GetString("count")
		if err != nil {
			panic(err)
		}

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

func lineCounter(r io.Reader) (int, error) {
	// create a buffer
	buffer := make([]byte, 32*1024)
	count := 0
	line_separator := []byte{'\n'} //separate by new line

	for {
		c, err := r.Read(buffer)
		count += bytes.Count(buffer[:c], line_separator)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func NumberofLines(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("lines") {

		file_name, err := cmd.Flags().GetString("lines")
		if err != nil {
			panic(err)
		}

		// returns an io.Reader for later
		file, err := os.Open(file_name)
		if err != nil {
			panic(err)
		}

		//  initialize file to be read by lineCounter function
		r := io.Reader(file)
		count, err := lineCounter(r)

		fmt.Println(count, file_name)

	}

}

// output the number of words in a file
func NumberofWords(cmd *cobra.Command, args []string) {

}
