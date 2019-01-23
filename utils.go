package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	msg := scanner.Text()
	return msg
}

func check_for_cmd(msg string) {
	active := 1
	for active == 1 {
		if strings.HasPrefix(msg, ">> cc") {
			fmt.Print("Enter new id: ")
			inp = read()
		} else if strings.HasPrefix(msg, ">> h") {
			fmt.Println("----------------\nIds:\n/dev/null:general: 422293824770146306\n PH:general:        469851459966730262\n FC:general:        439871916082331650\n----------------")
		} else if strings.HasPrefix(msg, ">> ls") {
			for i := range channels {
				fmt.Printf("%s\t%s\n", channels[i].Name, channels[i].ID)
			}
		} else {
			fmt.Printf("No command %s found\n", msg)
		}
		active = 0
	}
}
