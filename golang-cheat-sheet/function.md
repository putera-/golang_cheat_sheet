# Function

```go
func sayHello() {
    fmt.Println("Hello")
}

func main() {
    sayHello()
}
```

### parameter

```go
func sayHelloTo(firstname string, lastname string) {
    fmt.Println("Hello", firstname, lastname)
}

sayHelloTo("Uda", "Putera") // Hello Uda Putera
```

### return

```go
func getHello(name string) string {
    return "Hello " + name
}

fmt.Println(getHello("Putera")) // Hello Putera
```

### anonymous function

```go
getHello := func(name string) string {
    return "Hello bro " + name
}

fmt.Println(getHello("Putera"))
```

### multiple return value

```go
func getFullName() (string, string) {
    return "uda", "Putera"
}

// anonymus
getFullName := func() (string, string) {
    return "uda", "Putera"
}

firstname, lastname := getFullName();
fmt.Println(firstname, lastname)
```

### Ignore return value with \_

```go
// anonymus
getFullName := func() (string, string) {
    return "uda", "Putera"
}

firstname, _ := getFullName();
fmt.Println(firstname)
```

### Return value as variable

- If varibale is not set, it will return default value based on type

```go
// anonymus
getFullName := func() (firstname string, lastname string) {
    firstname = "Uda"
    lastname = "Putera"
    return firstname, lastname
}

firstname, lastname := getFullName()
fmt.Println(firstname, lastname)
```

or

```go
// return variable type at last variable
getFullName := func() (firstname, lastname string) {
    // auto asign first data as firstname, second as lastname
    return "uda", "Putera"
}

firstname, lastname := getFullName()
fmt.Println(firstname, lastname)
```

### Variadic Function

- Last argument as slice (array)

```go
sumAll := func(numbers ...int) int {
    total := 0

    for _, number := range numbers {
        total += number
    }

    return total
}

fmt.Println(sumAll(1, 2, 7, 7)) // 17
```

- using slice
- numbers... is converting to variable argument

```go
numbers := []int{1,2,7,7}
fmt.Println(sumAll(numbers...)) // 17
```

### Function as Value

```go
func getGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	goodbye := getGoodbye  // make variable from function
	fmt.Println(goodbye("Putera"))
}
```

### Function as Parameter

```go
import (
	"fmt"
	"strings"
)

sayHelloFilter := func(name string, filter func(string) string) string {
    filteredName := filter(name)
    return "Hello " + filteredName
}

spanFilter := func(name string) string {
    if strings.ToLower(name) == "anjing" {
        return "..."
    } else {
        return name
    }
}

fmt.Println(sayHelloFilter("Anjing", spanFilter))
```

### Function Type Declaration

```go
import (
	"fmt"
	"strings"
)

type Filter func(string) string
sayHelloFilter := func(name string, filter Filter) string {
    filteredName := filter(name)
    return "Hello " + filteredName
}

spanFilter := func(name string) string {
    if strings.ToLower(name) == "anjing" {
        return "..."
    } else {
        return name
    }
}

fmt.Println(sayHelloFilter("Anjing", spanFilter))
```
