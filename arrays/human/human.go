package human

import "fmt"

type EyeColor int

const (
	Brown EyeColor = iota + 1
	Blue
	Red
	Green
	White
	Black
)

func (h EyeColor) String() string {
	return [...]string{"brown", "blue", "red", "green", "white", "black"}[h]
}

type Person struct {
	FirstName string   `faker:"first_name"`
	LastName  string   `faker:"last_name"`
	Age       int      `faker:"boundary_start=16, boundary_end=99"`
	EyeColor  EyeColor `faker:"boundary_start=0, boundary_end=5"`
	Height    int      `faker:"boundary_start=150, boundary_end=210"`
}

func (h Person) String() string {
	return fmt.Sprintf("{\n"+
		"  First Name: %s\n"+
		"  Last Name: %s\n"+
		"  Age: %d\n"+
		"  Eye Color: %s\n"+
		"  Height (cm): %d\n"+
		"}\n", h.FirstName, h.LastName, h.Age, h.EyeColor, h.Height)
}
