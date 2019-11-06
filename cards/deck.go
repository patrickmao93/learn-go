package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	suits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	values := []string{"Ace", "Two", "Three", "Four"}
	d := deck{}
	for _, s := range suits {
		for _, v := range values {
			d = append(d, v+" of "+s)
		}
	}
	return d
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) saveToFile(filename string) error {
	str := strings.Join(d, ", ")
	err := ioutil.WriteFile(filename, []byte(str), 0666)
	return err
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return strings.Split(string(bs), ", ")
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		pos := r.Intn(len(d))
		d[i], d[pos] = d[pos], d[i]
	}
	return d
}
