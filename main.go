package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func dirTree(path string, add string) {
	dir, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for idx, file := range dir {
		var isLast bool

		if idx == len(dir) - 1 {
			isLast = true
		} else {
			isLast = false
		}

		fmt.Print(add)

		if isLast {
			fmt.Print("└")
		} else {
			fmt.Print("├")
		}

		fmt.Println("───" + file.Name())

		if file.IsDir() {
			newPath := path + file.Name() + "/"

			var newAdd string

			if isLast {
				newAdd = add + "  "
			} else {
				newAdd = add + "| "
			}

			dirTree(newPath, newAdd)
		}
	}
}

func main() {
	startPath := "" // you should write here full path to dir which you want to start

	_, err := ioutil.ReadDir(startPath)

	if err != nil {
		log.Fatal(err)
	}

	dirTree(startPath, "")
}
