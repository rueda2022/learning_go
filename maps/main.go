package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"blue":  "#4bf745",
	}

	print_map(colors)

}

func print_map(c map[string]string) {
	for color, hex := range c {
		fmt.Print(color, hex)
	}
}
