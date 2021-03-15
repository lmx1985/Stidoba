package main

import (
	"fmt"
	"io/ioutil"
)

func Dir(command string, path string) string {
	var ss string
	//var t string
	var dir string = "c:/Users"

	if command == string("cd") && len(path) != 0 {
		dir = dir + path + "/"
		return dir

	}
	if command == string("cd") && len(path) == 0 {
		dir = path + "/"
		return ""

	}
	if command == string("cd..") {
		dir = dir[0:len(path)] + "/"
		return ""
	}

	if command == string("dir") {

		fmt.Println(dir)

		filesFromDir, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Println(err)
		}

		for _, file := range filesFromDir {
			//if file.IsDir() {
			//	t = "Directory: "
			//} else {
			//	t = "File: "
			//}

			//fmt.Print(t)
			//fmt.Printf("%s, size: %d\n", file.Name(), file.Size())

			ss = ss + fmt.Sprintf("%s, size: %d\n", file.Name(), file.Size())

		}
		fmt.Print(ss)
		return ss
	}
	return ""
}
