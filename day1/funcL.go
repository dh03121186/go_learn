package main

// 值传递，参数是复制出来的，这里的操作不会影响原数据
func max(num1, num2 int) int {
	if num1 >= num2 {
		return num1
	} else {
		return num2
	}
}

func min(num1, num2 int) int {
	if num1 >= num2 {
		return num2
	} else {
		return num1
	}
}

// 引用传递，传递的是内存地址，操作会影响原数据
func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

// 函数也可以作为其他函数的参数，todo 但是出参入参如何兼容
func dofunc(num1 int, num2 int, funcL func(num1, num2 int) int) int {
	return funcL(num1, num2)
}

// 闭包 getSequence 返回一个匿名函数 匿名函数的优越性在于可以直接使用函数内的变量
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集
//func (variable_name variable_data_type) function_name() [return_type]{
//	/* 函数体*/
//}

type Circle struct {
	radius float64
}

// 该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func main() {
	println(max(1, 2))
	println(dofunc(1, 2, min))
	println(dofunc(1, 2, max))

	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	println(nextNumber()) // 1
	println(nextNumber()) // 2
	println(nextNumber()) // 3

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	println(nextNumber1()) // 1
	println(nextNumber1()) // 2
}
