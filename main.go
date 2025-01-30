package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	cards := new_deck()
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filname string) error {
	return ioutil.WriteFile(filname, []byte(d.toString()), 0666)
}

func new_deck_from_file(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit((1))
	}

	s := strings.Split((string(bs)), ",")

	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
