package main

import (
	"fmt"

	j "github.com/shateq/jumble"
)

var nl string = "\n"

func main() {
	jumbleTest()
}

func jumbleTest() {
	var text string = "Can Our Brains Really Read Jumbled Words as Long as The First And Last Letters Are Correct?"
	fmt.Println(j.Jumble(text))

	fmt.Print(
		j.Jumble("Our new!"),
		j.Jumble("     whitespace     \n"),
		j.Jumble("prooof "),
		j.Jumble("function function function\n"),
		j.Jumble("Some words!"),
	)
}
