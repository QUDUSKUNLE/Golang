package hello

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix

	switch language {
		case "french":
			prefix = frenchHelloPrefix

		case "spanish":
			prefix = spanishHelloPrefix
		
		default:
			prefix = englishHelloPrefix
	}
	return prefix + name
}

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func ReverseRunes(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
