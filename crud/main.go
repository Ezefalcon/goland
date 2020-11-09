package main

func main() {
	bagOfSwords := NewBagOfSwords("Short range")
	bagOfSwords.Add(Sword{ID: 1, name: "Exc", power: 1})
	bagOfSwords.Add(Sword{ID: 2, name: "Exc2", power: 50})
	bagOfSwords.Add(Sword{ID: 3, name: "Exc3", power: 100})
	bagOfSwords.Print()

	bagOfSwords.update(Sword{ID: 3, name: "Updated", power: 120})
	bagOfSwords.Print()

	bagOfSwords.deleteByID(3)
	bagOfSwords.Print()

	bagOfSwords.findByID(2)

}
