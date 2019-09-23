package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

func main() {

	//TestMap()
	//TestSwitch()
	//Test_goroutein_map()
	TestInt()

}

/*
 问题1： 将一个空字符串付给interface后，是否可以将interface转换成string类型？
 验证结果：可以，空字符串是有指针的。

*/
func TestString2Interface() {
	/* var str = ""
	   var test = make(map[string]interface{})

	   //test["guid"] = ""
	   fmt.Println(&str)

	   u := test["guid"].(string)
	   fmt.Println(u)*/
}

/*
 问题2：将map传入方法中删除一个元素，在方法外遍历这个map，它的元素是否也被删除了？
 验证结果：是的，map默认是指针传递。

*/
func TestMap() {
	var data = map[string]interface{}{}
	data["test"] = "test"
	fmt.Println(data)
	fmt.Println(&data)
	DeleteOne(data)
	fmt.Println(data)
}

func DeleteOne(data map[string]interface{}) {
	fmt.Println(fmt.Sprintf(">- %#v", data))
	delete(data, "test")
}

/*
问题3：测试switch case 多项匹配 多项使用 ‘，’ 而不是 ‘|’
*/

func TestSwitch() {
	d := 5
	switch d {
	case 1 | 5 | 3 | 4 | 5:
		fmt.Println(d)
	case 6:
		fmt.Println(d)
	}
}

/*
问题：为什么只有数字会被改变
*/
func Test_goroutein_map() {
	var chanMap sync.Map
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("i=%d", i)

		workerChan := make(chan map[string]interface{}, 1000)
		chanMap.Store(key, workerChan)

		go func(chan map[string]interface{}) {
			fmt.Println(fmt.Sprintf("%#v", workerChan))

			fmt.Println(key, <-workerChan)
		}(workerChan)

	}

	chanMap.Range(func(key, value interface{}) bool {
		var testmap = map[string]interface{}{key.(string): "test"}

		chans := value.(chan map[string]interface{})
		chans <- testmap
		return true
	})
	time.Sleep(time.Second)
}

const (
	i1 = iota
	i2
	i3
)

/*
问题： 一个int类型传入interface中 经过marshal unmarshal 变成float
结论：unmarshal时由于获取不到字段的实际类型，为了保证精确度不会丢失，所以采用精确度最高的字段接收；
*/
func TestInt() {
	var i int

	var x = map[string]interface{}{}

	i = i1

	x["t"] = i

	b, _ := json.Marshal(x)

	var x2 map[string]interface{}
	json.Unmarshal(b, &x2)

	var tt = x2["t"]
	fmt.Println(tt.(int))

}
