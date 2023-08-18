package main

import (
	"fmt"
	"reflect"
)

// Make the interface.
// We set the Animal to be interface for simplicity and efficient code.
// and also to show Go's interface :)
type Animal interface {
	Eat()
	Move()
	Speak()
}

// make species to handle user request to
// create new animals
type Species struct {
	Name       string
	Food       string
	Locomotion string
	Noise      string
}

func main() {

	/* Animal Code Begins: bring in the farm animals: */
	// dog
	dog := Species{"dog", "dogfood", "walk", "woof"}

	// pig
	pig := Species{"pig", "soybeans", "walk", "oink"}

	// donkey
	donkey := Species{"donkey", "hay", "walk", "hee-haw"}

	// horse
	horse := Species{"horse", "grass", "walk", "neigh"}

	// This next one is the space we set for if the user creates a new animal:
	// It is also the struct we will modify at runtime using Reflection :)
	newanimal := Species{"", "", "", ""}
	/* Animal instantiation concludes here. */

	/* Maps begin (we use map for random access: for scalability & speed) */

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

	// We need those dogMaps et al. in this var assignment:
	animalFarm := map[string]map[string]func(){
		dog.Name: dogMap,
		"pig":    pigMap,
		"donkey": donkeyMap,
		"horse":  horseMap,
		// newanimal.Name: newanimalMap,
	}
	/* Maps conclude. */

	/* Code for user input begins: */

	// declare var to hold user inputs:
	var nimal, action string

	// var to modify the thing:
	var speciesname, itsfood, itsloco, itsnoise string
	newnimal := reflect.ValueOf(&newanimal).Elem()
	fmt.Println("newnimal's settability: ", newnimal.CanSet())

	for { // Loop starts.
		fmt.Println("> (Type \"exit\" to exit the program.)")
		fmt.Println("> Which animal would you like to select?")
		fmt.Println("> Options: dog, pig, donkey, horse, or the name of the new animal you made:")
		fmt.Printf("> ")
		fmt.Scan(&nimal)
		if nimal == "exit" || (nimal != "dog" && nimal != "pig" &&
			nimal != "donkey" && nimal != "horse" && nimal != newanimal.Name) {
			fmt.Println("exiting...")
			return
		}
		fmt.Printf("> OK, what action would you like the %v to perform?\n", nimal)
		fmt.Println("> Options: eat, move, speak")
		fmt.Printf("> ")
		fmt.Scan(&action)
		if action == "exit" || (action != "eat" && action != "move" &&
			action != "speak") {
			fmt.Println("exiting...")
			return
		}

		// deliver user's request:
		animalFarm[nimal][action]()
		fmt.Println("")

		// conditional: if there is sitll spcae for newnimal
		if newanimal.Name == "" {
			fmt.Println("Hey let's make a new animal! We have one space available.")
			//...
			fmt.Println("Please enter the species name: (e.g. bird)")
			fmt.Scan(&speciesname)
			newnimal.Field(0).SetString(speciesname)

			fmt.Println("What does it eat? (e.g. worm)")
			fmt.Scan(&itsfood)
			newnimal.Field(1).SetString(itsfood)

			fmt.Println("What is its locomotion method? (e.g. fly)")
			fmt.Scan(&itsloco)
			newnimal.Field(2).SetString(itsloco)

			fmt.Println("What sound does it make when it speaks? (e.g. chirp)")
			fmt.Scan(&itsnoise)
			newnimal.Field(3).SetString(itsnoise)
			newanimalMap := map[string]func(){
				"eat":   newanimal.Eat,
				"move":  newanimal.Move,
				"speak": newanimal.Speak,
			}
			animalFarm[newanimal.Name] = newanimalMap
		} // Now we need to print the new user input

	} // Loop concludes.
	/* User input code concludes here. */

}

// These three functions are so that our animals satisfy the Animal interface. //
func (nimal Species) Eat() {
	fmt.Printf("%v eating %v...!\n", nimal.Name, nimal.Food)
}

func (nimal Species) Move() {
	fmt.Printf("%ving...", nimal.Locomotion)
}

func (nimal Species) Speak() {
	fmt.Printf("%v %v", nimal.Noise, nimal.Noise)
}
