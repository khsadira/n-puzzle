package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func showPrompt() {
	print("> ")
}

func getUserEntry() (string, []string) {
	var args []string

	showPrompt()
	in := bufio.NewReader(os.Stdin)
	userEntry, err := in.ReadString('\n')

	if err != nil {
		log.Fatal("n-puzzle: error:", err.Error())
	}

	userEntryArr := strings.Split(strings.TrimRight(userEntry, "\r\n"), " ")
	cmd := userEntryArr[0]

	if len(userEntryArr) > 1 {
		args = userEntryArr[1:]
	}

	return cmd, args
}
