// Unit Tests
// Functional Tests
// Integration Tests
package main

import (
	"fmt"
	"testing"
)

func translate(s string) string {
	switch s {
	case "en-US":
		return "Hello "
	case "fr-FR":
		return "Bonjour "
	case "it-IT":
		return "Ciao "
	default:
		return "Hello "
	}
}

func Greeting(name, locale string) string {
	salutation := translate(locale)
	return (salutation + name)
}

func TestFrTranslation(t *testing.T) {
	got := translate("fr")
	want := "Bonjour "
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}
func TestUSTranslation(t *testing.T) {
	got := Greeting("George", "en-US")
	want := "Hello George"
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}

func main() {
	fmt.Println("Hello World")
}
