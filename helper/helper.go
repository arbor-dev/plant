/**
* Copyright © 2018, ACM@UIUC
*
* This file is part of the Arbor Project.
*
* The Arbor Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package helper

import (
	"os"
	"path/filepath"
	"fmt"
)

func createDir(dir string, path string) {
	if _, e := os.Stat(path); e == nil {
		fmt.Println("Directory already exists; can't create arbor project.")
		os.Exit(-1)
	}
	
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func WriteToFile(script string, dir string, file string, newDir bool) error {
	if newDir {
		path, _ := filepath.Abs(dir)
		createDir(dir, path)
	}
	
	f, _ := os.Create(dir + "/" +file)


	defer f.Close()
	
	_, err := f.WriteString(script)
	f.Sync()
	return err
}

