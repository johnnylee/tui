package tui

import (
	"os"
	"fmt"
	"github.com/GeertJohan/go.linenoise"
	"strings"
)

func printBoxed(text string) {
	l := len(text)

	// Top of box. 
	fmt.Printf("╭")
	for i := 0; i < l + 2; i++ {
		fmt.Printf("─")
	}
	fmt.Printf("╮\n")
	
	// Text. 
	fmt.Printf("│ %v │\n", text)
	
	// Bottom of box. 
	fmt.Printf("╰")
	for i := 0; i < l + 2; i++ {
		fmt.Printf("─")
	}
	fmt.Printf("╯\n")
}

func paddedString(text string, length int) string {
	for len(text) < length {
		text = text + " "
	}
	return text
}

// Clear clears the screen. 
func Clear() {
	linenoise.Clear()
}

// Line prints a line across the screen ending in a newline.
func Line() {
	fmt.Printf("┅")
	for i := 0; i < 77; i++ {
		fmt.Printf("━")
	}
	fmt.Printf("┅\n")
}

// String prompts a user for a string and returns the result. 
// The prompt will have a colon appended to it. 
// The resulting string is stripped of any whitespace.
func String(prompt string) string {
	prompt = fmt.Sprintf("%v: ", prompt)

	s, err := linenoise.Line(prompt)
	if err != nil {
		os.Exit(0)
	}
	s = strings.TrimSpace(s)

	if len(s) > 0 {
		linenoise.AddHistory(s)
	}

	return s
}

// StringNotEmpty prompts the user for a non-empty string. 
// It keeps asking until a string is returned. 
func StringNotEmpty(prompt string) string {
	for {
		s := String(prompt)
		if len(s) > 0 {
			return s
		}
	}
}

// Int prompts a user for an integer. 
func Int(prompt string) int64 {
	for {
		s := String(prompt)

		value := int64(0)
		
		n, err := fmt.Sscanf(s, "%d", &value)
		if n != 1 || err != nil {
			continue
		}
		return value
	}
}

// Float prompts the user for a floating point value. 
func Float(prompt string) float64 {
	for {
		s := String(prompt)

		value := float64(0)
		
		n, err := fmt.Sscanf(s, "%g", &value)
		if n != 1 || err != nil {
			continue
		}
		return value
	}
}

// Menu displays a menu allowing a user to choose a single item.
// Lines in text will be displayed before the menu items. 
// Menu items are just key-value string pairs. 
func Menu(title string, text []string, items ...string) string {
	// Create map of known keys, and find the maximum key length. 
	keyMap := make(map[string]bool)
	keyLen := 0
	
	for i := 0; i < len(items); i+=2 {
		key := items[i]
		if len(key) > keyLen {
			keyLen = len(key)
		}
		keyMap[key] = true
	}
	
	for {
		// Print title. 
		printBoxed(title)
		fmt.Printf("┆\n")
		
		// Print any extra text. 
		if len(text) > 0 {
			for _, t := range text {
				fmt.Printf("│ %v\n", t)
			}
			fmt.Printf("│\n")
		}
		
		// Print items. 
		for i := 0; i < len(items); i+=2 {
			key := paddedString(items[i], keyLen)
			txt := items[i+1]
			
			fmt.Printf("│ %v %v\n", key, txt)
		}
		
		fmt.Printf("│\n")
		fmt.Printf("╰─┄\n")

		// Prompt user. 
		s := String("Selection")

		if _, ok := keyMap[s]; ok {
			return s
		} 

		Clear()
	}
}

