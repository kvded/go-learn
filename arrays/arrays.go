package arrays

import (
	"fmt"

	"github.com/kvded/go-learn/arrays/human"
)

func Exec() {
	people := human.Generate(20)
	fmt.Println(people)
}
