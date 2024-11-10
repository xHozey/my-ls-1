package ls

import (
	"fmt"
	"log"
	"os"
)

func Ls(dir string) {
	file, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for i, val := range file {
		fmt.Print(val.Name())
		if i != len(file)-1 {
			fmt.Print(" ")
		}
	}
}

func Ls_R(dir string) {
	file, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	folders := make([]string, 0, len(file))
	fmt.Print(dir + ":")
	if len(file) > 0 {
		fmt.Println()
	}
	for i, val := range file {
		if val.Name()[0] == '.' {
			continue
		}
		if val.IsDir() {
			fmt.Print(val.Name())
			folders = append(folders, dir+"/"+val.Name())
		} else {
			fmt.Print(val.Name())
		}
		if i < len(file)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	fmt.Println()
	for _, val := range folders {
		Ls_R(val)
	}
}

func Ls_l(dir string) {
	// file, err := os.ReadDir(dir)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("total %d", len(file))
}

func Ls_a(dir string) {
}

func Ls_r(dir string) {
}

func Ls_t(dir string) {
}
