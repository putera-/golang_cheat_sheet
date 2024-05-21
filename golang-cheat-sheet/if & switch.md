# IF

```go
name := "putera"

if name == "putera" {
    fmt.Println(name)
}

if name == "budi" {
    fmt.Println(name)
} else {
    fmt.Println("bukan budi")
}

if name == "budi" {
    fmt.Println(name)
} else if name == "joko" {
    fmt.Println(name)
} else {
    fmt.Println("bukan budi atau joko")
}
```

### short statement

declare variable before check condition

```go
// if [declaration] [condition] { code block }
if length := len(name); length > 5 {
    fmt.Println("nama terlalu panjang")
} else {
    fmt.Println("nama tidak terlalu panjang")
}
```

# SWITCH

```go
name := "putera"

switch name {
case "putera":
    fmt.Println("Hello Putera")
case "joko":
    fmt.Println("Hello Joko")
default:
    fmt.Println("kamu siapa")
}
```

### short statement

```go
// switch [declaration] [condition] { code block }
switch length := len(name); length > 5 {
case true:
    fmt.Println("nama terlalu panjang")
case false:
    fmt.Println("nama tidak terlalu panjang")
}
```

### without condition

```go
// switch [declaration] [condition] { code block }
length := len(name);
switch {
case length > 5:
    fmt.Println("nama terlalu panjang")
case length <= 5:
    fmt.Println("nama tidak terlalu panjang")
}
```
