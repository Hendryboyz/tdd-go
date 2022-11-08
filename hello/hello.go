package main

import "fmt"

const SPANISH_LANG = "Spanish"
const FRENCH_LANG = "French"
const ENGLISH_GREETING = "Hello"
const SPANISH_GREETING = "Hola"
const FRENCH_GREETING = "Bonjour"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	greetingPrefix := ENGLISH_GREETING
	if language == SPANISH_LANG {
		greetingPrefix = SPANISH_GREETING
	} else if language == FRENCH_LANG {
		greetingPrefix = FRENCH_GREETING
	}
	return fmt.Sprintf("%s, %s", greetingPrefix, name)
}

func main() {
	fmt.Println(Hello("World", ""))
}
