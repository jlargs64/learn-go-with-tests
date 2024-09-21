package hello

import "fmt"

const (
	English            = "English"
	EnglishHelloPrefix = "Hello, "
	SpanishHelloPrefix = "Hola, "
	Spanish            = "Spanish"
	FrenchHelloPrefix  = "Bonjour, "
	French             = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case Spanish:
		prefix = SpanishHelloPrefix
	case French:
		prefix = FrenchHelloPrefix
	default:
		prefix = EnglishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Justin", "English"))
}
