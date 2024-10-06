// cmd/golang-package/main.go
package main

import (
    "fmt"
	"golang-package/pkg/exinit"
)

func main() {
	fmt.Println("main function")
	exinit.PrintD()
}