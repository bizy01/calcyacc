package main

import (
	"github.com/bizy01/calcyacc/parse"
	"fmt"
	"bufio"
	"os"
	"strings"
)

var (
	input string
)

func main() {
	f := bufio.NewReader(os.Stdin)

    for {
        fmt.Print(">")
        input, _ = f.ReadString('\n')
        if len(input) == 1 {
            continue
        }

        input = strings.TrimSpace(input)

        if input == "exit" || input == "quit" || input == "bye" {
        	os.Exit(1)
        }

        res, err := parse.Parse(input)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
    }
}


