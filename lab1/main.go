package main

import "fmt"

// go mod init example/hello
// go run main.go
// go build -o name.exe

func main() {
	var liczba1 int64 = 1
	var (
		x float64 = 2.5
		y string  = "He"
	)

	liczba1++

	fmt.Println("Hello world", liczba1, x, y)
}
