package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("E:\\workspace\\src\\park_record", WalkFunc)
}

func WalkFunc(path string, info os.FileInfo, err error) error {
	if err := Skip(path, info); err != nil {
		return err
	}
	fmt.Println(path, "/", info.Name())
	return nil
}

func Skip(path string, f os.FileInfo) error {
	if f.IsDir() && f.Name() == "vendor" {
		return filepath.SkipDir
	}
	// exclude all hidden folder
	if f.IsDir() && len(f.Name()) > 1 && f.Name()[0] == '.' {
		return filepath.SkipDir
	}
	return nil
}
