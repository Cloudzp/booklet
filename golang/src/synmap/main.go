package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var syncMap sync.Map

	// 存储一个数据
	syncMap.Store("key", "value")
	// 根据key获取value，如果不存在返回nil，并且ok为false
	v, ok := syncMap.Load("key")
	if ok {
		fmt.Println(v)
	}

	// 删除一个值
	syncMap.Delete(v)
	v, ok = syncMap.Load("key")
	fmt.Println("删除后", v, ok)

	// 查找或保存值
	// 如果能够根据key查找到值，就返回改值及true
	// 如果不能根据key查询到值，就保存key及value并返回false；
	v, ok = syncMap.LoadOrStore("key2", "1")
	fmt.Println("LoadOrStore", v, ok)
	v, ok = syncMap.LoadOrStore("key2", "2")
	fmt.Println("LoadOrStore", v, ok)
	time.Sleep(2 * time.Second)
}
