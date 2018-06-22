package helper

import (
	//"io"
	"os"
	//"path/filepath"
	"fmt"
	//"strings"
)

func createDir(dir string) {
	err := os.MkdirAll(dir, 0777)
	if (err != nil) {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func WriteToFile(script string, dir string, file string) {
	createDir(dir)

	f, _ := os.Create(dir + "/" +file)


	defer f.Close()
	
	f.WriteString(script)
	f.Sync()
}
