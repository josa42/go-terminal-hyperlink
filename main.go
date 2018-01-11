package hyperlink

import "fmt"

// Create :
func Create(text string, href string) string {
	return fmt.Sprintf("\x1b]8;;%s\a%s\x1b]8;;\a", href, text)
}

// Print :
func Print(text string, href string) {
	link := Create(text, href)
	fmt.Print(link)
}

// Println :
func Println(text string, href string) {
	link := Create(text, href)
	fmt.Println(link)
}
