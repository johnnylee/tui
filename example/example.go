package main

import (
	"fmt"
	"github.com/johnnylee/tui"
)

func main() {
	tui.Clear()

	s := tui.Menu(
		"My Title",
		nil,
		"a", "Ace",
		"b", "Base",
		"c", "Case?")

	fmt.Println("Input:", s)
	tui.Line()

	s = tui.String("String input")
	fmt.Println("Input:", s)
	tui.Line()
	
	s = tui.StringNotEmpty("String input (not empty)")
	fmt.Println("Input:", s)
	tui.Line()

	i := tui.Int("Integer input")
	fmt.Println("Input:", i)
	tui.Line()

	j := tui.Float("Float input")
	fmt.Println("Input:", j)
	tui.Line()
}
