package hyperlink

import (
	"fmt"
	"os"
)

var (
	// NoLinks :
	NoLinks = os.Getenv("TERM") == "dumb"
)

// Create :
func Create(text string, href string) string {
	if NoLinks {
		return text
	}
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
