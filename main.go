package main

import (
	"fmt"
	"os"
	"slices"

	ls "ls/fucns"
)

func main() {
	arg := os.Args[1:]

	flags := []string{"-l", "-R", "-a", "-r", "-t"}
	switch len(arg) {
	case 0:
		ls.Ls(".")
		fmt.Println()
		break
	case 1:
		if slices.Contains(flags, arg[0]) {
			switch arg[0] {
			case "-l":
				ls.Ls_l(".")
			case "-R":
				ls.Ls_R(".")
				fmt.Println()
			case "-a":
				ls.Ls_a(".")
			case "-r":
				ls.Ls_r(".")
			case "-t":
				ls.Ls_t(".")
			}
		} else {
			ls.Ls(arg[0])
			fmt.Println()
		}
		break
	case 2:
		if !slices.Contains(flags, arg[0]) {
			fmt.Println("Usage: [OPTION] [FILE]")
		} else {
			switch arg[0] {
			case "-l":
				ls.Ls_l(arg[1])
			case "-R":
				ls.Ls_R(arg[1])
			case "-a":
				ls.Ls_a(arg[1])
			case "-r":
				ls.Ls_r(arg[1])
			case "-t":
				ls.Ls_t(arg[1])
			}
		}
	}
}
