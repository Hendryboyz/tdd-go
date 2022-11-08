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
	greetingPrefix := getGreetingPrefix(language)
	return fmt.Sprintf("%s, %s", greetingPrefix, name)
}

func getGreetingPrefix(language string) string {
	switch language {
	case SPANISH_LANG:
		return SPANISH_GREETING
	case FRENCH_LANG:
		return FRENCH_GREETING
	default:
		return ENGLISH_GREETING
	}
}

func main() {
	fmt.Println(Hello("World", ""))
}
