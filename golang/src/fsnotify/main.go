package main

import (
	"fmt"
	"gopkg.in/fsnotify.v1"
	"path/filepath"
)

func main() {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(filepath.Dir("E:/workspace/src/park_record/config/service.json"))
	if err := watcher.Add("E:/workspace/src/park_record/config"); err != nil {
		fmt.Println(err)
		return
	}

	for {
		select {
		case op := <-watcher.Events:
			switch op.Op {
			case fsnotify.Create:
				fmt.Println("创建了一个文件")

			case fsnotify.Remove:
				fmt.Println("删除了一个文件")
			case fsnotify.Write:
				fmt.Println("文件新添加了内容", op.Name)
			}
		}

	}

}
