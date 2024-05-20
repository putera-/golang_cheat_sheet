package main

import "fmt"

func main() {
	fmt.Println(string("test"[1]))
	fmt.Println("Hello World Putera")

	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32) // 32768
	fmt.Println(nilai64) // 32768
	fmt.Println(nilai16) // -32768

	var names [2]string
	names[0] = "Uda"
	names[1] = "Putera"
	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [3]int{1, 2, 3}

	fmt.Println(values)
}
