package main

func main() {
	// ... 标识不指定数组长度 根据初始化数据自动推断
	balance1 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	// 指定元素个数的时候 可以通过 index:value 格式初始化指定位置的元素
	balance2 := [5]float32{1: 2.0, 3: 7.0}

	println(balance1, balance2)
}
