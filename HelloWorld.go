package main

import "fmt"

const helloPrefix = "Hello "
const smile = " :)"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name + smile
}

func main() {
	fmt.Println(Hello("world"))
}
