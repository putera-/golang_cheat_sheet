# DEFER, PANIC, RECOVER

### Defer

- will be execute last, event if error happen

```go
func logging() {
    fmt.Println("Logging done")
}

func runApp() {
    defer logging()
    fmt.Println("App running")
}

func main() {
    runApp()
}

// output:
// App Running
// Logging Done
```
