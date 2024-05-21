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

	// slice

	var month = [12]string{
		"jan",
		"feb",
		"mar",
		"apr",
		"mei",
		"jun",
		"jul",
		"aug",
		"sep",
		"okt",
		"nov",
		"des",
	}
	var slice1 []string = month[4:7] // pointer -> "mei", length -> 3, capacity -> 8
	slice2 := month[6:9]             // pointer -> "jul", length -> 3, capacity -> 6
	slice3 := month[6:]              // pointer -> "jul", length -> 6, capacity -> 6
	slice4 := month[:4]              // pointer -> "jan", length -> 4, capacity -> 12
	slice5 := month[:]               // pointer -> "jan", length -> 12, capacity -> 12
	fmt.Println(slice1)              // [mei jun jul]
	fmt.Println(slice2)              // [jul aug sep]
	fmt.Println(slice3)              // [jul aug sep okt nov des]
	fmt.Println(slice4)              // [jan feb mar apr]
	fmt.Println(slice5)              // [jan feb mar apr mei jun jul aug sep okt nov des]

	fmt.Println(len(slice1)) // 3
	fmt.Println(len(slice2)) // 3
	fmt.Println(len(slice3)) // 6
	fmt.Println(len(slice4)) // 4
	fmt.Println(len(slice5)) // 12

	fmt.Println(cap(slice1)) // 8
	fmt.Println(cap(slice2)) // 6
	fmt.Println(cap(slice3)) // 6
	fmt.Println(cap(slice4)) // 12
	fmt.Println(cap(slice5)) // 12

	slice4[1] = "notFeb"
	fmt.Println(slice4) // / [jan notFeb mar apr]
	fmt.Println(month)  // [jan notFeb mar apr mei jun jul aug sep okt nov des]

	// this will create new array
	slice6 := append(slice5, "no 13")
	slice6[0] = "notJan"
	fmt.Println(slice6) // [notJan notFeb mar apr mei jun jul aug sep okt nov des no 13]
	fmt.Println(month)  // [jan notFeb mar apr mei jun jul aug sep okt nov des]

	fmt.Println("new Slice")
	newSlice := make([]string, 2, 5) // length -> 2, capacity -> 5
	newSlice[0] = "Uda"
	newSlice[1] = "Putera"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// newSlice[2] = "Padang" // will be error, must use append
	newSlice = (append(newSlice, "Padang"))

	fmt.Println(newSlice)      // [Uda Putera Padang]
	fmt.Println(len(newSlice)) // 3
	fmt.Println(cap(newSlice)) // 5

	fullname := make([]string, 2, 5) // length -> 2, capacity -> 5
	fullname[0] = "Uda"
	fullname[1] = "Putera"

	copyFullname := make([]string, len(fullname), cap(fullname))
	copy(copyFullname, fullname)

	fmt.Println(fullname)
	fmt.Println(copyFullname)

	fmt.Println("array vs slice declaration")
	iniArray := [...]int{1, 2, 3, 4}
	iniSlice := []int{1, 2, 3, 4}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
