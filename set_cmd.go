package main

import (
	"strconv"
	"strings"
)

func setCmd(args []string) {
	if len(args) == 0 {
		println("set: usage: set heur=X algo=X")
	}
	for _, arg := range args {
		setArg := strings.Split(arg, "=")

		if len(setArg) != 2 {
			println("set:", arg, "argument not well formatted. Should be \"set heur=X algo=X\"")
			continue
		}

		value, err := strconv.Atoi(setArg[1])

		if err != nil {
			println("set:", setArg[1], "value not well formatted.")
			continue
		}

		if value >= 0 && value <= 2 {
			if setArg[0] == "heur" {
				heur = value
				println("set: heuristic set to:", value)
			} else if setArg[0] == "algo" {
				algo = value
				println("set: algorithm set to:", value)
			}
		}
	}
}
