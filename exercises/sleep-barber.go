package main

// Person represents the barber and the customers
type Person struct {
	ID       int
	Sleeping bool
}

// Chair represents the chair where the barber is cutting hair and the waiting room
type Chair struct {
	Occupied bool
	Person   *Person
}

// WaitingRoom represents the waiting room
type WaitingRoom struct {
	Sits map[*Person]bool
}

func main() {
	customersCh := make(chan Person)

}
