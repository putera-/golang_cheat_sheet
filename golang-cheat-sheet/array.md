# Array

- data length must be declare

```go
var names [2]string
names[0] = "Uda"
names[1] = "Putera"
```

### set value on declaration

```go
var values = [3]int{
    1,
    2,
    3,
}

fmt.Println(values) // [1 2 3]

var values2 = [3]int{1, 2, 3}

fmt.Println(values2) // [1 2 3]
```

### default value based on data type

```go
var values = [3]int{
    1,
    2,
}

fmt.Println(values) // [1 2 0]
```

### dynamic array length

```go
var values = [...]int{
    1,
    2,
}

fmt.Println(values) // [1 2]
fmt.Println(len(values)) // 2
```

## Array function

```go
len(array)          // get length of array
array[index]        // get value on selected index
array[index] = 1    // set value on selected index
```
