# Slice

3 data types of slice

- Pointer
- Length
- Capacity

### Create slice from array

```go
array[low:high]     // create slice of array from index low to high index
array[low:]         // create slice of array from index low to last index of array
array[:high]        // create slice of array from index 0 to high index
array[:]            // create slice of array from index 0 to last index of array
```

```go
array[:]            // convert from array to slice
```

### Example

```go
var month = [12]string {
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
slice2 := month[6:9] // pointer -> "jul", length -> 3, capacity -> 6
slice3 := month[6:]  // pointer -> "jul", length -> 6, capacity -> 6
slice4 := month[:4]  // pointer -> "jan", length -> 4, capacity -> 4
slice5 := month[:]   // pointer -> "jan", length -> 12, capacity -> 12
fmt.Println(slice1)  // [mei jun jul]
fmt.Println(slice2)  // [jul aug sep]
fmt.Println(slice3)  // [jul aug sep okt nov des]
fmt.Println(slice4)  // [jan feb mar apr]
fmt.Println(slice5)  // [jan feb mar apr mei jun jul aug sep okt nov des]

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
```

### create slice

```go
newSlice := make([]string, 2, 5) // length -> 2, capacity -> 5
newSlice[0] = "Uda"
newSlice[1] = "Putera"

fmt.Println(newSlice) // [Uda Putera]
fmt.Println(len(newSlice)) // 2
fmt.Println(cap(newSlice)) // 5


newSlice[2] = "Padang" // will be error, must use append
newSlice = append(newSlice, "Padang")

fmt.Println(newSlice)		// [Uda Putera Padang]
fmt.Println(len(newSlice))	// 3
fmt.Println(cap(newSlice))	// 5
```

### copy slice

```go
fullname := make([]string, 2, 5) // length -> 2, capacity -> 5
fullname[0] = "Uda"
fullname[1] = "Putera"

copyFullname := make([]string, len(fullname), cap(fullname))
copy(copyFullname, fullname)

fmt.Println(fullname)
fmt.Println(copyFullname)
```

## ARRAY VS SLICE declaration

```go
iniArray := [...]int{1, 2, 3, 4}
iniSlice := []int{1, 2, 3, 4}

fmt.Println(iniArray)
fmt.Println(iniSlice)
```
