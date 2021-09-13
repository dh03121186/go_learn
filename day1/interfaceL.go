package main

import (
	"fmt"
	"reflect"
)

// 转型 如果 A 包含 B， 那么A可以转型成B，B不能转成A；

// Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。
type Nothing interface {
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone *NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone *IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

// 函数接收空接口作为参数，因此，可以给这个函数传递任何类型。
func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	ss := "ssss"
	var ii interface{} = 55
	describe(ss)
	nothing := new(Nothing)
	describe(nothing)
	//nothing = phone
	println(reflect.TypeOf(ss).Name())
	// interface.(int) 类型断言 interface 必须为接口 而不是基础类型
	value, ok := ii.(int) // 提取ii底层的int值
	println(value, ok)
	// s := ss.(string) // Invalid type assertion: ss.(string) (non-interface type string on left)
}
