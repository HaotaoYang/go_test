package main

import (
	"fmt"
	"log"
	"pack1" // 导入自定义包
	"runtime"
	"time"
)

func main() {
	start_time := time.Now()

	// 打印当前的go版本
	fmt.Printf("%s\n", runtime.Version())

	// 调试测试，打印调用文件和行数
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()

	// 数组和数组指针做参数的区别
	array := [3]float64{7.0, 8.5, 9.1}
	x := sum1(&array) // 这里的参数是数组array的指针，这样做的目的可以节省空间，只传递一个指针，同时也可以通过sum函数改变数组的值
	y := sum2(array)  // 这里的参数是数组array
	fmt.Println("The sum1 of the array is : ", x)
	fmt.Println("The sum2 of the array is : ", y)

	// 测试import自定义包
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pack1.Pack1Int)
	fmt.Printf("Float from package1: %f\n", pack1.Pack1Float)

	// 输出程序执行时间
	end_time := time.Now()
	delta := end_time.Sub(start_time)
	fmt.Printf("start_time: %s\n", start_time)
	fmt.Printf("end_time: %s\n", end_time)
	fmt.Printf("delta: %s\n", delta)
}

func sum1(a *[3]float64) (sum float64) {
	for _, x := range a {
		sum += x
	}
	return
}

func sum2(a [3]float64) (sum float64) {
	for _, y := range a {
		sum += y
	}
	return
}
