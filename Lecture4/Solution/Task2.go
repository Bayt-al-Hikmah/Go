package main

import "fmt"

type Writer interface {
	Write(data string) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data string) (int, error) {
	fmt.Println(data) // print to console
	return len(data), nil
}

type StringWriter struct {
	content string
}

func (sw *StringWriter) Write(data string) (int, error) {
	sw.content += data 
	return len(data), nil
}

func main() {
	cw := ConsoleWriter{}
	n1, err1 := cw.Write("Hello from ConsoleWriter!")
	fmt.Println("Bytes written:", n1, "Error:", err1)

	sw := StringWriter{}
	n2, err2 := sw.Write("Hello, ")
	n3, err3 := sw.Write("Golang Writer!")

	fmt.Println("Bytes written:", n2, err2)
	fmt.Println("Bytes written:", n3, err3)

	fmt.Println("StringWriter content:", sw.content)
}