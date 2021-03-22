package main

import "fmt"

func main() {
	fmt.Println(XShape('A'))
}

func XShape(char byte) (result string) {
	return fmt.Sprintf("%c", char)
}
