package chapter1

import "fmt"

// Person is a struct used to store details about people including friends.
type Person struct {
	FirstName   string
	LastName    string
	EmailID     string
	PhoneNumber string
	Birthday    string
}

// FullName returns the full name of a person
// as "<First Name> <Last Name>"
func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

// ChangeEmailID updates the Email ID of person
// with the new Email ID provided as second argument
func (p *Person) ChangeEmailID(newEmailID string) {
	p.EmailID = newEmailID
}

// Alice stores details about Alice in Person struct
var Alice = Person{
	FirstName:   "Alice",
	LastName:    "A.",
	EmailID:     "alice@email.com",
	PhoneNumber: "xx-xxx",
	Birthday:    "01-01-1972",
}

// Bob stores details about Bob in Person struct
var Bob = Person{
	FirstName:   "Bob",
	LastName:    "B.",
	EmailID:     "bob@email.com",
	PhoneNumber: "xx-xxx",
	Birthday:    "01-01-1972",
}

// GetFullName returns the full name of a person
// as "<First Name> <Last Name>"
func GetFullName(p Person) string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

// ChangeEmailID updates the Email ID of person
// with the new Email ID provided as second argument
func ChangeEmailID(p *Person, newEmailID string) {
	p.EmailID = newEmailID
}

func manipulateStructs() {

	fmt.Println("Bob's Full Name using C-style functions: " + GetFullName(Bob))

	fmt.Println("Alice's Email ID before change: " + Alice.EmailID)
	ChangeEmailID(&Alice, "new_alice@email.com")
	fmt.Println("Alice's Email ID after change: " + Alice.EmailID)
}
