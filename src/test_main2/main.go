package main

import (
	"fmt"
	"struct_pack"
)

type struct1 struct {
	i   int
	f   float32
	str string
}

// 这个方法改变了recv的值，所以要写成struct1指针的类型
func (recv *struct1) change() {
	recv.i = 100
}

func (recv struct1) write() string {
	return fmt.Sprint(recv)
}

type IntA []int

func (a IntA) Swap() {
	// temp := array[i]
	// array[i] = array[j]
	// array[j] = temp
	a[0], a[1] = a[1], a[0]
}

func main() {
	// 测试结构体
	s := new(struct1)
	s.i = 10
	s.f = 1.1
	s.str = "hello"
	fmt.Println(s)  // s是一个指针，这里打印出指向的地址
	fmt.Println(*s) // 这里打印出s指向的地址的值

	// 测试自定义包中的结构体以及自定义包中方法的使用
	person := struct_pack.NewPerson("Holtom", 27)
	fmt.Println(person)
	fmt.Println(*person)
	person.SetPersonName("Holtom1")
	fmt.Printf("My name is %s.\n", person.GetPersonName())
	person.SetPersonAge(18)
	fmt.Printf("I am %d years old.\n", person.GetPersonAge())

	// 测试值或指针作为方法的接收者，a是值b是指针上面的方法都支持了
	var a struct1
	a.change()
	fmt.Println(a.write())

	b := new(struct1)
	b.change()
	fmt.Println(b.write())

	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	t := IntA(data)
	t.Swap()
	fmt.Printf("The sorted array is: %v\n", t)
}
