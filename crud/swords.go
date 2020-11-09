package main

import "fmt"

type BagOfSwords struct {
	speciality string
	swords map[int]*Sword
}

type Sword struct {
	ID int
	name string
	power int
}

// Create new instance
func NewBagOfSwords(speciality string) BagOfSwords {
	return BagOfSwords {
		speciality: speciality,
		swords: make(map[int]*Sword),
	}
}

// Adds a new Sword
func (b BagOfSwords) Add(s Sword) {
	b.swords[s.ID] = &s
}

// Prints all swords
func (b BagOfSwords) Print() {
	for _, v := range b.swords {
		fmt.Printf("[%v]\t%v %v\n", v.ID, v.name, v.power)
	}
}

// Update a Sword
func (b BagOfSwords) update(sword Sword) {
	b.swords[sword.ID] = &sword
}

// Delete by ID
func (b BagOfSwords) deleteByID(ID int) {
	delete(b.swords, ID)
}

// Find by ID returns value of Sword
func (b BagOfSwords) findByID(ID int) *Sword {
	return b.swords[ID]
}