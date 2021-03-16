package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var t string
var dir string = "c:/Users/"

func Dir(command string, path string) string {
	var ss string

	if command == string("exit") {
		return "exit"
	}

	if command == string("cd") && len(path) != 0 {
		dir = filepath.Join(dir, path)
		return dir

	}
	if command == string("cd") && len(path) == 0 {
		dir = filepath.Join(path)
		return ""

	}
	if command == string("cd..") {
		dir = filepath.Join(dir[0:len(path)])
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

			if file.IsDir() {
				t = "Directory: "
			} else {
				t = "File: "
			}

			ss = ss + fmt.Sprintf("%s %s, size: %d\n", t, file.Name(), file.Size())

		}
		fmt.Print(ss)
		return ss
	}
	return ""
}
