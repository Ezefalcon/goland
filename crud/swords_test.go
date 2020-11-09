package main

import (
	"fmt"
	"testing"
)

var bag BagOfSwords

func init() {
	bag = NewBagOfSwords("Long range")
}

func TestBagOfSwords_Add(t *testing.T) {
	sword := bag.findByID(0)
	if sword != nil {
		t.Error("The sword with ID 0 already exists")
	}

	bag.Add(Sword{0,"Excalibur", 10})
	sword = bag.findByID(0)
	if sword == nil {
		t.Error("The sword could not be added")
	}
}

func TestBagOfSwords_findByID(t *testing.T) {
	sword := Sword{0, "Excalibur", 10}
	bag.Add(sword)
	if bag.findByID(0) != nil {
		t.Error("The sword could not be found")
	}
}

func TestBagOfSwords_Update(t *testing.T) {
	bag.Add(Sword{0,"Excalibur", 10})
	bag.update(Sword{0, "Updated", 1})
	sword := bag.findByID(0)

	fmt.Println(sword.name )
	if sword == nil {
		t.Error("The sword could not be updated")
	} else if sword.name != "Upadted" || sword.power != 1 {
		t.Error("The update is wrong")
	}
}

func TestBagOfSwords_Delete(t *testing.T) {
	bag.Add(Sword{0,"Excalibur", 10})
	sword := bag.findByID(0)
	if sword == nil {
		t.Error("The sword could not be added")
	}
	bag.deleteByID(0)
	sword = bag.findByID(0)
	if sword != nil {
		t.Error("The sword coult not be deleted")
	}
}

