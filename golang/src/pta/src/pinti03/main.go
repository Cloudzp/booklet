package main

/*
设计函数分别求两个一元多项式的乘积与和。

输入格式:
输入分2行，每行分别先给出多项式非零项的个数，再以指数递降方式输入一个多项式非零项系数和指数（绝对值均为不超过1000的整数）。数字间以空格分隔。

输出格式:
输出分2行，分别以指数递降方式输出乘积多项式以及和多项式非零项的系数和指数。数字间以空格分隔，但结尾不能有多余空格。零多项式应输出0 0。

输入样例:
4 3 4 -5 2 6 1 -2 0
3 5 20  -7 4  3  1
输出样例:
15 24 -25 22 30 21 -10 20 -21 8 35 6 -33 5 14 4 -15 3 18 2 -6 1
5 20 -4 4 -5 2 9 1 -2 0
*/

import (
	"fmt"
)

type Element struct {
	// 指数
	col int
	// 系数
	index int
	Next  *Element
}

func main() {
	var n1 int
	var n2 int

	fmt.Scan(&n1)

	ele01 := Create(n1)

	fmt.Scan(&n2)

	ele02 := Create(n2)

	ele01.multi(ele02).Print()

	ele01.plus(ele02).Print()

}

func Create(n int) *Element {
	var col int
	var index int

	ele := &Element{}
	for i := 0; i < n; i++ {
		fmt.Scan(&index)
		fmt.Scan(&col)
		if i == 0 {
			ele.index = index
			ele.col = col
		} else {
			if index != 0 {
				ele.addOne(&Element{col: col, index: index})
			}
		}
	}
	return ele
}

// addOne 向多项式中添加一个元素
func (e *Element) addOne(ele *Element) {
	for e != nil {
		if e.Next == nil {
			e.Next = ele
			return
		}

		e = e.Next
	}
	return
}

// 两个多项式相加，生成一个新的多项式
func (e *Element) plus(e2 *Element) (rootE3 *Element) {
	if e2 == nil {
		return e
	}
	// 遍历表达式
	e3 := rootE3

	for e2 != nil && e != nil {
		if e.col < e2.col {
			if e3 == nil {
				e3 = &Element{}
				e3.col = e2.col
				e3.index = e2.index
				rootE3 = e3
			} else {
				e3.Next = e2
				e3 = e3.Next
			}

			if e2.Next != nil {
				e2 = e2.Next
			} else {
				e3.Next = e
				return
			}

		} else if e.col > e2.col {

			if e3 == nil {
				e3 = &Element{}
				e3.col = e.col
				e3.index = e.index
				rootE3 = e3
			} else {
				e3.Next = e
				e3 = e3.Next
			}

			if e.Next != nil {
				e = e.Next
			} else {
				e3.Next = e2
				return
			}

		} else {
			if e3 == nil {
				rootE3 = e3
			}

			if e.index+e2.index != 0 {
				newE := &Element{}
				newE.col = e.col
				newE.index = e.index + e2.index
				e3.Next = newE
				e3 = e3.Next
			}

			if e3 == nil {
				return &Element{}
			}

			if e.Next != nil {
				e = e.Next
			} else {
				e3.Next = e2
				return
			}

			if e2.Next != nil {
				e2 = e2.Next
			} else {
				e3.Next = e
				return
			}
		}
	}
	return
}

func (e *Element) multi(e2 *Element) *Element {
	var root *Element
	e2Root := e2
	for e != nil {
		var e3Root *Element
		e3 := e3Root
		for e2 != nil {
			if e3 == nil {
				e3 = &Element{
					col:   e2.col + e.col,
					index: e2.index * e.index,
				}
				e3Root = e3
			} else {
				newE := &Element{}
				newE.col = e2.col + e.col
				newE.index = e2.index * e.index
				e3.Next = newE
				e3 = e3.Next
			}
			e2 = e2.Next
		}
		root = e3Root.plus(root)

		e = e.Next
		e2 = e2Root
	}

	return root
}

func (e *Element) Print() {
	var flag = true
	for e != nil {
		if flag {
			flag = false
			fmt.Print(e.index, " ", e.col)
		} else {
			fmt.Print(" ", e.index, " ", e.col)
		}

		e = e.Next
	}
	fmt.Println()
}
