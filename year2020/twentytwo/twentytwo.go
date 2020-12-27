package twentytwo

import (
	"aoc-go/files"
	"aoc-go/set"
	"aoc-go/utils"
	"fmt"
	"strconv"
	"strings"
)

// PartOne - not yet implemented
func PartOne(filename string) string {
	p1, p2 := parseDecks(filename)
	p1Winner := playGame(&p1, &p2, func(d1 *deck, d2 *deck) (bool, bool) { return d1.battle(d2) }, false)
	var winnerScore int
	if p1Winner {
		winnerScore = p1.calcScore()
	} else {
		winnerScore = p2.calcScore()
	}
	return fmt.Sprint(winnerScore)
}

// PartTwo - not yet implemented
func PartTwo(filename string) string {
	p1, p2 := parseDecks(filename)
	p1Winner := playGame(&p1, &p2, func(d1 *deck, d2 *deck) (bool, bool) { return d1.recursiveBattle(d2) }, true)
	var winnerScore int
	if p1Winner {
		winnerScore = p1.calcScore()
	} else {
		winnerScore = p2.calcScore()
	}
	return fmt.Sprint(winnerScore)
}

func parseDecks(filename string) (deck, deck) {
	fileStream := make(chan string)
	go files.StreamLines(filename, fileStream)
	p1cards := make([]int, 0)
	p2cards := make([]int, 0)
	onP1 := true
	for line := range fileStream {
		if line == "" || line == "Player 1:" {
			continue
		}
		if line == "Player 2:" {
			onP1 = false
			continue
		}
		card := utils.MustAtoi(line)
		if onP1 {
			p1cards = append(p1cards, card)
		} else {
			p2cards = append(p2cards, card)
		}
	}
	p1deck := deck{p1cards}
	p2deck := deck{p2cards}
	return p1deck, p2deck
}

func playGame(p1 *deck, p2 *deck, battle func(p1 *deck, p2 *deck) (bool, bool), finishOnRepeatedState bool) (p1Winner bool) {
	finished := false
	p1Winner = false
	previousGameStates := set.MakeStringSet()
	for !finished {
		gameState := serializeGameState(*p1, *p2)
		if finishOnRepeatedState && previousGameStates.Contains(gameState) {
			return true
		}
		previousGameStates.Add(serializeGameState(*p1, *p2))
		finished, p1Winner = battle(p1, p2)
	}
	return p1Winner
}

func serializeGameState(p1 deck, p2 deck) string {
	str := "p1:"
	p1strs := make([]string, len(p1.cards))
	p2strs := make([]string, len(p2.cards))
	for i, card := range p1.cards {
		p1strs[i] = strconv.Itoa(card)
	}
	for i, card := range p2.cards {
		p2strs[i] = strconv.Itoa(card)
	}
	str += strings.Join(p1strs, ",")
	str += ":p2:"
	str += strings.Join(p2strs, ",")
	return str
}

type deck struct {
	cards []int
}

func (d *deck) draw() (int, bool) {
	if d.isEmpty() {
		return -1, false
	}
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card, true
}

func (d *deck) addToBottom(card int) {
	d.cards = append(d.cards, card)
}

func (d *deck) battle(d2 *deck) (end bool, amWinner bool) {
	if d.isEmpty() {
		return true, false
	}
	if d2.isEmpty() {
		return true, true
	}
	p1card, _ := d.draw()
	p2card, _ := d2.draw()
	if p1card > p2card {
		d.addToBottom(p1card)
		d.addToBottom(p2card)
	} else {
		d2.addToBottom(p2card)
		d2.addToBottom(p1card)
	}
	return false, false
}

func (d *deck) recursiveBattle(d2 *deck) (end bool, amWinner bool) {
	if d.isEmpty() {
		return true, false
	}
	if d2.isEmpty() {
		return true, true
	}
	p1card, _ := d.draw()
	p2card, _ := d2.draw()
	p1CanRecurse := p1card <= len(d.cards)
	p2CanRecurse := p2card <= len(d2.cards)
	if p1CanRecurse && p2CanRecurse {
		dCopy := d.copy(p1card)
		d2Copy := d2.copy(p2card)
		p1SubWinner := playGame(&dCopy, &d2Copy, func(d1 *deck, d2 *deck) (bool, bool) { return d1.recursiveBattle(d2) }, true)
		if p1SubWinner {
			d.addToBottom(p1card)
			d.addToBottom(p2card)
		} else {
			d2.addToBottom(p2card)
			d2.addToBottom(p1card)
		}
	} else {
		if p1card > p2card {
			d.addToBottom(p1card)
			d.addToBottom(p2card)
		} else {
			d2.addToBottom(p2card)
			d2.addToBottom(p1card)
		}
	}
	return false, false
}

func (d deck) copy(amount int) deck {
	newCards := make([]int, amount)
	for i := 0; i < amount && i < len(d.cards); i++ {
		newCards[i] = d.cards[i]
	}
	return deck{newCards}
}

func (d deck) isEmpty() bool {
	return len(d.cards) == 0
}

func (d deck) calcScore() int {
	score := 0
	value := len(d.cards)
	for i := 0; i < len(d.cards); i++ {
		score += d.cards[i] * value
		value--
	}
	return score
}
