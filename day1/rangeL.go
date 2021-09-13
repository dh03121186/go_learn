package main

import "fmt"

func main() {

	//range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

	// 数组和切片
	nums := []int{1, 2, 3}
	for i, value := range nums {
		println(i, value)
	}

	// map
	map1 := make(map[string]string)

	map1["France"] = "巴黎"
	map1["Italy"] = "罗马"
	map1["Japan"] = "东京"
	map1["India "] = "新德里"
	for key := range map1 {
		println(key, "->", map1[key])

	}

	// range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
