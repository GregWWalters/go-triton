Triton🐚 is a text-styling package for Unix-like shells, providing easy-to-use
ANSI formatting.

# Usage

```go
package main

import "github.com/gregwwalters/go-triton"

func main() {
	boldRed := triton.NewStyle().
		WithForeground(triton.Red).
		WithAttributes(triton.Bold)
	log.Println(boldRed.Apply("You shall not pass!")
	warning := "You " + triton.Inverted("shall not") + " pass!"
	log.Println(boldRed.Apply(warning))
	log.Println("Tell me. Where is " +
		triton.White.WithAttributes(triton.Dim, triton.Underlined).Apply("Gandalf") +
		"? For I much desire to speak with him.")
}
```

# Colors

* Black
* Red
* Green
* Yellow
* Blue
* Magenta
* Cyan
* White

# Attributes

* Normal
* Bold
* Dim
* Italic
* Underlined
* Blinking
* Inverted
* Hidden
* Strikethrough

