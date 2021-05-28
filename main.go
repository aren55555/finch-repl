package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	delimitter = " "
)

func main() {
	fmt.Println("REPL...")

	s := NewStore()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input := scanner.Text()

			args := strings.Split(input, " ")

			if len(args) == 0 || args[0] == "" {
				fmt.Println("input error")
				continue
			}

			switch cmd := args[0]; cmd {
			case "BEGIN":
				s.Begin()
				fmt.Println("Done!")
			case "ROLLBACK":
				s.Rollback()
				fmt.Println("Done!")
			case "COMMIT":
				s.Commit()
				fmt.Println("Done!")
			case "GET":
				fmt.Printf("%s => %d\n", args[1], s.Get(args[1]))
			case "SET":
				v, err := strconv.Atoi(args[2])
				if err != nil {
					fmt.Printf("non integer input detected %q", args[2])
					continue
				}
				if err := s.Set(args[1], v); err != nil {
					fmt.Printf("error: %v\n", err)
				}
			case "DEL":
				if err := s.Del(args[1]); err != nil {
					fmt.Printf("error: %v\n", err)
				}
			default:
				fmt.Printf("invalid argument %q\n", cmd)
			}
		}
	}
}
