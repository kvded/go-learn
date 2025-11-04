package human

import (
	"fmt"

	"github.com/go-faker/faker/v4"
)

func Generate(amount int) []Person {
	var people []Person
	numOfPeople := amount

	for i := 0; i < numOfPeople; i++ {
		var person Person
		err := faker.FakeData(&person)
		if err != nil {
			fmt.Printf("Error when generating Humans with Faker; err: %v+", err)
		}
		people = append(people, person)
	}
	return people
}
