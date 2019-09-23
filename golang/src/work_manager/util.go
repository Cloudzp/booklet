package work_manager

import (
	"log"
	"runtime/debug"
)

// HandleCrash 捕获panic 打印堆栈信息到日志文件及stdout中
func HandleCrash() {
	if r := recover(); r != nil {
		log.Fatalf("panic error: %v \n %s", r, string(debug.Stack()))
	}
}
