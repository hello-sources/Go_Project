package main

import _case "goBasic/interface/case"

func main() {
	cat := _case.NewCat()
	animalLife(cat)
	dog := _case.NewDog()
	animalLife(dog)
	dove := _case.NewDove()
	animalLife(dove)
}

func animalLife(a _case.Animal) {
	a.Run()
	a.Sleep()
	a.Each()
	a.Drink()
}
