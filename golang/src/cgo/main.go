/*
参考： https://studygolang.com/articles/2629
*/

//TODO 回调函数还没有练习，目前还不知道有什么作用
package main

/*
struct POINT_ALPHA
{
   int x;
   int y;
};
// 简化版
typedef struct _POINT_BETA
{
   int x;
   int y;
}POINT_BETA;

// 联合体
typedef long LONG;
typedef unsigned long DWORD;
typedef long long LONGLONG;

typedef union _LARGE_INTEGER {
    struct {
        DWORD LowPart;
        LONG HighPart;
    };
    struct {
        DWORD LowPart;
        LONG HighPart;
    } u;
    LONGLONG QuadPart;
} LARGE_INTEGER, *PLARGE_INTEGER;

void AAA(LARGE_INTEGER li)
{
    li.u.LowPart = 1;
    li.u.HighPart = 4;
}


int b = 6;
int *p = &b;

 int PlusOne(int n)
{
   return n + 1;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
	"golang.org/x/text/encoding"
)

func main() {
	//fmt.Println(C.PlusOne(1))
	//TestPointer()
	//TestStrust()
	encoding.ASCIISub
	TestUnion()
}

// c指针转换为go指针
// 设C代码中有int*类型的指针p，利用*C.p就可以获取到它所指向的变量的值（和C语言中指针的用法相同），利用*C.p则可以修改它所指向的变量的值。
func TestPointer(){
	var c *int32
    fmt.Println("b1 = ", C.b)
    *C.p = 99
	fmt.Println("b2 = ", C.b)
    c = (*int32)(unsafe.Pointer(C.p))
	*c = 100
	fmt.Println("b2 = ", C.b)
}

// 结构体

func TestStrust(){
   var p C.struct_POINT_ALPHA
   p.x = 6
   p.y = 7
   fmt.Println(p)
   // 简化版
   var p2 C.POINT_BETA
	p2.x = 6
	p2.y = 7
   fmt.Println(p2)
}

// 4.6 联合体
//  Go中使用C的联合体是比较少见的，而且稍显麻烦，因为Go将C的联合体视为字节数组。比方说，下面的联合体LARGE_INTEGER被视为[8]byte。
func TestUnion(){
	var li C.LARGE_INTEGER
	var b [8]byte = li //正确，因为[8]byte和C.LARGE_INTEGER相同
	b[0]=75
	C.AAA(b) //参数类型为LARGE_INTEGER，可以接收[8]byte
	fmt.Println("li = ", li)
	li[1] = 99
	fmt.Println("li = ", li)
}



