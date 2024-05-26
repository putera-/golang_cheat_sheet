package main

import "fmt"

func main() {
	fmt.Println(_getHello("Uda"))
}
func _getHello(name string) string {
	return "Hello " + name
}
