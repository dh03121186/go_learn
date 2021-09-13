package main

func main() {

	// map定义
	var map1 map[string]string // nil
	map1 = make(map[string]string)

	map1["France"] = "巴黎"
	map1["Italy"] = "罗马"
	map1["Japan"] = "东京"
	map1["India "] = "新德里"

	// 删除元素
	delete(map1, "France")

	// 键不存在时返回值 的缺省值（类型的0值）
	println("Franced", map1["Franced"])
}
