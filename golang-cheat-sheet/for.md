# FOR LOOP

```go
counter := 1

for counter <= 10 {
    fmt.Println(counter)
    counter++
}
// output:
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10
```

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// output:
// 0
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
```

### for range

```go
_names := []string{"Uda", "Putera", "Ganteng"}
for i := 0; i < len(_names); i++ {
    fmt.Println(_names[i])
}

// output:
// Uda
// Putera
// Ganteng

// first key -> index
// second key -> value
for i, name := range _names {
    fmt.Println(i, name)
}
// output:
// 0 Uda
// 1 Putera
// 2 Ganteng
for _i_, name := range _names {
    fmt.Println(name)
}
// output:
// Uda
// Putera
// Ganteng
```
