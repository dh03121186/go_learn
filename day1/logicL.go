package main

import "fmt"

func main() {
	value1 := -2
	value1 += 1

	// go 不支持 py 类似的语法 if value1:
	if value1 > 0 {
		fmt.Println("value1 大于0")
	} else {
		fmt.Println("value1 小于0")
	}

	sum := 0
	// for init; condition; post { }
	for i := 0; i <= 10; i++ {
		sum += i
		fmt.Println(i)
	}
	fmt.Println(sum)

	// for condition { }
	i := 0
	for i <= 10 {
		sum += i
		fmt.Println(i)
		i++
	}
	fmt.Println(sum)

	// 死循环
	//for {
	//	fmt.Println(i)
	//}

	//for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环
	str1 := "abcd"
	for index, value := range str1 {
		fmt.Println(index, value)
	}
}
