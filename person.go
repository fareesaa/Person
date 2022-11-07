package person

import "fmt"

type Person struct {
	Id    int
	Name  string
	Hobby string
}

// SetId receives a pointer to Person so it can modify it.
func (f *Person) SetId(id int) {
	f.Id = id
}

// Id receives a copy of Person since it doesn't need to modify it.
func (f Person) GetId() int {
	return f.Id
}

// SetName receives a pointer to Person so it can modify it.
func (f *Person) SetName(name string) {
	f.Name = name
}

// Name receives a copy of Person since it doesn't need to modify it.
func (f Person) GetName() string {
	return f.Name
}

// SetHobby receives a pointer to Person so it can modify it.
func (f *Person) SetHobby(hobby string) {
	f.Hobby = hobby
}

// Name receives a copy of Person since it doesn't need to modify it.
func (f Person) GetHobby() string {
	return f.Hobby
}

func (f Person) ToString() {
	fmt.Printf("Halo nama saya %s, hobi saya adalah %s", f.Name, f.Hobby)
}
