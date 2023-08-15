package main

import "fmt"

/* This line begins our Animal class (... well, its eequivalent in Go) */

// Here we define the animal struct: the animal "class" sans the verbs/methods.
type Animal struct {
	species    string
	food       string
	locomotion string
	noise      string
}

/* We write these three methods with a receiver, so that they only work on the Animal struct.*/
// we create the "verbs" for our object the animal.
func (a Animal) Eat() {
	fmt.Printf("%v eating %v...!\n", a.species, a.food)
}

func (a Animal) Move() {
	fmt.Printf("%v", a.locomotion)
}

func (a Animal) Speak() {
	fmt.Printf("%v", a.noise)
}

/* The afore three methods only work on the Animal struct. There is no outher outlet to call them. */
/* This line concludes our Animal "class" */

func main() {
	/* This line begins our object code: "turning class into objects by instantiating them." */
	// dog
	dog := Animal{"dog", "dogfood", "walk", "woof"}

	// pig
	pig := Animal{"pig", "soybeans", "walk", "oink"}

	// donkey
	donkey := Animal{"donkey", "hay", "", ""}

	// horse
	horse := Animal{"horse", "grass", "", ""}
	/* This line concludes our object instantiations code. */

	// We are going to use maps to handle user choices.
	// We can go with if or switch statements, but
	// what if there are 1,000 animals, and each animal has 1,000 possible verbs?
	// We wouldn't want the inefficiency of going through a list;
	// we want instant, random access.
	// So we use map: we are scalable and ready for 3 or 3,000 options.
	dogMap := map[string]func(){
		"eat":   dog.Eat,
		"move":  dog.Move,
		"speak": dog.Speak,
	}

	pigMap := map[string]func(){
		"eat":   pig.Eat,
		"move":  pig.Move,
		"speak": pig.Speak,
	}

	donkeyMap := map[string]func(){
		"eat":   donkey.Eat,
		"move":  donkey.Move,
		"speak": donkey.Speak,
	}

	horseMap := map[string]func(){
		"eat":   horse.Eat,
		"move":  horse.Move,
		"speak": horse.Speak,
	}

	// we define dogMaps et al. above this line,
	// because we need them in this animalFarm assignment:
	animalFarm := map[string]map[string]func(){
		"dog":    dogMap,
		"pig":    pigMap,
		"donkey": donkeyMap,
		"horse":  horseMap,
	}
	/* The code lines for maps conclude here */

	// declare var to hold the animal that the user wants:
	var nimal, action string

	// Loop for user input:
	for {
		fmt.Println("> (Type \"exit\" to exit the program.)")
		fmt.Println("> Which animal would you like to select?")
		fmt.Println("> Options: dog, pig, donkey, horse")
		fmt.Printf("> ")
		fmt.Scan(&nimal)
		if nimal == "exit" || (nimal != "dog" && nimal != "pig" && nimal != "donkey" && nimal != "horse") {
			fmt.Println("exiting...")
			return
		}
		fmt.Printf("> OK, what action would you like the %v to perform?\n", nimal)
		fmt.Println("> Options: eat, move, speak")
		fmt.Printf("> ")
		fmt.Scan(&action)
		if action == "exit" || (action != "eat" && action != "move" && action != "speak") {
			fmt.Println("exiting...")
			return
		}

		// deliver user's request:
		animalFarm[nimal][action]()
		fmt.Println("")
	}

}
