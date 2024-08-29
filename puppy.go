package puppy

import (
	"fmt"

	"github.com/EthaGarz/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Bark())
}

func From1() {
	fmt.Println("Version 1.0.0")
}
