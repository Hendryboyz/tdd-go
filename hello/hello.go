package main

import "fmt"

const SPANISH_LANG = "Spanish"
const ENGLISH_GREETING = "Hello"
const SPANISH_GREETING = "Hola"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	greetingPrefix := ENGLISH_GREETING
	if language == SPANISH_LANG {
		greetingPrefix = SPANISH_GREETING
	}
	return fmt.Sprintf("%s, %s", greetingPrefix, name)
}

func main() {
	fmt.Println(Hello("World", ""))
}
