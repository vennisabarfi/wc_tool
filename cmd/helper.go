package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

// function to open and read files
func FileHelper() {

}

// save file in root with main.go
// count number of bytes in a function
func byteCounter(file_name string) (byte_length int) {
	buffer, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}
	byte_length = len(buffer)
	return byte_length
}
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

		// use byte counter function
		byte_count := byteCounter(file_name)
		fmt.Println(byte_count, file_name)

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
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(count, file_name)

	}

}

func wordCounter(file_name string) (count int) {
	// initialize a buffer to access file content
	buffer, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}
	// initialize a scanner to be used by ScanWords for word count
	scanner := bufio.NewScanner(bytes.NewReader(buffer))
	scanner.Split(bufio.ScanWords)
	// count words
	count = 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error while scanning file file:", err)
	}
	return count
}

// output the number of words in a file
func NumberofWords(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("words") {

		file_name, err := cmd.Flags().GetString("words")
		if err != nil {
			panic(err)
		}

		// returns an io.Reader for later
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

		word_count := wordCounter(file_name)

		fmt.Println(word_count, file_name)

	}
}

func characterCounter(file_name string) (char int) {
	// initialize a buffer to access file content
	buffer, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}

	char = utf8.RuneCount(buffer)

	fmt.Println("char", utf8.RuneCount(buffer))
	return char
}

// output the number of characters in a text
func NumberofCharacters(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("char") {

		file_name, err := cmd.Flags().GetString("char")
		if err != nil {
			panic(err)
		}

		// returns an io.Reader for later
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
		char := characterCounter(file_name)
		fmt.Println(char, file_name)
	}
}

// no flag option. output -c, -l and -w options (byte count, # of lines, # of words)
func NoFlagOption(cmd *cobra.Command, args []string) {

	//  convert array of strings to just string
	file_name := strings.Join(args, " ")

	file, err := os.Open(file_name)
	if err != nil {
		// panic(err)
		log.Fatal("File failed to open: ", err)
	}
	// check for error to be returned and exit if file not found
	defer func() {
		if err := file.Close(); err != nil {
			// panic("File closed" err)
			log.Fatal("File closed: ", err)
		}
	}()

	//  initialize file to be read by lineCounter function with reader
	reader := io.Reader(file)
	count, err := lineCounter(reader)
	if err != nil {
		log.Fatal("Error counting file lines", err)
	}

	byte_count := byteCounter(file_name)

	word_count := wordCounter(file_name)

	fmt.Println(count, word_count, byte_count, file_name)

	// // initialize a buffer to access file content
	// buffer, err := os.ReadFile(file_name)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("char", utf8.RuneCount(buffer))

}
