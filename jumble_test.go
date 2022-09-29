package jumble

import (
	"fmt"
)

func main() {
	var text string = "Can Our Brains Really Read Jumbled Words as Long as The First And Last Letters Are Correct?"
	fmt.Println(Jumble(text))
}

func jumbleTest() {

	fmt.Print(
		Jumble("Our new!"),
		Jumble("     whitespace     \n"),
		Jumble("prooof "),
		Jumble("function function function\n"),
		Jumble("Some words!"),
	)
}
