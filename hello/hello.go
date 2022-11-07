package main

import "fmt"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	const greetingPrefix = "Hello"
	return fmt.Sprintf("%s, %s", greetingPrefix, name)
}

func main() {
	fmt.Println(Hello("World"))
}
