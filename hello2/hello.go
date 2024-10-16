package hello

import "fmt"

// Language Code
const (
	spanish = "Spanish"
	french  = "French"
)

// Prefix
const (
	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	frenchHelloPrefix  = "Bonjour"
)

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s, %s", getGreetingPrefix(language), name)
}

func main() {
	fmt.Printf("%s\n", Hello("World", ""))
}
