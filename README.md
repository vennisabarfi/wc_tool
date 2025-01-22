# wc_tool
Built a clone of the wc_tool in Unix
# Unix WC_TOOL Clone in Go!

Built a clone of the Unix wc_tool in Go. This is a first in a series of projects in Go following [codingchallenges.fyi](https://codingchallenges.fyi/challenges/challenge-wc/)

## Description

Using this cli tool, you can view the:
- Number of bytes in a text file.
- Number of lines in a text file.
- Number of words in a text file.
- Number of characters in a text file.

This command line tool is built using Cobra CLI in Go.

## Getting Started

1. Clone the repo and use the help flag to view commands

   ```sh
    git clone https://github.com/vennisabarfi/wc_tool
    go run main.go -h 
    ```
    or 
    ```sh
    git clone https://github.com/vennisabarfi/wc_tool
    go run main.go --help 
    ``` 

2. Current Commands
    ```sh
   This is a clone of the wc tool in Unix but in Go! 
   Run go run main.go -flag file_name
   wctool [flags]
   ccwc        Default command to instantly find count of bytes, number of words and number of lines in a file.

   Flags:
   -c, --count string   Count number of bytes in a file
   -h, --help           help for wctool
   -l, --lines string   Find number of lines in a file
   -w, --words string   Find number of words in a file
    ```

3. To run the default command
    ```sh
    go run main.go ccwc "test.txt"
    ``` 
    - This should output (corresponding to -c, -l and -w options )
    ```sh
    7145 58164 342190 test.txt
    ``` 

4. Example commands. Let's say the number of bytes in a file (-c option)
    ```sh
    go run main.go -c "test.txt"
    ``` 
     - This should output (corresponding to -c, -l and -w options )
    ```sh
    342190 test.txt
    ``` 
### Dependencies

Currently runs on go 1.23.4. Please view ([go.mod](https://github.com/vennisabarfi/wc_tool/blob/main/go.mod)) for more details on dependencies.


### Installing

- Save text file in root of directory (same level as main.go)


## Authors


[Vennisa Barfi](https://github.com/vennisabarfi)

## Version History
* 0.1
    * Initial Release

## Acknowledgments


* Problem found at: [codingchallenges.fyi](https://codingchallenges.fyi/challenges/challenge-wc/)